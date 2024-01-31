package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"

	"github.com/jackc/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// BarberRoles
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

// / BarberShop
// / chains
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

// / barbershops
func convertBarberShops(barberShop db.BarberShop) *barber.BarberShops {
	var barberShopChainId *string
	if barberShop.BarbershopChainID.Valid {
		*barberShopChainId = barberShop.BarbershopChainID.UUID.String()
	}
	return &barber.BarberShops{
		Id:                barberShop.ID.String(),
		BarbershopChainId: barberShopChainId,
		Name:              barberShop.Name,
		IsMainBranch:      barberShop.IsMainBranch,
		BranchCount:       barberShop.BranchCount,
		Coordinates:       barberShop.Coordinates,
		Address:           barberShop.Address,
		Image:             barberShop.Image.String,
		Status:            barberShop.Status,
		Rate:              float32(barberShop.Rate),
		StartTime:         convertToTimeOfDay(barberShop.StartTime),
		EndTime:           convertToTimeOfDay(barberShop.EndTime),
		BreakTime:         convertToTimeOfDay(barberShop.BreakTime),
		BreakMinutes:      barberShop.BreakMinutes,
		IntervalScheduler: barberShop.IntervalScheduler,
		CreateAt:          timestamppb.New(barberShop.CreateAt),
		UpdateAt:          timestamppb.New(barberShop.UpdateAt),
	}
}

func convertBarberManagers(res db.BarberManager) *barber.BarberManagers {
	return &barber.BarberManagers{
		BarberId:  res.BarberID.String(),
		ManagerId: res.ManagerID.String(),
		CreateAt:  timestamppb.New(res.CreateAt),
		UpdateAt:  timestamppb.New(res.UpdateAt),
	}
}

func convertBSServiceCategories(serviceCategory db.BarberShopServiceCategory) *barber.BarberShopServiceCategories {
	return &barber.BarberShopServiceCategories{
		Id:                serviceCategory.ID.String(),
		BarbershopChainId: serviceCategory.BarbershopChainID.UUID.String(),
		BarbershopId:      serviceCategory.BarbershopID.UUID.String(),
		ServiceCategoryId: serviceCategory.ServiceCategoryID.String(),
		CreateAt:          timestamppb.New(serviceCategory.CreateAt),
		UpdateAt:          timestamppb.New(serviceCategory.UpdateAt),
	}
}

func convertListBSServiceCategories(res []db.ListBarberShopServiceCategoriesRow) []*barber.BarberShopServiceCategories {
	var barbers []*barber.BarberShopServiceCategories
	for _, item := range res {
		barber := &barber.BarberShopServiceCategories{
			Id:                item.ID.String(),
			BarbershopChainId: item.BarbershopChainID.UUID.String(),
			BarbershopId:      item.BarbershopID.UUID.String(),
			ServiceCategoryId: item.ServiceCategoryID.String(),
			CreateAt:          timestamppb.New(item.CreateAt),
			UpdateAt:          timestamppb.New(item.UpdateAt),
		}
		barbers = append(barbers, barber)
	}
	return barbers
}

func convertBarberShopServices(service db.BarberShopService) *barber.BarberShopServices {
	return &barber.BarberShopServices{
		Id:                   service.ID.String(),
		BarbershopCategoryId: service.BarbershopCategoryID.String(),
		BarbershopChainId:    service.BarbershopChainID.UUID.String(),
		BarbershopId:         service.BarbershopID.UUID.String(),
	}
}

func convertAppointments(appointment db.CreateAppointmentsRow) *barber.Appointments {
	return &barber.Appointments{
		Id:                  appointment.ID.String(),
		BarbershopId:        appointment.BarbershopID.String(),
		CustomerId:          appointment.CustomerID.String(),
		BarberId:            appointment.BarberID.String(),
		AppointmentDateTime: timestamppb.New(appointment.AppointmentDateTime),
		Status:              appointment.Status,
		CreateAt:            timestamppb.New(appointment.CreateAt),
		UpdateAt:            timestamppb.New(appointment.UpdateAt),
	}
}

func convertListAppointmentsByDate(res []db.ListAppointmentsByDateRow) []*barber.Appointments {
	var appointments []*barber.Appointments
	for _, appointment := range res {
		appointment := &barber.Appointments{
			Id:                  appointment.ID.String(),
			CustomerId:          appointment.CustomerID.String(),
			BarberId:            appointment.BarberID.String(),
			BarbershopId:        appointment.BarbershopID.String(),
			Status:              appointment.Status,
			AppointmentDateTime: timestamppb.New(appointment.AppointmentDateTime),
			CreateAt:            timestamppb.New(appointment.CreateAt),
			UpdateAt:            timestamppb.New(appointment.UpdateAt),
			ServiceTimer:        int32(appointment.ServiceTimer),
		}
		appointments = append(appointments, appointment)
	}
	return appointments
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
			Price:        helpers.ConvertFloat64ToFloat32Pointer(item.Price.Float64),
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
