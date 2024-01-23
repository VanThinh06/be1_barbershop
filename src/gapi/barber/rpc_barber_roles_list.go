package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ListBarberRoles(ctx context.Context, req *barber.ListBarberRolesRequest) (*barber.ListBarberRolesResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var barberShopID = uuid.MustParse(req.BarbershopId)

	barberRoles, err := server.Store.ListBarbersRoles(ctx, barberShopID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Barbershop_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberRolesResponse{
		BarberRoles: ConvertListBarbersRoles(barberRoles),
	}
	return rsp, nil
}
