package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateServiceCategory(ctx context.Context, req *barber.CreateServiceCategoryRequest) (*barber.CreateServiceCategoryResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ManageService),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopId,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, noPermissionError(err)
	}

	arg := db.CreateServiceCategoryParams{
		BarberShopID: uuid.NullUUID{
			UUID:  barberShopId,
			Valid: true,
		},
		Name: req.Name,
	}

	serviceCategory, err := server.Store.CreateServiceCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "ServiceCategories_name_key":
				return nil, status.Errorf(codes.Internal, "service category already exists.")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateServiceCategoryResponse{
		ServiceCategory: convertBSServiceCategories(serviceCategory),
	}
	return rsp, nil
}
