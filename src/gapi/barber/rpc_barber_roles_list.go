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
			}
		}
		return nil, internalError(err)
	}

	rsp := &barber.ListBarberRolesResponse{
		BarberRole: ConvertListBarbersRoles(barberRoles),
	}
	return rsp, nil
}
