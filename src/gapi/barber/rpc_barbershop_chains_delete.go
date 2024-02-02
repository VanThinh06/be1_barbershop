package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteBarberShopChains(ctx context.Context, req *barber.DeleteBarberShopChainsRequest) (*barber.DeleteBarberShopChainsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if payload.Barber.BarberID.String() != req.GetId() {
		return nil, unauthenticatedError(err)
	}

	var id = uuid.MustParse(req.Id)
	err = server.Store.DeleteBarberShopChains(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberShopChainsResponse{
		Status: "success",
	}
	return rsp, nil
}
