package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginBarber(ctx context.Context, req *pb.LoginBarberRequest) (*pb.LoginBarberResponse, error) {
	// Retrieve barber information from the store
	barber, err := server.Store.GetEmailBarber(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			// If barber does not exist, return a 404 error
			return nil, status.Errorf(codes.NotFound, "failed to get email barber %s", err)
		}
		// If there's an error getting the barber, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to get email barber %s", err)
	}

	// Check the password provided against the stored hashed password
	err = utils.CheckPassword(req.Password, barber.HashedPassword)
	if err != nil {
		// If the password check fails, return an unauthenticated error
		return nil, status.Errorf(codes.Unauthenticated, "failed to check the password for %s", err)
	}

	barberPayload := token.BarberPayload{
		BarberID: barber.BarberID,
		Role: barber.Role,
		Phone: barber.Phone,
		Email: barber.Email,
	}
	// Create an access token for the authenticated barber
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		// If there's an error creating the access token, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create token %s", err)
	}

	// Create a refresh token for the authenticated barber
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		// If there's an error creating the refresh token, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create refresh token %s", err)
	}

	mtdt := server.extractMetadata(ctx)
	// Create a session for the barber
	session, err := server.Store.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
		ID:           refreshPayload.ID,
		BarberID:     refreshPayload.Barber.BarberID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		// If there's an error creating the session, return an internal server error
		return nil, status.Errorf(codes.Internal, "failed to create session %s", err)
	}

	// Prepare the response with barber and session information
	rsp := &pb.LoginBarberResponse{
		Barber:                convertBarber(barber),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}

