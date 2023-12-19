package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateChain(ctx context.Context, req *barber.CreateChainRequest) (*barber.CreateChainResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	arg := db.CreateChainParams{
		Name:    req.Name,
		OwnerID: uuid.MustParse(req.OwnerId),
	}

	services, err := server.Store.CreateChain(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateChainResponse{
		Chain: ConvertChain(services),
	}
	return rsp, nil
}
