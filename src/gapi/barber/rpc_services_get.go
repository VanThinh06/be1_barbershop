package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetServices(ctx context.Context, req *barber.GetServicesRequest) (*barber.GetServicesResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.GetListServicesParams{
		ShopID: uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetShopId()),
			Valid: req.ChainId != "",
		},

		ChainID: uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetChainId()),
			Valid: req.ChainId != "",
		},
	}

	services, err := server.Store.GetListServices(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetServicesResponse{
		Services: ConvertListSerivces(services),
	}
	return rsp, nil
}
