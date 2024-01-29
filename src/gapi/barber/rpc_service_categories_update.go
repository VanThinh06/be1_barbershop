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

func (server *Server) UpdateServiceCategories(ctx context.Context, req *barber.UpdateServiceCategoriesRequest) (*barber.UpdateServiceCategoriesResponse, error) {

	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.UpdateServiceCategoriesParams{
		ID: uuid.MustParse(req.Id),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		IsGlobal: sql.NullBool{
			Bool:  req.GetIsGlobal(),
			Valid: req.IsGlobal != nil,
		},
	}

	servicecategory, err := server.Store.UpdateServiceCategories(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.UpdateServiceCategoriesResponse{
		ServiceCategories: convertServiceCategories(servicecategory),
	}
	return rsp, nil
}
