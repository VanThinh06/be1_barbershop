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

func (server *Server) GetBarberShopServices(ctx context.Context, req *barber.GetBarberRolesRequest) (*barber.GetBarberRolesResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var barberID = uuid.MustParse(req.BarberId)
	var barberShopID = uuid.MustParse(req.BarbershopId)
	arg := db.GetBarberRolesParams{
		BarberID:     barberID,
		BarbershopID: barberShopID,
	}

	barberRoles, err := server.Store.GetBarberRoles(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Barber_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			case "Barbershop_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopServicesResponse{
		BarberShopService: ConvertBarberRolesDetail(barberRoles),
	}
	return rsp, nil
}
