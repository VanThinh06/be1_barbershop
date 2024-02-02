package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberShopChains(ctx context.Context, req *barber.GetBarberShopChainsRequest) (*barber.GetBarberShopChainsResponse, error) {

	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var idBarberShopChain = uuid.MustParse(req.Id)
	barberRole, err := server.Store.GetBarberShopChains(ctx, idBarberShopChain)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
			return nil, status.Errorf(codes.Internal, "internal")
		}
	}

	rsp := &barber.GetBarberShopChainsResponse{
		BarbershopChain: ConvertBarberShopChains(barberRole),
	}
	return rsp, nil
}
