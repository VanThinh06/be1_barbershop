package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (*customer.LoginCustomerResponse, error) {

	contact := db.GetContactCustomerParams{
		TypeUsername: "phone",
		Email:        req.Username,
	}

	err := utils.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.TypeUsername = "email"
		contact.Email = req.Username
	}

	res, err := server.store.GetContactCustomer(ctx, contact)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "information is incorrect")
		}
		return nil, status.Error(codes.Internal, "internal")
	}

	if req.IsSocialAuth {
		email, err := server.authVerifyJWTGG(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal")
		}
		if email != req.Username {
			return nil, status.Errorf(codes.Unauthenticated, "information is incorrect")
		}
	}

	if req.IsSocialAuth == false {
		err = utils.CheckPassword(req.Password, res.HashedPassword.String)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "information is incorrect")
		}
	}

	customerPayload := token.Customer{
		CustomerID: res.CustomerID,
		Name:       res.Name,
		Gender:     res.Gender,
		Phone:      res.Phone,
		Email:      res.Email,
		FcmDevice:  req.GetFcmDevice(),
		Timezone:   req.GetTimezone(),
	}
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSessionsCustomer(ctx, db.CreateSessionsCustomerParams{
		ID:           refreshPayload.ID,
		CustomerID:   refreshPayload.Customer.CustomerID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
		FcmDevice:    req.GetFcmDevice(),
		Longitude:    req.GetLatitude().GetValue(),
		Latitude:     req.GetLatitude().GetValue(),
		Timezone:     req.GetTimezone(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &customer.LoginCustomerResponse{
		Customer:              convertCustomer(res),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		Longitude:             req.Longitude,
		Latitude:              req.Latitude,
	}
	return rsp, nil
}
