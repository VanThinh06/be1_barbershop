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

func (server *Server) CreateBarberRoles(ctx context.Context, req *barber.CreateBarberRolesRequest) (*barber.CreateBarberRolesResponse, error) {

	payload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.BarberId {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}
	var barberID = uuid.MustParse(req.BarberId)
	var barberShopID = uuid.MustParse(req.BarbershopId)
	arg := db.CreateBarberRolesParams{
		BarberID:     barberID,
		BarbershopID: barberShopID,
		RoleID:       req.RoleId,
	}

	barberRole, err := server.Store.CreateBarberRoles(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Chains_owner_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberRolesResponse{
		BarberRole: ConvertBarberRoles(barberRole),
	}
	return rsp, nil
}
