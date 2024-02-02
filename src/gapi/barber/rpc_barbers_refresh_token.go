package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshTokenBarber(ctx context.Context, req *barber.RefreshTokenBarberRequest) (*barber.RefreshTokenBarberResponse, error) {
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	session, err := server.Store.GetSessionBarber(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			if err != nil {
				return nil, status.Error(codes.NotFound, "notFound")
			}
		}

		return nil, status.Error(codes.Internal, "unauthenticated")
	}

	if session.IsBlocked {
		_ = fmt.Errorf("incorrect block user")
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")

	}

	if session.BarberID != payload.Barber.BarberID {
		_ = fmt.Errorf("incorrect session user")

		return nil, unauthenticatedError(err)

	}

	if session.RefreshToken != req.RefreshToken {
		_ = fmt.Errorf("incorrect session user")

		return nil, status.Error(codes.Unauthenticated, "unauthenticated")

	}

	if time.Now().After(session.ExpiresAt) {
		_ = fmt.Errorf("expired session")
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")

	}

	if session.ClientIp != server.extractMetadata(ctx).ClientIP {
		err := fmt.Errorf("incorrect clientIP")
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated")
		}
	}

	access_token, accessPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Barber,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	refresh_token, refreshPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Barber,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	arg := db.UpdateSessionRefreshTokenParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
	}
	err = server.Store.UpdateSessionRefreshToken(ctx, arg)
	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	rsp := &barber.RefreshTokenBarberResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
		RefreshToken:          refresh_token,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}
	return rsp, nil
}
