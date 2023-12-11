package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func convertBarber(res db.Barber) *barber.Barber {
	return &barber.Barber{
		BarberId:          res.BarberID.String(),
		ShopId:            res.ShopID.UUID.String(),
		NickName:          res.NickName,
		FullName:          res.FullName,
		Phone:             res.Phone,
		Email:             res.Email,
		Gender:            res.Gender,
		Role:              res.Role,
		Status:            int32(res.Status.Int32),
		CreatedAt:         timestamppb.New(res.CreatedAt),
		UpdateAt:          timestamppb.New(res.UpdateAt.Time),
		PasswordChangedAt: timestamppb.New(res.PasswordChangedAt),
	}
}

func convertBarberShop(barberShop db.BarberShop) *barber.BarberShops {
	return &barber.BarberShops{
		OwnerId:     barberShop.OwnerID.String(),
		ShopId:      barberShop.ShopID.String(),
		Status:      barberShop.Status,
		Name:        barberShop.Name,
		Coordinates: barberShop.Coordinates,
		Address:     barberShop.Address,
		Image:       barberShop.Image.String,
		ListImage:   barberShop.ListImage,
		CreatedAt:   timestamppb.New(barberShop.CreatedAt),
	}
}

func ConvertServiceCategory(servicecategory db.ServiceCategory) *barber.ServiceCategory {
	return &barber.ServiceCategory{
		Id:        servicecategory.ID.String(),
		ShopId:    servicecategory.ShopID.String(),
		Name:      servicecategory.Name,
		CreatedAt: timestamppb.New(servicecategory.CreatedAt),
		UpdateAt:  timestamppb.New(servicecategory.UpdateAt.Time),
	}
}

func ConvertServices(service db.Service) *barber.Services {
	return &barber.Services{
		Id:          service.ID.String(),
		CategoryId:  service.CategoryID.String(),
		Name:        service.Name,
		Timer:       &service.Timer.Int32,
		Price:       &service.Timer.Int32,
		Description: &service.Description.String,
		Image:       &service.Image.String,
		CreatedAt:   timestamppb.New(service.CreatedAt),
		UpdateAt:    timestamppb.New(service.UpdateAt.Time),
	}
}

func ConvertListBarberShopsNearby(res []db.FindBarberShopNearbyLocationsRow) []*barber.BarberShops {
	var barberShops []*barber.BarberShops

	for _, barberShop := range res {
		barberShopPB := &barber.BarberShops{
			OwnerId:     barberShop.OwnerID.String(),
			ShopId:      barberShop.ShopID.String(),
			Status:      barberShop.Status,
			Name:        barberShop.Name,
			Coordinates: barberShop.Coordinates,
			Address:     barberShop.Address,
			Image:       barberShop.Image.String,
			ListImage:   barberShop.ListImage,
			CreatedAt:   timestamppb.New(barberShop.CreatedAt),
			Distance:    float32(barberShop.Distance),
			Longitude:   wrapperspb.Double(barberShop.Longitude),
			Latitude:    wrapperspb.Double(barberShop.Latitude),
		}

		barberShops = append(barberShops, barberShopPB)
	}

	return barberShops
}


