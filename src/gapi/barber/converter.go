package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utils"

	"github.com/jackc/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ConvertBarberRoles(res db.BarberRole) *barber.BarberRoles {
	return &barber.BarberRoles{
		Id:           res.BarberID.String(),
		BarberId:     res.BarberID.String(),
		BarbershopId: res.BarbershopID.String(),
		RoleId:       res.RoleID,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
	}
}

func ConvertBarberRolesDetail(res db.GetBarberRolesRow) *barber.BarberRoles {
	return &barber.BarberRoles{
		Id:           res.BarberID.String(),
		BarberId:     res.BarberID.String(),
		BarbershopId: res.BarbershopID.String(),
		RoleId:       res.RoleID,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
	}
}

func ConvertListBarbersRoles(res []db.ListBarbersRolesRow) []*barber.BarberRoles {
	var barbersRoles []*barber.BarberRoles

	for _, item := range res {
		barberRole := &barber.BarberRoles{
			Id:           item.ID.String(),
			BarberId:     item.BarberID.String(),
			BarbershopId: item.BarbershopID.String(),
			RoleId:       item.RoleID,
			CreateAt:     timestamppb.New(item.CreateAt),
			UpdateAt:     timestamppb.New(item.UpdateAt),
		}
		barbersRoles = append(barbersRoles, barberRole)
	}

	return barbersRoles
}

func ConvertBarberShopChains(res db.BarberShopChain) *barber.BarberShopChains {
	return &barber.BarberShopChains{
		Id:           res.ID.String(),
		Name:         res.Name,
		Description:  res.Description.String,
		Founder:      res.Founder,
		FoundingDate: timestamppb.New(res.FoundingDate),
		Website:      res.Website.String,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
	}
}

func convertBarber(res db.Barber) *barber.Barber {
	return &barber.Barber{
		BarberId: res.ID.String(),
		ShopId:   res.ShopID.UUID.String(),
		NickName: res.NickName,
		FullName: res.FullName,
		Phone:    res.Phone,
		Email:    res.Email,
		Gender:   res.Gender,
		Role:     res.Role,
		Status:   int32(res.Status.Int32),
		Haircut:  res.Haircut,
	}
}

func convertBarberDetail(res db.ReadBarberRow) *barber.GetBarberResponse {
	return &barber.GetBarberResponse{
		Barber: &barber.Barber{
			BarberId: res.ID.String(),
			ShopId:   res.ShopID.UUID.String(),
			NickName: res.NickName,
			FullName: res.FullName,
			Phone:    res.Phone,
			Email:    res.Email,
			Gender:   res.Gender,
			Role:     res.Role,
			Status:   int32(res.Status.Int32),
			Haircut:  res.Haircut,
		},
		ShopName:          res.ShopName,
		ShopAddress:       res.ShopAddress,
		ShopImage:         res.Avatar.String,
		ShopFacility:      res.ShopFacility,
		Reputation:        res.ShopReputation,
		ShopRate:          float32(res.ShopRate),
		StartTime:         convertToTimeOfDay(res.ShopStartTime),
		EndTime:           convertToTimeOfDay(res.ShopEndTime),
		BreakTime:         convertToTimeOfDay(res.ShopBreakTime),
		BreakMinutes:      res.ShopBreakMinutes,
		IntervalScheduler: res.ShopIntervalScheduler,
	}
}

func convertGetBarberInShop(res []db.Barber) []*barber.Barber {
	var barbers []*barber.Barber
	for _, item := range res {
		barber := &barber.Barber{
			BarberId: item.ID.String(),
			Status:   item.Status.Int32,
			ShopId:   item.ShopID.UUID.String(),
			NickName: item.NickName,
			FullName: item.FullName,
			Phone:    item.Phone,
			Email:    item.Email,
			Gender:   item.Gender,
			Role:     item.Role,
			Avatar:   item.Avatar.String,
			Haircut:  item.Haircut,
		}
		barbers = append(barbers, barber)
	}
	return barbers
}

func convertBarberShop(barberShop db.BarberShop) *barber.BarberShops {
	return &barber.BarberShops{
		Id:            barberShop.ID.String(),
		Status:            barberShop.Status,
		Name:              barberShop.Name,
		Coordinates:       barberShop.Coordinates,
		Address:           barberShop.Address,
		Image:             barberShop.Image.String,
		StartTime:         convertToTimeOfDay(barberShop.StartTime),
		EndTime:           convertToTimeOfDay(barberShop.EndTime),
		BreakTime:         convertToTimeOfDay(barberShop.BreakTime),
		BreakMinutes:      barberShop.BreakMinutes,
		IntervalScheduler: barberShop.IntervalScheduler,
		CreateAt:         timestamppb.New(barberShop.CreateAt),
	}
}

func ConvertServiceCategory(servicecategory db.ServiceCategory) *barber.ServiceCategory {
	return &barber.ServiceCategory{
		Id:        servicecategory.ID.String(),
		ShopId:    servicecategory.ShopID.UUID.String(),
		Name:      servicecategory.Name,
		CreatedAt: timestamppb.New(servicecategory.CreatedAt),
		UpdatedAt: timestamppb.New(servicecategory.UpdatedAt.Time),
		Gender:    servicecategory.Gender,
		ChainId:   servicecategory.ChainID.UUID.String(),
	}
}

func ConvertServices(service db.Service) *barber.Service {
	return &barber.Service{
		Id:          service.ID.String(),
		CategoryId:  service.CategoryID.String(),
		Name:        service.Name,
		Timer:       &service.Timer.Int32,
		Price:       utils.ConvertFloat64ToFloat32Pointer(service.Price.Float64),
		Description: &service.Description.String,
		Image:       &service.Image.String,
		CreatedAt:   timestamppb.New(service.CreatedAt),
		UpdatedAt:   timestamppb.New(service.UpdatedAt.Time),
		ShopId:      service.ShopID.UUID.String(),
	}
}

func ConvertListSerivces(res []db.Service) []*barber.Service {
	var services []*barber.Service

	for _, service := range res {
		barberShopPB := &barber.Service{
			Id:          service.ID.String(),
			CategoryId:  service.CategoryID.String(),
			ShopId:      service.ShopID.UUID.String(),
			Name:        service.Name,
			Image:       &service.Image.String,
			CreatedAt:   timestamppb.New(service.CreatedAt),
			UpdatedAt:   timestamppb.New(service.UpdatedAt.Time),
			Timer:       &timestamppb.Now().Nanos,
			Price:       utils.ConvertFloat64ToFloat32Pointer(service.Price.Float64),
			Description: &service.Description.String,
		}
		services = append(services, barberShopPB)
	}

	return services
}

func ConvertListCategorySerivceDetails(res []db.GetListServiceDetailsRow) []*barber.ServiceDetail {
	var services []*barber.ServiceDetail

	for _, item := range res {
		serviceDetail := &barber.ServiceDetail{
			CategoryId:   item.CategoryID.String(),
			NameCategory: item.CategoryName,
			Id:           item.ServiceID.String(),
			Name:         item.ServiceName,
			Image:        &item.Image.String,
			Timer:        &timestamppb.Now().Nanos,
			Price:        utils.ConvertFloat64ToFloat32Pointer(item.Price.Float64),
			Description:  &item.Description.String,
		}
		services = append(services, serviceDetail)
	}

	return services
}

func ConvertListSerivceCategories(res []db.ServiceCategory) []*barber.ServiceCategory {
	var serviceCategories []*barber.ServiceCategory

	for _, serviceCategory := range res {
		item := &barber.ServiceCategory{
			Id:        serviceCategory.ID.String(),
			ShopId:    serviceCategory.ShopID.UUID.String(),
			ChainId:   serviceCategory.ChainID.UUID.String(),
			Name:      serviceCategory.Name,
			Gender:    serviceCategory.Gender,
			CreatedAt: timestamppb.New(serviceCategory.CreatedAt),
			UpdatedAt: timestamppb.New(serviceCategory.UpdatedAt.Time),
		}
		serviceCategories = append(serviceCategories, item)
	}

	return serviceCategories
}

func ConvertListBarberShopsNearby(res []db.FindBarberShopsNearbyLocationsRow) []*barber.BarberShop {
	var barberShops []*barber.BarberShop

	for _, barberShop := range res {
		barberShopPB := &barber.BarberShop{
			OwnerId:     barberShop.OwnerID.String(),
			ShopId:      barberShop.ID.String(),
			Status:      barberShop.Status,
			Name:        barberShop.Name,
			Coordinates: barberShop.Coordinates,
			Address:     barberShop.Address,
			Image:       barberShop.Image.String,
			CreatedAt:   timestamppb.New(barberShop.CreatedAt),
			Distance:    float32(barberShop.Distance),
			Longitude:   wrapperspb.Double(barberShop.Longitude),
			Latitude:    wrapperspb.Double(barberShop.Latitude),
		}

		barberShops = append(barberShops, barberShopPB)
	}

	return barberShops
}

func ConvertListBarberShops(res []db.BarberShop) []*barber.BarberShop {
	var barberShops []*barber.BarberShop

	for _, barberShop := range res {
		chainIDString := barberShop.ChainID.UUID.String()
		barberShopPB := &barber.BarberShop{
			OwnerId:     barberShop.OwnerID.String(),
			ShopId:      barberShop.ID.String(),
			Status:      barberShop.Status,
			Name:        barberShop.Name,
			Coordinates: barberShop.Coordinates,
			Address:     barberShop.Address,
			Image:       barberShop.Image.String,

			CreatedAt:         timestamppb.New(barberShop.CreatedAt),
			StartTime:         convertToTimeOfDay(barberShop.StartTime),
			EndTime:           convertToTimeOfDay(barberShop.EndTime),
			BreakTime:         convertToTimeOfDay(barberShop.BreakTime),
			BreakMinutes:      barberShop.BreakMinutes,
			IntervalScheduler: barberShop.IntervalScheduler,
			IsReputation:      barberShop.Reputation,
			Facility:          barberShop.Facility,
			Rate:              float32(barberShop.Rate),
			ChainId:           &chainIDString,
		}
		barberShops = append(barberShops, barberShopPB)
	}

	return barberShops
}

func ConvertChain(res db.Chain) *barber.Chain {
	return &barber.Chain{
		ChainId:   res.ID.String(),
		OwnerId:   res.OwnerID.String(),
		Name:      res.Name,
		CreatedAt: timestamppb.New(res.CreatedAt),
		UpdatedAt: timestamppb.New(res.UpdatedAt.Time),
	}
}

func convertToTimeOfDay(pgTime pgtype.Time) *barber.TimeOfDay {
	if pgTime.Status == pgtype.Null || pgTime.Microseconds < 0 {
		return nil
	}

	totalMinutes := pgTime.Microseconds / 60e6
	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	return &barber.TimeOfDay{
		Hours:   int32(hours),
		Minutes: int32(minutes),
	}
}
