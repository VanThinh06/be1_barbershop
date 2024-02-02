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

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var barberShopChainId uuid.NullUUID
	var barberShopID uuid.NullUUID
	if req.BarberShopChainId != nil {
		barberShopChainId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopChainId()),
			Valid: req.BarberShopChainId != nil,
		}
	}
	if req.BarberShopId != nil {
		barberShopID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopId()),
			Valid: req.BarberShopId != nil,
		}
	}

	arg := db.ListBarberShopServiceCategoriesParams{
		BarberShopChainID: barberShopChainId,
		BarberShopID:      barberShopID,
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
