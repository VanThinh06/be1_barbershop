package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (*customer.LoginCustomerResponse, error) {

	contact := db.GetContactCustomerParams{
		Column1: "phone",
		Email:   req.Username,
	}
	// validate
	err := utils.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.Column1 = "email"
	}

	// Retrieve barber information from the store
	res, err := server.store.GetContactCustomer(ctx, contact)
	if err != nil {
		if err == sql.ErrNoRows {
			// If barber does not exist, return a 404 error
			return nil, status.Errorf(codes.NotFound, "failed to get email barber %s", err)
		}
		// If there's an error getting the barber, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to get email barber %s", err)
	}

	// Check the password provided against the stored hashed password
	err = utils.CheckPassword(req.Password, res.HashedPassword)
	if err != nil {
		// If the password check fails, return an unauthenticated error
		return nil, status.Errorf(codes.Unauthenticated, "failed to check the password for %s", err)
	}

	customerPayload := token.Customer{
		CustomerID: res.CustomerID,
		Role:       0,
		Phone:      res.Phone,
		Email:      res.Email,
	}
	// Create an access token for the authenticated barber
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		// If there's an error creating the access token, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create token %s", err)
	}

	// Create a refresh token for the authenticated barber
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		customerPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		// If there's an error creating the refresh token, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create refresh token %s", err)
	}

	mtdt := server.extractMetadata(ctx)
	// Create a session for the barber
	coordinates := fmt.Sprintf("POINT(%.7f %.7f)", 16.059663965551348, 108.21712790268852)
	session, err := server.store.CreateSessionsCustomer(ctx, db.CreateSessionsCustomerParams{
		ID:           refreshPayload.ID,
		CustomerID:   refreshPayload.Customer.CustomerID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
		FcmDevice:    "FcmDevice",
		Coordinates:  coordinates,
	})
	if err != nil {
		// If there's an error creating the session, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create session %s", err)
	}

	// Prepare the response with barber and session information
	rsp := &customer.LoginCustomerResponse{
		Customer:              convertCustomer(res),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}