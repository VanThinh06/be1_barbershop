package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetServiceCategories(ctx context.Context, req *barber.GetServiceCategoriesRequest) (*barber.GetServiceCategoriesResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	var shopId uuid.NullUUID = uuid.NullUUID{
		UUID:  uuid.MustParse(req.ShopId),
		Valid: req.ShopId != "",
	}
	serviceCategories, err := server.Store.GetListServiceCategories(ctx, shopId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetServiceCategoriesResponse{
		ServiceCategories: ConvertListSerivceCategories(serviceCategories),
	}
	return rsp, nil
}
