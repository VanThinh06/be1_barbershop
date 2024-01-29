package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"strings"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateServiceCategories(ctx context.Context, req *barber.CreateServiceCategoriesRequest) (*barber.CreateServiceCategoriesResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	arg := db.CreateServiceCategoriesParams{
		Name:     strings.TrimSpace(req.GetName()),
		IsGlobal: req.GetIsGlobal(),
	}

	servicecategory, err := server.Store.CreateServiceCategories(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {

			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateServiceCategoriesResponse{
		ServiceCategories: convertServiceCategories(servicecategory),
	}
	return rsp, nil
}
