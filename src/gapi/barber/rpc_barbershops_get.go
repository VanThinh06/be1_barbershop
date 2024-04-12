package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberShop(ctx context.Context, req *barber.GetBarberShopRequest) (*barber.GetBarberShopResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	arg := db.GetBarberShopParams{
		ID:       barberShopId,
		BarberID: payload.Barber.BarberID,
	}
	res, err := server.Store.GetBarberShop(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopResponse{
		BarberShop: convertGetBarberShop(res),
	}
	return rsp, nil
}
