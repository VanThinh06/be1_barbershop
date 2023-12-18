package customergapi

import (
	"barbershop/src/pb/customer"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarberInShop(ctx context.Context, req *customer.GetBarberInShopRequest) (*customer.GetBarberInShopResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, UnauthenticatedError(err)
	}

	res, err := server.store.GetBarberInBarberShop(ctx, uuid.NullUUID{
		UUID:  uuid.MustParse(req.BarbershopId),
		Valid: req.BarbershopId != "",
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update account")
	}

	rsp := &customer.GetBarberInShopResponse{
		Barbers: convertGetBarberInShop(res),
	}
	return rsp, nil
}
