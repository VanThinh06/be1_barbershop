package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) NewServiceCategory(ctx context.Context, req *pb.CreateServiceCategoryRequest) (*pb.CreateServiceCategoryResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	arg := db.CreateServiceCategoryParams{
		ShopID: uuid.MustParse(req.GetBarberShopId()),
		Name:   req.GetName(),
	}

	servicecategory, err := server.Store.CreateServiceCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &pb.CreateServiceCategoryResponse{
		Barber: ConvertServiceCategory(servicecategory),
	}
	return rsp, nil
}
