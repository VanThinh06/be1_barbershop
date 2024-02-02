package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShopChains(ctx context.Context, req *barber.CreateBarberShopChainsRequest) (*barber.CreateBarberShopChainsResponse, error) {

	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.CreateBarberShopChainsParams{
		Name: req.Name,
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		Founder:      req.Founder,
		FoundingDate: req.FoundingDate.AsTime(),
		Website: sql.NullString{
			String: req.GetWebsite(),
			Valid:  req.Website != "",
		},
	}

	barberShopChain, err := server.Store.CreateBarberShopChains(ctx, arg)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopChainsResponse{
		BarbershopChain: ConvertBarberShopChains(barberShopChain),
	}
	return rsp, nil
}
