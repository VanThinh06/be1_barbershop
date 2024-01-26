package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberManagers(ctx context.Context, req *barber.CreateBarberManagersRequest) (*barber.CreateBarberManagersResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.BarberId {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}
	var barberID = uuid.MustParse(req.BarberId)
	var managerID = uuid.MustParse(req.ManagerId)

	arg := db.CreateBarberManagersParams{
		BarberID:  barberID,
		ManagerID: managerID,
	}

	barberManager, err := server.Store.CreateBarberManagers(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberManagersResponse{
		BarberManager: convertBarberManagers(barberManager),
	}
	return rsp, nil
}
