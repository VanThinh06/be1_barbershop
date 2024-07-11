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

// CreateServiceItem
func (server *Server) CreateServiceItem(ctx context.Context, req *barber.CreateServiceItemRequest) (*barber.CreateServiceItemResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	permissionService := server.checkPermissionManageService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	arg := db.CreateServiceItemParams{
		CategoryID: int16(req.GetCategoryId()),
		GenderID:   int16(req.GetGenderId()),
		Name:       req.GetName(),
		Timer:      int16(req.GetTimer()),
		Price:      req.GetPrice(),
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		ImageUrl: sql.NullString{
			String: req.GetImageUrl(),
			Valid:  req.ImageUrl != nil,
		},
		BarberShopID: barberShopId,
	}
	service, err := server.Store.CreateServiceItem(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "ServiceItems_category_id_fkey":
				return nil, status.Errorf(codes.NotFound, "service category don't exist")
			case "ServiceItems_category_id_name_idx":
				return nil, status.Errorf(codes.AlreadyExists, "service already exists.")
			case "ServiceItems_barber_shop_id":
				return nil, status.Errorf(codes.NotFound, "barbershops don't exist")
			}
		}
		return nil, utils.InternalError(err)
	}

	rsp := &barber.CreateServiceItemResponse{
		Service: &barber.ServiceItem{
			Id:                service.ID.String(),
			CategoryId:        int32(service.CategoryID),
			BarberShopId:      service.BarberShopID.String(),
			GenderId:          int32(service.GenderID),
			Name:              service.Name,
			Timer:             int32(service.Timer),
			Price:             service.Price,
			Description:       service.Description.String,
			ImageUrl:          service.ImageUrl.String,
			DiscountPrice:     service.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(service.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(service.DiscountEndTime.Time),
			IsActive:          service.IsActive,
		},
	}
	return rsp, nil
}

// GetServiceItem
func (server *Server) GetServiceItem(ctx context.Context, req *barber.GetServiceItemRequest) (*barber.GetServiceItemResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetServiceItem(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.GetServiceItemResponse{
		Service: &barber.ServiceItem{
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
			DiscountPrice:     res.DiscountPrice.Float32,
			DiscountStartTime: utils.ConvertToTimestampProto(res.DiscountStartTime),
			DiscountEndTime:   utils.ConvertToTimestampProto(res.DiscountEndTime),
			IsActive:          res.IsActive,
		},
	}
	return rsp, nil
}

// ListServiceItem
func (server *Server) ListServiceItem(ctx context.Context, req *barber.ListServiceItemRequest) (*barber.ListServiceItemResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	permissionService := server.checkPermissionViewService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	res, err := server.Store.ListServicesByCategory(ctx, barberShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListServiceItemResponse{
		Services: convertListService(res),
	}
	return rsp, nil
}

// UpdateServiceItem
func (server *Server) UpdateServiceItem(ctx context.Context, req *barber.UpdateServiceItemRequest) (*barber.UpdateServiceItemResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	permissionService := server.checkPermissionManageService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	idService, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "service don't exist")
	}

	arg := db.UpdateServiceItemParams{
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
			Int16: int16(req.GetTimer()),
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
		ID: idService,
	}
	res, err := server.Store.UpdateServiceItem(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "ServiceItems_category_id_fkey":
				return nil, status.Errorf(codes.NotFound, "service category don't exist")
			case "ServiceItems_category_id_name_idx":
				return nil, status.Errorf(codes.AlreadyExists, "service already exists.")
			case "ServiceItems_barber_shop_id":
				return nil, status.Errorf(codes.NotFound, "barbershops don't exist")
			}
		}
		return nil, utils.InternalError(err)
	}

	rsp := &barber.UpdateServiceItemResponse{
		Service: &barber.ServiceItem{
			Id:                res.ID.String(),
			CategoryId:        int32(res.CategoryID),
			BarberShopId:      res.BarberShopID.String(),
			GenderId:          int32(res.GenderID),
			Name:              res.Name,
			Timer:             int32(res.Timer),
			Price:             res.Price,
			Description:       res.Description.String,
			ImageUrl:          res.ImageUrl.String,
			DiscountPrice:     res.DiscountPrice.Float32,
			DiscountStartTime: utils.ConvertToTimestampProto(res.DiscountStartTime),
			DiscountEndTime:   utils.ConvertToTimestampProto(res.DiscountEndTime),
			IsActive:          res.IsActive,
		},
	}
	return rsp, nil
}

// DeleteServiceItem
func (server *Server) DeleteServiceItem(ctx context.Context, req *barber.DeleteServiceItemRequest) (*barber.DeleteServiceItemResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	idService, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "service don't exist")
	}

	barberShopId, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	permissionService := server.checkPermissionManageService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	err = server.Store.DeleteServiceItem(ctx, idService)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "ServicePackageItems_service_item_id_fkey":
				return nil, status.Errorf(codes.FailedPrecondition, "unable to delete this service because it is still in use elsewhere.")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to delete service item")
	}

	rsp := &barber.DeleteServiceItemResponse{
		Message: "success",
	}
	return rsp, nil
}

func (server *Server) checkPermissionManageService(ctx context.Context, barberShopID uuid.UUID, barberID uuid.UUID) error {
	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ManageService),
		BarberID:     barberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return noPermissionError(err)
	}
	return nil
}

func (server *Server) checkPermissionViewService(ctx context.Context, barberShopID uuid.UUID, barberID uuid.UUID) error {
	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ViewServiceList),
		BarberID:     barberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return noPermissionError(err)
	}
	return nil
}
