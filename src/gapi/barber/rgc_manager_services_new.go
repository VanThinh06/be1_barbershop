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

func (server *Server) CreateService(ctx context.Context, req *barber.CreateServiceRequest) (*barber.CreateServiceResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	err = server.IsManager(ctx, payload)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "No permission")
	}

	arg := db.CreateServiceParams{
		CategoryID: uuid.MustParse(req.GetCategoryId()),
		Timer: sql.NullInt32{
			Int32: req.GetTimer(),
			Valid: req.Timer != nil,
		},
		Price: sql.NullFloat64{
			Float64: float64(req.GetPrice()),
			Valid:   req.Price != nil,
		},
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},
		Name: req.Name,
	}

	services, err := server.Store.CreateService(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	rsp := &barber.CreateServiceResponse{
		Services: ConvertServices(services),
	}
	return rsp, nil
}
