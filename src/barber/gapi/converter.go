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


func ConvertServiceCategory(servicecategory db.ServiceCategory) *pb.ServiceCategory {
	return &pb.ServiceCategory{
		Id: servicecategory.ID.String(),
		ShopId: servicecategory.ShopID.String(),
		Name: servicecategory.Name,
		CreatedAt: timestamppb.New(servicecategory.CreatedAt),
		UpdateAt: timestamppb.New(servicecategory.UpdateAt.Time),
	}
}

func ConvertServices(service db.Service) *pb.Services {
	return &pb.Services{
		Id: service.ID.String(),
		CategoryId: service.CategoryID.String(),
		Name: service.Name,
		Timer: &service.Timer.Int32,
		Price: &service.Timer.Int32,
		Description: &service.Description.String,
		Image: &service.Image.String,
		CreatedAt: timestamppb.New(service.CreatedAt),
		UpdateAt: timestamppb.New(service.UpdateAt.Time),
	}
}
