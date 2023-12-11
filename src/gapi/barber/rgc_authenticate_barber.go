package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginBarber(ctx context.Context, req *barber.LoginBarberRequest) (*barber.LoginBarberResponse, error) {
	// Retrieve barber information from the store
	res, err := server.Store.GetEmailBarber(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "failed to get email barber ")
		}
		return nil, status.Error(codes.Internal, "failed to get email barber")
	}

	// Check the password provided against the stored hashed password
	err = utils.CheckPassword(req.Password, res.HashedPassword)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "failed to check the password for")
	}

	barberPayload := token.BarberPayload{
		BarberID:  res.BarberID,
		Role:      res.Role,
		Phone:     res.Phone,
		Email:     res.Email,
		FcmDevice: req.FcmDevice,
		Timezone:  req.Timezone,
	}

	// Create an access token for the authenticated barber
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create token")
	}

	// Create a refresh token for the authenticated barber
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create refresh token %s")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.Store.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
		ID:           refreshPayload.ID,
		BarberID:     refreshPayload.Barber.BarberID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
		FcmDevice:    req.FcmDevice,
		Timezone:     req.Timezone,
	})
	if err != nil {
		// If there's an error creating the session, return an internal server error
		return nil, status.Error(codes.Internal, "failed to create session")
	}

	// Prepare the response with barber and session information
	rsp := &barber.LoginBarberResponse{
		Barber:                convertBarber(res),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
