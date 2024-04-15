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

func (server *Server) GetBarberEmployees(ctx context.Context, req *barber.GetBarberEmployeeRequest) (*barber.GetBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ViewEmployeeInformation),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, noPermissionError(err)
	}

	arg := db.ListEmployeesParams{
		BarberShopID: barberShopID,
		Limit:        req.Limit,
		Offset:       req.Limit * (req.Page - 1),
	}
	barberEmployees, err := server.Store.ListEmployees(ctx, arg)
	if err != nil {
		return nil, internalError(err)
	}

	var totalEmployees int16 = 0
	if len(barberEmployees) > 0 {
		totalEmployees = int16(barberEmployees[0].TotalEmployees)
	}
	rsp := &barber.GetBarberEmployeeResponse{
		BarberDetails:  convertBarberEmployees(barberEmployees),
		TotalEmployees: int32(totalEmployees),
	}
	return rsp, nil
}
