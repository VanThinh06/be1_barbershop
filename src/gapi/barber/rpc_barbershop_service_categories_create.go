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

func (server *Server) CreateBarberShopServiceCategories(ctx context.Context, req *barber.CreateBarberShopServiceCategoriesRequest) (*barber.CreateBarberShopServiceCategoriesResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	var chainID uuid.NullUUID
	var barberShopID uuid.NullUUID
	if req.BarbershopChainId != nil {
		chainID = uuid.NullUUID{
			UUID: uuid.MustParse(req.GetBarbershopChainId()),
			Valid: req.BarbershopChainId != nil,
		}
	}
	if req.BarbershopId != nil {
		barberShopID = uuid.NullUUID{
			UUID: uuid.MustParse(req.GetBarbershopId()),
			Valid: req.BarbershopId != nil,
		}
	}

	arg := db.CreateBarberShopServiceCategoriesParams{
		BarbershopChainID: chainID,
		BarbershopID: barberShopID,
	}

	serviceCategory, err := server.Store.CreateBarberShopServiceCategories(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateBarberShopServiceCategoriesResponse{
		ServiceCategory: convertBSServiceCategories(serviceCategory),
	}
	return rsp, nil
}
