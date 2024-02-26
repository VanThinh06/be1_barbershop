package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShopServices(ctx context.Context, req *barber.CreateBarberShopServicesRequest) (*barber.CreateBarberShopServicesResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var BarberShopId uuid.NullUUID
	if req.BarberShopId != nil {
		BarberShopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopId()),
			Valid: req.BarberShopId != nil,
		}
	}


	service, err := server.Store.CreateBarberShopServices(ctx, BarberShopId.UUID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberShopServicesResponse{
		BarbershopService: convertBarberShopServices(service),
	}
	return rsp, nil
}
