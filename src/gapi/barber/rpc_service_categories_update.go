package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"database/sql"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateServiceCategory(ctx context.Context, req *barber.UpdateServiceCategoryRequest) (*barber.UpdateServiceCategoryResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.UpdateServiceCategoryParams{
		ID: int16(req.Id),
		Name: sql.NullString{
			String: req.GetName(),
			Valid: req.Name != nil,
		},
	}

	 err = server.Store.UpdateServiceCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UpdateServiceCategoryResponse{
		Message: "Message",
	}
	return rsp, nil
}
