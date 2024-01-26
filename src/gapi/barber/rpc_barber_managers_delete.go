package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteBarberManagers(ctx context.Context, req *barber.DeleteBarberManagersRequest) (*barber.DeleteBarberManagersResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	
	var barberID = uuid.MustParse(req.BarberId)
	var managerID = uuid.MustParse(req.BarberId)

	arg :=  db.DeleteBarberManagersParams{
		BarberID:  barberID,
		ManagerID: managerID,
	}
	err = server.Store.DeleteBarberManagers(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberManagersResponse{
		Status: "success",
	}
	return rsp, nil
}
