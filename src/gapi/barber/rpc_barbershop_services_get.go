package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberShopServices(ctx context.Context, req *barber.GetBarberShopServicesRequest) (*barber.GetBarberShopServicesResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetBarberShopService(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopServicesResponse{
		BarberShopService: convertBarberShopService(res),
	}
	return rsp, nil
}
