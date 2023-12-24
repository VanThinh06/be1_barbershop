package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShop(ctx context.Context, req *barber.CreateBarberShopRequest) (*barber.CreateBarberShopResponse, error) {

	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var chainId uuid.NullUUID
	var facility = req.Facility
	if req.ChainId != nil {
		chainId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetChainId()),
			Valid: req.ChainId != nil,
		}
		facility = req.GetFacility()
	}
	arg := db.CreateBarberShopParams{
		OwnerID:   authPayload.Barber.BarberID,
		Name:      req.GetName(),
		ChainID:   chainId,
		Longitude: req.Longitude.Value,
		Latitude:  req.Latitude.Value,
		Address:   req.GetAddress(),
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
		Facility: facility,
	}

	isBarberShopInChain, err := server.Store.BarberShopInChain(ctx, authPayload.Barber.BarberID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}
	if isBarberShopInChain {
		return nil, status.Errorf(codes.PermissionDenied, "account cannot create many other barber shops")
	}
	
	barberShop, err := server.Store.CreateBarberShop(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.ConstraintName {
			case "BarberShops_pkey":
				return nil, status.Errorf(codes.AlreadyExists, "barber shop already existed")

			case "BarberShops_chain_id_facility_idx":
				return nil, status.Errorf(codes.AlreadyExists, "barber shop already existed")

			case "BarberShops_owner_id_facility_idx":
				return nil, status.Errorf(codes.AlreadyExists, "barber shop already existed")
			}
		}
		return nil, status.Error(codes.Internal, "internal")
	}

	if req.IsMainBranch {
		requestBarber := db.UpdateBarberParams{BarberID: authPayload.Barber.BarberID, ShopID: uuid.NullUUID{UUID: barberShop.ShopID, Valid: true}}
		_, errUpdateBarber := server.Store.UpdateBarber(ctx, requestBarber)
		if errUpdateBarber != nil {
			return nil, status.Errorf(codes.Internal, "internal")
		}
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopResponse{
		BarberShop: convertBarberShop(barberShop),
	}
	return rsp, nil
}
