package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteServiceCategories(ctx context.Context, req *barber.DeleteServiceCategoriesRequest) (*barber.DeleteServiceCategoriesResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.GetId() {
		return nil, status.Errorf(codes.PermissionDenied, "no permission")
	}

	var id = uuid.MustParse(req.Id)
	err = server.Store.DeleteServiceCategories(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteServiceCategoriesResponse{
		Status: "success",
	}
	return rsp, nil
}
