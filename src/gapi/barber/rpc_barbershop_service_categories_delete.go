package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteBarberShopServiceCategories(ctx context.Context, req *barber.DeleteBarberShopServiceCategoriesRequest) (*barber.DeleteBarberShopServiceCategoriesResponse, error) {

	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var id = uuid.MustParse(req.Id)
	err = server.Store.DeleteBarberShopServiceCategories(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberShopServiceCategoriesResponse{
		Status: "success",
	}
	return rsp, nil
}
