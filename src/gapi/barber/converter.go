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
		UpdatedAt:         timestamppb.New(res.UpdatedAt.Time),
		PasswordChangedAt: timestamppb.New(res.PasswordChangedAt),
		Haircut:           res.Haircut,
	}
}

func convertBarberShop(barberShop db.BarberShop) *barber.BarberShop {
	return &barber.BarberShop{
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
		ShopId:    servicecategory.ShopID.UUID.String(),
		Name:      servicecategory.Name,
		CreatedAt: timestamppb.New(servicecategory.CreatedAt),
		UpdatedAt: timestamppb.New(servicecategory.UpdatedAt.Time),
		Gender:    "",
		IsHidden:  servicecategory.Hidden,
		ChainId:   servicecategory.ChainID.UUID.String(),
	}
}

func ConvertServices(service db.Service) *barber.Service {
	return &barber.Service{
		Id:          service.ID.String(),
		CategoryId:  service.CategoryID.String(),
		Name:        service.Name,
		Timer:       &service.Timer.Int32,
		Price:       &service.Timer.Int32,
		Description: &service.Description.String,
		Image:       &service.Image.String,
		CreatedAt:   timestamppb.New(service.CreatedAt),
		UpdatedAt:   timestamppb.New(service.UpdatedAt.Time),
		ShopId:      service.ShopID.UUID.String(),
		ChainId:     service.ChainID.UUID.String(),
		IsHidden:    service.Hidden,
	}
}

func ConvertListBarberShopsNearby(res []db.FindBarberShopsNearbyLocationsRow) []*barber.BarberShop {
	var barberShops []*barber.BarberShop

	for _, barberShop := range res {
		barberShopPB := &barber.BarberShop{
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

func ConvertChain(res db.Chain) *barber.Chain {
	return &barber.Chain{
		ChainId:   res.ChainID.String(),
		OwnerId:   res.OwnerID.String(),
		Name:      res.Name,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt.Time),
	}
}
