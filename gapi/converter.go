package gapi

import (
	db "barbershop/db/sqlc"
	"barbershop/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertBarber(barber db.Barber) *pb.Barber {
	return &pb.Barber{
		BarberId:  barber.BarberID.String(),
		ShopId:    barber.ShopID.UUID.String(),
		NickName:  barber.NickName,
		FullName:  barber.FullName,
		Phone:     barber.Phone,
		Email:     barber.Email,
		Gender:    barber.Gender,
		Role:      barber.Role,
		Status:    int32(barber.Status.Int64),
		CreatedAt: timestamppb.New(barber.CreatedAt),
	}
}
