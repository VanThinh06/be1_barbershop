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

func (server *Server) CreateBarberShops(ctx context.Context, req *barber.CreateBarberShopsRequest) (*barber.CreateBarberShopsResponse, error) {

	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var barberShopChainId = uuid.NullUUID{
		UUID:  uuid.MustParse(req.GetBarbershopChainId()),
		Valid: req.BarbershopChainId != nil,
	}
	arg := db.CreateBarberShopParams{
		Name:              req.Name,
		BarbershopChainID: barberShopChainId,
		IsMainBranch: sql.NullBool{
			Bool:  req.GetIsMainBranch(),
			Valid: req.IsMainBranch != nil,
		},
		BranchCount: req.BranchCount,
		Longitude:   req.Longitude.Value,
		Latitude:    req.Latitude.Value,
		Address:     req.Address,
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
	}

	barberShop, err := server.Store.CreateBarberShop(ctx, arg)
	if err != nil {

		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopsResponse{
		BarberShop: convertBarberShops(barberShop),
	}
	return rsp, nil
}
