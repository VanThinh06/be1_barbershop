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

func (server *Server) UpdateServicesCategoryPublicSeparate(ctx context.Context, req *barber.UpdateServicesCategoryPublicSeparateRequest) (*barber.CreateServicesCategoryResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	arg := db.UpdateServicesCategoryPublicSeparateParams{
		ChainID: uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetChainId()),
			Valid: req.ChainId != "",
		},

		Name: req.GetName(),
	}

	servicecategory, err := server.Store.UpdateServicesCategoryPublicSeparate(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateServicesCategoryResponse{
		Barber: ConvertServiceCategory(servicecategory),
	}
	return rsp, nil
}