package gapi

import (
	db "barbershop/db/sqlc"
	"barbershop/pb"
	"barbershop/utils"
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarber(ctx context.Context, req *pb.CreateBarberRequest) (*pb.CreateBarberResponse, error) {
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unimplemented, "method CreateBarber not implemented")
	}

	codeBarberShop, err := server.store.GetCodeBarberShop(ctx, req.CodeBarberShop)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to crate user: %s", err)
	}

	arg := db.CreateBarberParams{
		NickName:       req.GetNickname(),
		HashedPassword: hashedPassword,
		Phone:          req.GetPhone(),
		FullName:       req.GetEmail(),
		Gender:         req.GetGender(),
		Email:          req.GetEmail(),
		ShopID:         uuid.NullUUID{UUID: codeBarberShop, Valid: true},
	}

	barber, err := server.store.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Barbers_pkey" {
					return nil, status.Errorf(codes.AlreadyExists, "This account has already existed")
				}
				if pqErr.Constraint == "Barbers_phone_key" {
					return nil, status.Errorf(codes.AlreadyExists, "This phone has already existed")

				}
				if pqErr.Constraint == "Barbers_email_key" {
					return nil, status.Errorf(codes.AlreadyExists, "This email has already existed")
				}
			}
		}
		return nil, status.Errorf(http.StatusForbidden, "Failed to create barber")
	}
	rsp := &pb.CreateBarberResponse{
		Barber: convertBarber(barber),
	}
	return rsp, nil
}
