package gapi

// import (
// 	db "barbershop/src/db/sqlc"
// 	"barbershop/src/pb/barber"
// 	"context"

// 	"github.com/google/uuid"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func (server *Server) ListAppointmentsByDate(ctx context.Context, req *barber.ListAppointmentsByDateRequest) (*barber.ListAppointmentsByDateResponse, error) {
// 	_, err := server.authorizeBarber(ctx)
// 	if err != nil {
// 		return nil, unauthenticatedError(err)
// 	}

// 	arg := db.ListAppointmentsByDateParams{
// 		BarberID:            uuid.MustParse(req.BarberId),
// 		AppointmentDateTime: req.AppointmentDatetime.AsTime(),
// 	}

// 	res, err := server.Store.ListAppointmentsByDate(ctx, arg)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to update account")
// 	}

// 	rsp := &barber.ListAppointmentsByDateResponse{
// 		Appointment: convertListAppointmentsByDate(res),
// 	}
// 	return rsp, nil
// }
