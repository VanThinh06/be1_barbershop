package gapi

import (
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) QueryBarberShops(ctx context.Context, req *barber.QueryBarberShopsRequest) (*barber.QueryBarberShopsResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	res, err := server.Store.QueryBarberShops(ctx, req.Barbershop)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.QueryBarberShopsResponse{
		BarberShops: ConvertListBarberShops(res),
	}
	return rsp, nil
}
