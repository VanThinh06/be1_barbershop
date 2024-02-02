package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberManagers(ctx context.Context, req *barber.GetBarberManagersRequest) (*barber.GetBarberManagersResponse, error) {

	payload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.Id {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}

	var barberID = uuid.MustParse(req.Id)
	barberManager, err := server.Store.GetBarberManagers(ctx, barberID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberManagersResponse{
		BarberManager: convertBarberManagers(barberManager),
	}
	return rsp, nil
}
