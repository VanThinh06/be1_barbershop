package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SearchByNameBarberShops(ctx context.Context, req *barber.SearchByNameBarberShopsRequest) (*barber.SearchByNameBarberShopsResponse, error) {
	_, err := server.authorizeBarber(ctx)
	if err != nil {
		_, err = server.authorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}
	arg := db.SearchByNameBarberShopsParams{
		CurrentLongitude: req.GetLongitude().GetValue(),
		CurrentLatitude:  req.GetLatitude().GetValue(),
		Name:             req.GetName(),
	}
	res, err := server.Store.SearchByNameBarberShops(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.SearchByNameBarberShopsResponse{
		BarberShops: ConvertSearchByNameBarberShops(res),
	}
	return rsp, nil
}
