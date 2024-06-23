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

	permissionService := server.checkPermissionManageService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	arg := db.CreateBarberShopServiceParams{
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
			DiscountPrice:     &service.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(service.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(service.DiscountEndTime.Time),
			IsActive:          service.IsActive,
		},
	}
	return rsp, nil
}

func (server *Server) CreateComboService(ctx context.Context, req *barber.CreateComboServiceRequest) (*barber.CreateComboServiceResponse, error) {

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

	var timer = int64(req.GetTimer())
	if len(req.GetServices()) != 0 {
		var listServices []uuid.UUID
		for _, item := range req.GetServices() {
			uuidService, err := uuid.Parse(item)
			if err != nil {
				return nil, status.Errorf(codes.NotFound, "service don't exist")
			}
			listServices = append(listServices, uuidService)
		}
		timer, err = server.Store.GetTimerBarberShopServices(ctx, listServices)
		if err != nil {
			return nil, utils.InternalError(err)
		}
	}

	var createdService db.ComboService
	err = server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateComboServicesParams{
			GenderID: int16(req.GetGenderId()),
			Name:     req.GetName(),
			Timer:    int16(timer),
			Price:    req.GetPrice(),
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
		service, err := server.Store.CreateComboServices(ctx, arg)
		if err != nil {
			return utils.InternalError(err)
		}

		createdService = service
		for _, v := range req.GetServices() {
			barberShopServiceID, err := uuid.Parse(v)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "barbershops don't exist")
			}
			arg := db.CreateComboServiceItemsParams{
				ComboServiceID:      service.ID,
				BarberShopServiceID: barberShopServiceID,
			}
			_, err = server.Store.CreateComboServiceItems(ctx, arg)
			if err != nil {
				return utils.InternalError(err)
			}
		}
		return nil
	})

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

	rsp := &barber.CreateComboServiceResponse{
		Service: &barber.ComboService{
			Id:           createdService.ID.String(),
			BarberShopId: createdService.BarberShopID.String(),
			GenderId:     int32(createdService.GenderID),
			Name:         createdService.Name,
			Timer:        int32(createdService.Timer),
			Price:        createdService.Price,
			Description:  createdService.Description.String,
			ImageUrl:     createdService.ImageUrl.String,
			IsActive:     createdService.IsActive,
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
			DiscountPrice:     &res.DiscountPrice.Float32,
			DiscountStartTime: discountStartTime,
			DiscountEndTime:   discountEndTime,
			IsActive:          res.IsActive,
		},
	}
	return rsp, nil
}

// combo service detail
func (server *Server) GetComboService(ctx context.Context, req *barber.GetComboServiceRequest) (*barber.GetComboServiceResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetComboService(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.GetComboServiceResponse{
		BarberShopService: &barber.ComboService{
			Id:           res.ComboServiceID.String(),
			BarberShopId: res.BarberShopServiceID.String(),
			Name:         res.BarberShopServiceName,
			
			GenderId:     int32(res.GenderID),
			Timer:        int32(res.Timer),
			Price:        res.Price,
			Description:  res.Description.String,
			ImageUrl:     res.BarberShopServic,

			IsActive: res.IsActive,
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

	permissionService := server.checkPermissionViewService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
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

	permissionService := server.checkPermissionViewService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
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

	permissionService := server.checkPermissionManageService(ctx, barberShopId, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
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
			return nil, utils.InternalError(err)
		}
	}

	idService, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "service don't exist")
	}

	arg := db.UpdateBarberShopServiceParams{
		CategoryID: int16(req.GetCategoryId()),
		GenderID:   int16(req.GetGenderId()),
		Name:       req.GetName(),
		IsActive:   req.GetIsActive(),
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
		return nil, utils.InternalError(err)
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
		Message: "success",
	}
	return rsp, nil
}

// delete
func (server *Server) UpdateCategoryPosition(ctx context.Context, req *barber.UpdateCategoryPositionRequest) (*barber.UpdateCategoryPositionResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	permissionService := server.checkPermissionManageService(ctx, barberShopID, payload.Barber.BarberID)
	if permissionService != nil {
		return nil, permissionService
	}

	err = server.Store.ExecTx(ctx, func(q *db.Queries) error {
		for _, v := range req.CategoryPositions {
			arg := db.UpdateCategoryPositionParams{
				BarberShopID: barberShopID,
				CategoryID:   int16(v.CategoryId),
				Position:     int16(v.Position),
				Visible:      v.Visible,
			}

			err := q.UpdateCategoryPosition(ctx, arg)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.UpdateCategoryPositionResponse{
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
