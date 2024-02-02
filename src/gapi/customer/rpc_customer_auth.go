package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (*customer.LoginCustomerResponse, error) {

	contact := db.GetUserCustomerParams{
		TypeUsername: "phone",
		Email:        req.Username,
	}

	err := helpers.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.TypeUsername = "email"
		contact.Email = req.Username
	}

	socialEmail := &SocialEmail{}
	res, err := server.store.GetUserCustomer(ctx, contact)
	if err != nil {
		if err == sql.ErrNoRows {
			if req.IsSocialAuth == false {
				return nil, returnError(codes.NotFound, "incorrect account or password", err)
			}

			// create account for the first time login
			if req.IsSocialAuth {
				socialEmail, err = server.authVerifyJWTGG(ctx)
				if err != nil {
					return nil, internalError(err)
				}
				if socialEmail.Email != req.GetUsername() {
					return nil, status.Error(codes.Unauthenticated, "information is incorrect")
				}
				argCustomer := db.CreateCustomerParams{
					Name:           socialEmail.GivenName,
					Phone:          sql.NullString{},
					Gender:         int32(utilities.Male),
					Email:          socialEmail.Email,
					IsSocialAuth:   req.GetIsSocialAuth(),
					HashedPassword: sql.NullString{},
				}
				customerSocial, err := server.store.CreateCustomer(ctx, argCustomer)
				if err != nil {
					return nil, status.Error(codes.Unauthenticated, "information is incorrect")
				}
				res = customerSocial
			}
		}

		if socialEmail == nil {
			return nil, internalError(err)
		}
	}

	if req.IsSocialAuth && socialEmail != nil {
		socialEmail, err = server.authVerifyJWTGG(ctx)
		if err != nil {
			return nil, internalError(err)
		}
		if socialEmail.Email != req.Username {
			return nil,unauthenticatedError(err)
		}
	}

	if req.IsSocialAuth == false {
		err = utilities.CheckPassword(req.Password, res.HashedPassword.String)
		if err != nil {
			return nil, unauthenticatedError(err)
		}
	}

	customerPayload := token.Customer{
		CustomerID: res.ID,
		Name:       res.Name,
		Phone:      res.Phone,
		Email:      res.Email,
		FcmDevice:  req.GetFcmDevice(),
	}
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, internalError(err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, internalError(err)
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSessionsCustomer(ctx, db.CreateSessionsCustomerParams{
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		FcmDevice:    req.GetFcmDevice(),
		Longitude:    req.GetLatitude().GetValue(),
		Latitude:     req.GetLatitude().GetValue(),
	})
	if err != nil {
		return nil, internalError(err)
	}

	rsp := &customer.LoginCustomerResponse{
		Customer:              convertCustomer(res),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
		Longitude:             req.Longitude,
		Latitude:              req.Latitude,
	}
	return rsp, nil
}
