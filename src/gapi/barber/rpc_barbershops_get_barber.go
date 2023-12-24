package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberInShop(ctx context.Context, req *barber.GetBarberInShopRequest) (*barber.GetBarberInShopResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		_, err = server.AuthorizeCustomer(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	res, err := server.Store.GetBarbersInBarberShop(ctx, uuid.NullUUID{
		UUID:  uuid.MustParse(req.BarbershopId),
		Valid: req.BarbershopId != "",
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update account")
	}

	rsp := &barber.GetBarberInShopResponse{
		Barbers: convertGetBarberInShop(res),
	}
	return rsp, nil
}
