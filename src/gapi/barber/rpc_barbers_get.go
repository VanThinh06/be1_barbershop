package gapi

import (
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetBarber(ctx context.Context, req *barber.GetBarberRequest) (*barber.GetBarberResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, UnauthenticatedError(err)
	}

	res, err := server.Store.ReadBarber(ctx, uuid.MustParse(req.BarberId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "barber not found")
		}
		return nil, status.Errorf(codes.Internal, "barber not found")
	}

	rsp := convertBarberDetail(res)
	return rsp, nil
}
