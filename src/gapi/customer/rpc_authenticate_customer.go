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

	// Retrieve barber information from the store
	res, err := server.store.GetContactCustomer(ctx, contact)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "failed to get email barber %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get email barber %s", err)
	}

	/// check social
	if req.IsSocialAuth {
		email, err := server.authVerifyJWTGG(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get email barber %s", err)
		}
		if email != req.Username {
			return nil, status.Errorf(codes.Internal, "failed to get email barber %s", err)
		}
	}

	// Check the password provided against the stored hashed password
	if req.IsSocialAuth == false {
		err = utils.CheckPassword(req.Password, res.HashedPassword.String)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "failed to check the password for")
		}
	}

	// create token
	customerPayload := token.Customer{
		CustomerID: res.CustomerID,
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
		return nil, status.Error(codes.Internal, "failed to create token")
	}

	// Create a refresh token for the authenticated barber
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create refresh token")
	}

	// Create a session for the barber
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
		return nil, status.Error(codes.Internal, "failed to create session")
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
