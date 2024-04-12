package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberEmployees(ctx context.Context, req *barber.GetBarberEmployeeRequest) (*barber.GetBarberEmployeeResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	// if payload.Barber.BarberRole != int32(utilities.Admin) {
	// 	return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	// }

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	// check admin role
	arg := db.ListEmployeesAdminParams{
		BarberShopID: barberShopID,
		Limit:        req.Limit,
		Offset:       req.Limit * (req.Page - 1),
	}
	barberEmployees, err := server.Store.ListEmployeesAdmin(ctx, arg)
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
