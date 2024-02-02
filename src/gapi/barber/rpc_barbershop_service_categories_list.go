package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListBarberShopServiceCategories(ctx context.Context, req *barber.ListBarberShopServiceCategoriesRequest) (*barber.ListBarberShopServiceCategoriesResponse, error) {

	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var chainID uuid.NullUUID
	var barberShopID uuid.NullUUID
	if req.BarbershopChainId != nil {
		chainID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarbershopChainId()),
			Valid: req.BarbershopChainId != nil,
		}
	}
	if req.BarbershopId != nil {
		barberShopID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarbershopId()),
			Valid: req.BarbershopId != nil,
		}
	}

	arg := db.ListBarberShopServiceCategoriesParams{
		BarbershopChainID: chainID,
		BarbershopID:      barberShopID,
	}

	serviceCategories, err := server.Store.ListBarberShopServiceCategories(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberShopServiceCategoriesResponse{
		ServiceCategories: convertListBSServiceCategories(serviceCategories),
	}
	return rsp, nil
}
