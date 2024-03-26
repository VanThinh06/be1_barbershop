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

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	idBarberShop, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetBarberShop(ctx, idBarberShop)
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
