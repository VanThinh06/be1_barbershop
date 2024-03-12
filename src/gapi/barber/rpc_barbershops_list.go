package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListBarberShops(ctx context.Context, req *barber.ListBarberShopsRequest) (*barber.ListBarberShopsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	res, err := server.Store.ListBarberShops(ctx, payload.Barber.BarberID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberShopsResponse{
		BarberShops: convertListBarberShops(res),
	}
	return rsp, nil
}
