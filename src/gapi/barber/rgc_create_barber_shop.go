package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) NewBarberShops(ctx context.Context, req *barber.CreateBarberShopsRequest) (*barber.CreateBarberShopsResponse, error) {

	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	coordinates := fmt.Sprintf("POINT(%.7f %.7f)", req.Longitude, req.Latitude)

	arg := db.CreateBarberShopsParams{
		CodeBarberShop: req.GetCodeBarberShop(),
		OwnerID:        authPayload.Barber.BarberID,
		Name:           req.GetName(),
		Coordinates:    coordinates,
		Address:        req.GetAddress(),
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
	}

	barberShop, err := server.Store.CreateBarberShops(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.ConstraintName {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop ", err)
	}

	requestBarber := db.UpdateBarberParams{BarberID: authPayload.Barber.BarberID, ShopID: uuid.NullUUID{UUID: barberShop.ShopID, Valid: true}}
	_, errUpdateBarber := server.Store.UpdateBarber(ctx, requestBarber)
	if errUpdateBarber != nil {
		return nil, status.Errorf(codes.Internal, "Intenal")
	}

	rsp := &barber.CreateBarberShopsResponse{
		BarberShop: convertBarberShop(barberShop),
	}
	return rsp, nil
}
