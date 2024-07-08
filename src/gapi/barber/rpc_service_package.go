package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/utils"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateServicePackage
func (server *Server) CreateServicePackage(ctx context.Context, req *barber.CreateServicePackageRequest) (*barber.CreateServicePackageResponse, error) {

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
	if len(req.GetServiceItems()) != 0 {
		var listServices []uuid.UUID
		for _, item := range req.GetServiceItems() {
			uuidService, err := uuid.Parse(item)
			if err != nil {
				return nil, status.Errorf(codes.NotFound, "service don't exist")
			}
			listServices = append(listServices, uuidService)
		}
		timer, err = server.Store.GetTimerServiceItem(ctx, listServices)
		if err != nil {
			return nil, utils.InternalError(err)
		}
	}

	var createdService db.ServicePackage
	err = server.Store.ExecTx(ctx, func(q *db.Queries) error {
		arg := db.CreateServicePackageParams{
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
		service, err := server.Store.CreateServicePackage(ctx, arg)
		if err != nil {
			return utils.InternalError(err)
		}

		createdService = service
		for _, v := range req.GetServiceItems() {
			barberShopServiceID, err := uuid.Parse(v)
			if err != nil {
				return status.Errorf(codes.InvalidArgument, "service don't exist")
			}
			arg := db.CreateServicePackageItemParams{
				ServicePackageID: service.ID,
				ServiceItemID:    barberShopServiceID,
			}
			_, err = server.Store.CreateServicePackageItem(ctx, arg)
			if err != nil {
				return utils.InternalError(err)
			}
		}
		return nil
	})

	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			}
		}
		return nil, utils.InternalError(err)
	}

	res, err := server.Store.GetServicePackage(ctx, createdService.ID)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.CreateServicePackageResponse{
		ServicePackage: convertServicePackageItem(res),
	}

	return rsp, nil
}

// GetServicePackage
func (server *Server) GetServicePackage(ctx context.Context, req *barber.GetServicePackageRequest) (*barber.GetServicePackageResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "service don't exist")
	}

	res, err := server.Store.GetServicePackage(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.GetServicePackageResponse{
		ServicePackageItem: convertServicePackageItem(res),
	}

	return rsp, nil
}

// ListServicePackage
func (server *Server) ListServicePackage(ctx context.Context, req *barber.ListServicePackageRequest) (*barber.ListServicePackageResponse, error) {

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

	res, err := server.Store.ListServicePackages(ctx, barberShopId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListServicePackageResponse{
		ServicePackages: convertListServiceItem(res),
	}
	return rsp, nil
}

// UpdateServicePackage
func (server *Server) UpdateServicePackage(ctx context.Context, req *barber.UpdateServicePackageRequest) (*barber.UpdateServicePackageResponse, error) {
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

	idServicePackage, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "service don't exist")
	}

	var listServiceItems []uuid.UUID
	for _, item := range req.GetServiceItems() {
		uuidServiceItem, err := uuid.Parse(item)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "service don't exist")
		}
		listServiceItems = append(listServiceItems, uuidServiceItem)
	}

	arg := db.UpdateServicePackageParams{
		GenderID: int16(req.GetGenderId()),
		Name:     req.GetName(),
		IsActive: req.GetIsActive(),
		Timer:    int16(req.GetTimer()),
		Price:    req.GetPrice(),
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
		ID: idServicePackage,
	}

	argServicePackageItem := db.UpsertServicePackageItemParams{
		ServicePackageID: idServicePackage,
		Column2:          listServiceItems,
	}

	var servicePackage db.ServicePackage
	err = server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		servicePackage, err = q.UpdateServicePackage(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pgconn.PgError); ok {
				switch pqErr.ConstraintName {
				}
			}
			return err
		}

		err = q.UpsertServicePackageItem(ctx, argServicePackageItem)
		if err != nil {
			if pqErr, ok := err.(*pgconn.PgError); ok {
				switch pqErr.ConstraintName {
				}
			}
			return err
		}

		return nil
	})

	if err != nil {
		return nil, utils.InternalError(err)
	}

	res, err := server.Store.GetServicePackage(ctx, servicePackage.ID)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.UpdateServicePackageResponse{
		ServicePackage: convertServicePackageItem(res),
	}
	return rsp, nil
}
