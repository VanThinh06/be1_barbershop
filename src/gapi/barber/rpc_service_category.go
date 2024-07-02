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
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateServiceCategory
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
				return nil, status.Errorf(codes.AlreadyExists, "service category already exists.")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateServiceCategoryResponse{
		ServiceCategory: convertServiceCategory(serviceCategory),
	}
	return rsp, nil
}

// ListServiceCategory
func (server *Server) ListServiceCategory(ctx context.Context, req *barber.ListServiceCategoryRequest) (*barber.ListServiceCategoryResponse, error) {

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

	rsp := &barber.ListServiceCategoryResponse{
		ServiceCategories: convertListServiceCategory(res),
	}
	return rsp, nil
}

// UpdateServiceCategory
func (server *Server) UpdateServiceCategory(ctx context.Context, req *barber.UpdateServiceCategoryRequest) (*barber.UpdateServiceCategoryResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.UpdateServiceCategoryParams{
		ID: int16(req.Id),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
	}

	err = server.Store.UpdateServiceCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UpdateServiceCategoryResponse{
		Message: "Message",
	}
	return rsp, nil
}

// UpdateCategoryPosition
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

// DeleteServiceCategory
func (server *Server) DeleteServiceCategory(ctx context.Context, req *barber.DeleteServiceCategoryRequest) (*barber.DeleteServiceCategoryResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.GetId() {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}

	err = server.Store.DeleteServiceCategories(ctx, 1)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteServiceCategoryResponse{
		Message: "success",
	}
	return rsp, nil
}
