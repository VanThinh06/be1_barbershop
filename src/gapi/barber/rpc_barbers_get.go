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

func (server *Server) GetBarber(ctx context.Context, req *barber.GetBarbersRequest) (*barber.GetBarbersResponse, error) {
	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	idBarber, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barber don't exist")
	}
	idBarberShop, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershop don't exist")
	}
	
	arg := db.GetBarberParams{
		ID:           idBarber,
		BarberShopID: idBarberShop,
	}
	res, err := server.Store.GetBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
		return nil, internalError(err)
	}

	rsp := &barber.GetBarbersResponse{
		BarberDetail : convertBarberEmployee(res),
	}
	return rsp, nil
}
