package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"

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
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}

	if payload.Barber.BarberID.String() != req.OwnerId {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}
	arg := db.CreateChainParams{
		Name:    req.Name,
		OwnerID: payload.Barber.BarberID,
	}

	services, err := server.Store.CreateChain(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Chains_owner_id_key":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	argUpdateChain := db.UpdateChainForBarberShopsParams{
		ChainID: services.ID,
		OwnerID: services.OwnerID,
	}
	err = server.Store.UpdateChainForBarberShops(ctx, argUpdateChain)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_chain_id_facility_idx":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			case "BarberShops_owner_id_facility_idx":
				return nil, status.Errorf(codes.AlreadyExists, "no permission")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.CreateChainResponse{
		Chain: ConvertChain(services),
	}
	return rsp, nil
}
