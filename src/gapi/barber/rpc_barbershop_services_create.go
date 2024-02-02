package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShopServices(ctx context.Context, req *barber.CreateBarberShopServicesRequest) (*barber.CreateBarberShopServicesResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var chainID uuid.NullUUID
	var BarberShopId uuid.NullUUID
	if req.BarberShopChainId != nil {
		chainID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopChainId()),
			Valid: req.BarberShopChainId != nil,
		}
	}
	if req.BarberShopId != nil {
		BarberShopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopId()),
			Valid: req.BarberShopId != nil,
		}
	}
	arg := db.CreateBarberShopServicesParams{
		BarbershopCategoryID: uuid.MustParse(req.BarbershopCategoryId),
		BarberShopChainID:    chainID,
		BarberShopID:         BarberShopId,
		GenderID:             req.GenderId,
		Name:                 req.Name,
		Timer:                req.Timer,
		Price:                req.Price,
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != "",
		},
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
	}

	service, err := server.Store.CreateBarberShopServices(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopServicesResponse{
		BarbershopService: convertBarberShopServices(service),
	}
	return rsp, nil
}
