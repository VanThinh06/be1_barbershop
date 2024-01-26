package gapi

import (
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberShops(ctx context.Context, req *barber.GetBarberShopsRequest) (*barber.GetBarberShopsResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	res, err := server.Store.GetBarberShop(ctx, uuid.MustParse(req.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopsResponse{
		BarberShop: convertBarberShops(res),
	}
	return rsp, nil
}
