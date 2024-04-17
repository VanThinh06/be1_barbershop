package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListServiceCategories(ctx context.Context, req *barber.ListServiceCategoriesRequest) (*barber.ListServiceCategoriesResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ViewServiceList),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopId,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, noPermissionError(err)
	}

	res, err := server.Store.ListServiceCategories(ctx, uuid.NullUUID{
		UUID:  barberShopId,
		Valid: true,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListServiceCategoriesResponse{
		ServiceCategories: convertListServiceCategories(res),
	}
	return rsp, nil
}
