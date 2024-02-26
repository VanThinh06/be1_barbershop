package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarberShopChains(ctx context.Context, req *barber.UpdateBarberShopChainsRequest) (*barber.UpdateBarberShopChainsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.Id {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}
	var barberShopChainID = uuid.MustParse(req.Id)
	arg := db.UpdateBarberShopChainsParams{
		ID: barberShopChainID,
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Founder: sql.NullString{
			String: req.GetFounder(),
			Valid:  req.Founder != nil,
		},
		FoundingDate: sql.NullTime{
			Time:  req.GetFoundingDate().AsTime(),
			Valid: req.FoundingDate != nil,
		},
		Website: sql.NullString{
			String: req.GetWebsite(),
			Valid:  req.Website != nil,
		},
	}

	barberShopChain, err := server.Store.UpdateBarberShopChains(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {

			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UpdateBarberShopChainsResponse{
		BarbershopChain: ConvertBarberShopChains(barberShopChain),
	}
	return rsp, nil
}
