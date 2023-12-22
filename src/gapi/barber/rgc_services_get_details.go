package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetServiceDetails(ctx context.Context, req *barber.GetServiceDetailsRequest) (*barber.GetServiceDetailsResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}
	var shopId uuid.NullUUID
	var chainId uuid.NullUUID
	if req.GetShopId() != "" {
		shopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetShopId()),
			Valid: req.ShopId != "",
		}
	}
	if req.GetChainId() != "" {
		shopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetChainId()),
			Valid: req.ShopId != "",
		}
	}
	arg := db.GetListServiceDetailsParams{
		ChainID: chainId,
		ShopID:  shopId,
	}

	serviceDetail, err := server.Store.GetListServiceDetails(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetServiceDetailsResponse{

		ServiceDetails: ConvertListCategorySerivceDetails(serviceDetail),
	}
	return rsp, nil
}
