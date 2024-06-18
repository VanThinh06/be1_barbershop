package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"barbershop/src/utils"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
			return nil, utils.InternalError(err)
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
		return nil, utils.InternalError(err)
	}

	rsp := &barber.CreateBarberShopServiceResponse{
		BarberShopService: &barber.BarberShopService{
			Id:                service.ID.String(),
			CategoryId:        int32(service.CategoryID),
			BarberShopId:      service.BarberShopID.String(),
			GenderId:          int32(service.GenderID),
			Name:              service.Name,
			Timer:             int32(service.Timer),
			Price:             service.Price,
			Description:       service.Description.String,
			ImageUrl:          service.ImageUrl.String,
			ComboServices:     service.ComboServices,
			DiscountPrice:     &service.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(service.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(service.DiscountEndTime.Time),
			IsActive:          service.IsActive,
		},
	}
	return rsp, nil
}

// service detail
func (server *Server) GetBarberShopService(ctx context.Context, req *barber.GetBarberShopServiceRequest) (*barber.GetBarberShopServiceResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetBarberShopService(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	discountStartTime := timestamppb.New(res.DiscountStartTime.Time)
	if !res.DiscountStartTime.Valid {
		discountStartTime = nil
	}

	discountEndTime := timestamppb.New(res.DiscountEndTime.Time)
	if !res.DiscountEndTime.Valid {
		discountEndTime = nil
	}

	rsp := &barber.GetBarberShopServiceResponse{
		BarberShopService: &barber.BarberShopService{
			Id:                res.ID.String(),
			CategoryId:        int32(res.CategoryID),
			CategoryName:      res.CategoryName.String,
			BarberShopId:      res.BarberShopID.String(),
			GenderId:          int32(res.GenderID),
			Name:              res.Name,
			Timer:             int32(res.Timer),
			Price:             res.Price,
			Description:       res.Description.String,
			ImageUrl:          res.ImageUrl.String,
			ComboServices:     res.ComboServices,
			DiscountPrice:     &res.DiscountPrice.Float32,
			DiscountStartTime: discountStartTime,
			DiscountEndTime:   discountEndTime,
			IsActive:          res.IsActive,
		},
	}
	return rsp, nil
}

// service list
func (server *Server) ListBarberShopService(ctx context.Context, req *barber.ListBarberShopServiceRequest) (*barber.ListBarberShopServiceResponse, error) {

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

	res, err := server.Store.ListServicesByCategory(ctx, barberShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberShopServiceResponse{
		BarberShopServices: convertListService(res),
	}
	return rsp, nil
}


func (server *Server) ListServiceForComboService(ctx context.Context, req *barber.ListServiceForComboServiceRequest) (*barber.ListServiceForComboServiceResponse, error) {

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

	res, err := server.Store.ListServiceForComboService(ctx, barberShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListServiceForComboServiceResponse{
		BarberShopServices: convertListServiceForComboService(res),
	}
	return rsp, nil
}


// service list combo
func (server *Server) ListComboService(ctx context.Context, req *barber.ListBarberShopServiceRequest) (*barber.ListBarberShopServiceResponse, error) {

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

	res, err := server.Store.ListComboServices(ctx, barberShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberShopServiceResponse{
		BarberShopServices: convertListComboService(res),
	}
	return rsp, nil
}

// update
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
		DiscountPrice: pgtype.Float4{
			Float32: req.GetDiscountPrice(),
			Valid:   req.DiscountPrice != nil,
		},
		DiscountStartTime: pgtype.Timestamp{
			Time:  req.GetDiscountStartTime().AsTime(),
			Valid: req.DiscountStartTime != nil,
		},
		DiscountEndTime: pgtype.Timestamp{
			Time:  req.GetDiscountEndTime().AsTime(),
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

// delete
func (server *Server) DeleteBarberShopService(ctx context.Context, req *barber.DeleteBarberShopServiceRequest) (*barber.DeleteBarberShopServiceResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.GetId() {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}

	var id = uuid.MustParse(req.Id)
	err = server.Store.DeleteBarberShopChains(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberShopServiceResponse{
		Status: "success",
	}
	return rsp, nil
}
