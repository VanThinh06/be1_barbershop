package gapi

import (
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

	if payload.Barber.BarberRoleType != string(utilities.Administrator) {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	barberEmployees, err := server.Store.GetBarberEmployees(ctx, barberShopID)
	if err != nil {
		return nil, internalError(err)
	}

	rsp := &barber.GetBarberEmployeeResponse{
		BarberEmployees: convertBarberEmployees(barberEmployees),
	}
	return rsp, nil
}
