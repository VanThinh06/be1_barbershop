package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertBarber(barber db.Barber) *pb.Barber {
	return &pb.Barber{
		BarberId:          barber.BarberID.String(),
		ShopId:            barber.ShopID.UUID.String(),
		NickName:          barber.NickName,
		FullName:          barber.FullName,
		Phone:             barber.Phone,
		Email:             barber.Email,
		Gender:            barber.Gender,
		Role:              barber.Role,
		Status:            int32(barber.Status.Int32),
		CreatedAt:         timestamppb.New(barber.CreatedAt),
		UpdateAt:          timestamppb.New(barber.UpdateAt.Time),
		PasswordChangedAt: timestamppb.New(barber.PasswordChangedAt),
	}
}

func convertBarberShop(barberShop db.BarberShop) *pb.BarberShops {
	return &pb.BarberShops{
		OwnerId:   barberShop.OwnerID.String(),
		ShopId:    barberShop.ShopID.String(),
		Status:    barberShop.Status,
		Name:      barberShop.Name,
		Location:  uint32(barberShop.Location),
		Address:   barberShop.Address,
		Image:     barberShop.Image.String,
		ListImage: barberShop.ListImage,
		CreatedAt: timestamppb.New(barberShop.CreatedAt),
	}
}
