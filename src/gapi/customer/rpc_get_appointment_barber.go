package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetAppointmentByDate(ctx context.Context, req *customer.GetAppointmentByDateRequest) (*customer.GetAppointmentByDateResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, UnauthenticatedError(err)
	}

	arg := db.GetAppointmentByDateWithServiceParams{
		AppointmentDatetime: req.GetDayAppointment().AsTime(),
		BarberID:            uuid.MustParse(req.BarberId),
	}

	res, err := server.store.GetAppointmentByDateWithService(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update account")
	}

	rsp := &customer.GetAppointmentByDateResponse{
		Appointment: convertGetAppointmentByDate(res),
	}
	return rsp, nil
}
