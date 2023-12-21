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
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}
	var shopId uuid.NullUUID
	if req.GetShopId() != "" {
		shopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetShopId()),
			Valid: req.ShopId != "",
		}
	}
	arg := db.GetListServicesParams{
		CategoryID: uuid.MustParse(req.GetCategoryId()),
		ShopID:     shopId,
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
