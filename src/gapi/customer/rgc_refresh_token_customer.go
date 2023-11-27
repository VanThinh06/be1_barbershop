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
		return nil, status.Errorf(codes.Unauthenticated, "Unauthenticated", err)
	}

	session, err := server.store.GetSessionsCustomer(ctx, payload.ID)
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

	if session.CustomerID != payload.Customer.CustomerID {
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

	access_token, accessPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Customer,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	refresh_token, refreshPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Customer,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	sessionCustomer, err := server.store.UpdateRefreshTokenSessionsCustomer(ctx, db.UpdateRefreshTokenSessionsCustomerParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error Intenal Server", err)

	}

	rsp := &customer.RefreshTokenCustomerResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          sessionCustomer.RefreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
