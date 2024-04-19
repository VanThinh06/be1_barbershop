package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShopService(ctx context.Context, req *barber.CreateBarberShopServiceRequest) (*barber.CreateBarberShopServiceResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.GetBarberShopId())
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

	// check combo services
	var timer = int64(req.GetTimer())
	if len(req.ComboServices) != 0 {
		var listComboServices []uuid.UUID
		for _, item := range req.ComboServices {
			uuidService, err := uuid.Parse(item)
			if err != nil {
				return nil, status.Errorf(codes.NotFound, "service don't exist")
			}
			listComboServices = append(listComboServices, uuidService)
		}
		timer, err = server.Store.GetTimerBarberShopServices(ctx, listComboServices)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal")
		}
	}

	arg := db.CreateBarberShopServiceParams{
		CategoryID: int16(req.GetCategoryId()),
		GenderID:   int16(req.GetGenderId()),
		Name:       req.GetName(),
		Timer:      int16(timer),
		Price:      req.GetPrice(),
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		ImageUrl: sql.NullString{
			String: req.GetImageUrl(),
			Valid:  req.ImageUrl != nil,
		},
		BarberShopID:  barberShopId,
		ComboServices: req.ComboServices,
	}
	service, err := server.Store.CreateBarberShopService(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "BarberShopServices_category_id_fkey":
				return nil, status.Errorf(codes.NotFound, "service category don't exist")
			case "BarberShopServices_category_id_name_idx":
				return nil, status.Errorf(codes.AlreadyExists, "service already exists.")
			case "BarberShopServices_barber_shop_id":
				return nil, status.Errorf(codes.NotFound, "barbershops don't exist")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopServiceResponse{
		BarberShopService: convertBarberShopService(service),
	}
	return rsp, nil
}
