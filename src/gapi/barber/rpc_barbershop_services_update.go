package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarberShopService(ctx context.Context, req *barber.UpdateBarberShopServiceRequest) (*barber.UpdateBarberShopServiceResponse, error) {

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
	if len(req.GetComboServices()) != 0 {
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

	idService, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "service don't exist")
	}

	arg := db.UpdateBarberShopServiceParams{
		CategoryID: pgtype.Int2{
			Int16: int16(req.GetCategoryId()),
			Valid: req.CategoryId != nil,
		},
		GenderID: pgtype.Int2{
			Int16: int16(req.GetGenderId()),
			Valid: req.GenderId != nil,
		},
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		IsActive: pgtype.Bool{
			Bool:  req.GetIsActive(),
			Valid: req.IsActive != nil,
		},
		Timer: pgtype.Int2{
			Int16: int16(timer),
			Valid: req.Timer != nil,
		},
		Price: pgtype.Float4{
			Float32: req.GetPrice(),
			Valid:   req.Price != nil,
		},
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		ImageUrl: sql.NullString{
			String: req.GetImageUrl(),
			Valid:  req.ImageUrl != nil,
		},
		DiscountPrice:  pgtype.Float4{
			Float32: req.GetDiscountPrice(),
			Valid:   req.DiscountPrice != nil,
		},
		DiscountStartTime: pgtype.Timestamp{
			Time: req.GetDiscountStartTime().AsTime(),
			Valid: req.DiscountStartTime != nil,
		},
		DiscountEndTime: pgtype.Timestamp{
			Time: req.GetDiscountEndTime().AsTime(),
			Valid: req.DiscountEndTime != nil,
		},
		ComboServices: req.ComboServices,
		ID:            idService,
	}
	err = server.Store.UpdateBarberShopService(ctx, arg)
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

	rsp := &barber.UpdateBarberShopServiceResponse{
		Message: "the service has been updated successfully.",
	}
	return rsp, nil
}
