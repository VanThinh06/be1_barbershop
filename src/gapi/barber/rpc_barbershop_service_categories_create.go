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

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}


	var barberShopID uuid.NullUUID
	if req.BarberShopId != nil {
		barberShopID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopId()),
			Valid: req.BarberShopId != nil,
		}
	}

	arg := db.CreateBarberShopServiceCategoriesParams{
		BarberShopID:      barberShopID.UUID,
	}

	serviceCategory, err := server.Store.CreateBarberShopServiceCategories(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateBarberShopServiceCategoriesResponse{
		ServiceCategory: convertBSServiceCategories(serviceCategory),
	}
	return rsp, nil
}
