package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"context"
	"database/sql"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshTokenCustomer(ctx context.Context, req *customer.RefreshTokenCustomerRequest) (*customer.RefreshTokenCustomerResponse, error) {
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	session, err := server.store.GetSessionsCustomer(ctx, payload.ID)
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

	if session.CustomerID != payload.Customer.CustomerID {
		_ = fmt.Errorf("incorrect session user")

		return nil, status.Error(codes.Unauthenticated, "unauthenticated")

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
		payload.Customer,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	refresh_token, refreshPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Customer,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	sessionCustomer, err := server.store.UpdateRefreshTokenSessionsCustomer(ctx, db.UpdateRefreshTokenSessionsCustomerParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "intenal")
	}

	rsp := &customer.RefreshTokenCustomerResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          sessionCustomer.RefreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
