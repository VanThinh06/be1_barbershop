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

func (server *Server) UpdateBarberRoles(ctx context.Context, req *barber.UpdateBarberRolesRequest) (*barber.UpdateBarberRolesResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if payload.Barber.BarberID.String() != req.Id {
		return nil, noPermissionError(err)
	}
	var barberRoleID = uuid.MustParse(req.Id)
	arg := db.UpdateBarberRolesParams{
		RoleID: int16(req.RoleId),
		ID:     barberRoleID,
	}

	barberRole, err := server.Store.UpdateBarberRoles(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Barber_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			case "barber_shop_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UpdateBarberRolesResponse{
		BarberRole: ConvertBarberRoles(barberRole),
	}
	return rsp, nil
}
