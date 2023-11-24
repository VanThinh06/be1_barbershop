package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshTokenBarber(ctx context.Context, req *pb.RefreshTokenBarberRequest) (*pb.RefreshTokenBarberResponse, error) {
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)
	}

	session, err := server.Store.GetSessionsBarber(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			if err != nil {
				return nil, status.Errorf(codes.NotFound, "NotFound", err)
			}
		}

		return nil, status.Errorf(codes.Internal, "Unauthenticated", err)

	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")

		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)

	}

	if session.BarberID != payload.Barber.BarberID {
		err := fmt.Errorf("incorrect session user")

		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)

	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("incorrect session user")

		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)

	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")

		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)

	}

	if session.ClientIp != server.extractMetadata(ctx).ClientIP {
		err := fmt.Errorf("incorrect clientIP")
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)
		}
	}
	
	access_token, accessPayload, err := server.tokenMaker.CreateToken(
		payload.Barber,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	refresh_token, refreshPayload, err := server.tokenMaker.CreateToken(
		payload.Barber,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	ssUpdate, err := server.Store.UpdateRefreshTokenSessionsBarber(ctx, db.UpdateRefreshTokenSessionsBarberParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	rsp := &pb.RefreshTokenBarberResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          ssUpdate.RefreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
