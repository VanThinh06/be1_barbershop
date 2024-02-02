package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

func (server *Server) GetBarber(ctx context.Context, req *barber.GetBarbersRequest) (*barber.GetBarbersResponse, error) {
	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.GetBarbersParams{
		ID:   uuid.MustParse(req.Id),
		ID_2: uuid.MustParse(req.BarberShopId),
	}
	res, err := server.Store.GetBarbers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
		return nil, internalError(err)
	}

	rsp := &barber.GetBarbersResponse{
		Barber: convertBarbers(res),
	}
	return rsp, nil
}
