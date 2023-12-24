package gapi

import (
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberShop(ctx context.Context, req *barber.GetBarberShopRequest) (*barber.GetBarberShopResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	res, err := server.Store.GetBarberShop(ctx, uuid.MustParse(req.ShopId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopResponse{
		BarberShop: convertBarberShop(res),
	}
	return rsp, nil
}
