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

func (server *Server) NewBarberShops(ctx context.Context, req *pb.CreateBarberShopsRequest) (*pb.CreateBarberShopsResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	arg := db.CreateBarberShopsParams{
		CodeBarberShop: req.GetCodeBarberShop(),
		OwnerID:        authPayload.BarberID,
		Name:           req.GetName(),
		Location:       float32(req.GetLocation()),
		Address:        req.GetAddress(),
	}

	barberShop, err := server.store.CreateBarberShops(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	requestBarber := db.UpdateBarberParams{BarberID: authPayload.BarberID, ShopID: uuid.NullUUID{UUID: barberShop.ShopID, Valid: true}}
	_, errUpdateBarber := server.store.UpdateBarber(ctx, requestBarber)
	if errUpdateBarber != nil {
		return nil, status.Errorf(codes.Internal, "Intenal")
	}

	rsp := &pb.CreateBarberShopsResponse{
		BarberShop: convertBarberShop(barberShop),
	}
	return rsp, nil
}
