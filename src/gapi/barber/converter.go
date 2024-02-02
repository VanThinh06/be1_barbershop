package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"

	"github.com/jackc/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// BarberRoles
func ConvertBarberRoles(res db.BarberRole) *barber.BarberRoles {
	return &barber.BarberRoles{
		Id:           res.BarberID.String(),
		BarberId:     res.BarberID.String(),
		BarberShopId: res.BarberShopID.UUID.String(),
		RoleId:       res.RoleID,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
	}
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

// barbershops
func convertBarberShops(barberShop db.BarberShop) *barber.BarberShops {
	var BarberShopChainId *string
	if barberShop.BarberShopChainID.Valid {
		*BarberShopChainId = barberShop.BarberShopChainID.UUID.String()
	}
	return &barber.BarberShops{
		Id:                barberShop.ID.String(),
		BarberShopChainId: BarberShopChainId,
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

func ConvertSearchByNameBarberShops(res []db.SearchByNameBarberShopsRow) []*barber.BarberShops {
	var barberShops []*barber.BarberShops

	for _, barberShop := range res {
		chainIDString := barberShop.BarberShopChainID.UUID.String()
		barberShopPB := &barber.BarberShops{
			Id:                barberShop.ID.String(),
			BarberShopChainId: &chainIDString,
			BranchCount:       barberShop.BranchCount,
			Distance:          float32(barberShop.Distance),
			Longitude:         wrapperspb.Double(barberShop.Longitude),
			Latitude:          wrapperspb.Double(barberShop.Latitude),
			Status:            barberShop.Status,
			Name:              barberShop.Name,
			Coordinates:       barberShop.Coordinates,
			Address:           barberShop.Address,
			Image:             barberShop.Image.String,
			IsReputation:      barberShop.IsReputation,
			Rate:              float32(barberShop.Rate),
		}
		barberShops = append(barberShops, barberShopPB)
	}
	return barberShops
}

func ConvertListBarberShopsNearby(res []db.ListNearbyBarberShopsRow) []*barber.BarberShops {
	var barberShops []*barber.BarberShops
	for _, barberShop := range res {
		barberShopPB := &barber.BarberShops{
			Id:          barberShop.ID.String(),
			Status:      barberShop.Status,
			Name:        barberShop.Name,
			Coordinates: barberShop.Coordinates,
			Address:     barberShop.Address,
			Image:       barberShop.Image.String,
			Distance:    float32(barberShop.Distance),
			Longitude:   wrapperspb.Double(barberShop.Longitude),
			Latitude:    wrapperspb.Double(barberShop.Latitude),
		}

		barberShops = append(barberShops, barberShopPB)
	}

	return barberShops
}

// barbers

func convertCreateBarbers(res db.Barber) *barber.Barbers {
	return &barber.Barbers{
		Id:        res.ID.String(),
		GenderId:  res.GenderID,
		Email:     res.Email,
		Phone:     res.Phone,
		NickName:  res.NickName,
		FullName:  res.FullName,
		Haircut:   res.Haircut,
		AvatarUrl: res.AvatarUrl.String,
		CreateAt:  timestamppb.New(res.CreateAt),
		UpdateAt:  timestamppb.New(res.UpdateAt),
	}
}
func convertBarbers(res db.GetBarbersRow) *barber.Barbers {
	return &barber.Barbers{
		Id:           res.ID.String(),
		GenderId:     res.GenderID,
		Email:        res.Email,
		Phone:        res.Phone,
		NickName:     res.NickName,
		FullName:     res.FullName,
		Haircut:      res.Haircut,
		AvatarUrl:    res.AvatarUrl.String,
		BarberRoleId: res.BarberRoleID,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
	}
}

func convertBarbersEmail(res db.GetBarbersEmailRow) *barber.Barbers {
	return &barber.Barbers{
		Id:           res.ID.String(),
		GenderId:     res.GenderID,
		Email:        res.Email,
		Phone:        res.Phone,
		NickName:     res.NickName,
		FullName:     res.FullName,
		Haircut:      res.Haircut,
		AvatarUrl:    res.AvatarUrl.String,
		BarberRoleId: res.BarberRole,
		CreateAt:     timestamppb.New(res.CreateAt),
		UpdateAt:     timestamppb.New(res.UpdateAt),
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

func ConvertListBarberManagers(res []db.BarberManager) []*barber.BarberManagers {
	var barberManagers []*barber.BarberManagers

	for _, item := range res {
		barberManager := &barber.BarberManagers{

			BarberId:  item.BarberID.String(),
			ManagerId: item.ManagerID.String(),
			CreateAt:  timestamppb.New(item.CreateAt),
			UpdateAt:  timestamppb.New(item.UpdateAt),
		}
		barberManagers = append(barberManagers, barberManager)
	}

	return barberManagers
}




func convertBSServiceCategories(serviceCategory db.BarberShopServiceCategory) *barber.BarberShopServiceCategories {
	return &barber.BarberShopServiceCategories{
		Id:                serviceCategory.ID.String(),
		BarberShopChainId: serviceCategory.BarberShopChainID.UUID.String(),
		BarberShopId:      serviceCategory.BarberShopID.UUID.String(),
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
			BarberShopChainId: item.BarberShopChainID.UUID.String(),
			BarberShopId:      item.BarberShopID.UUID.String(),
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
		BarberShopChainId:    service.BarberShopChainID.UUID.String(),
		BarberShopId:         service.BarberShopID.UUID.String(),
	}
}

func convertAppointments(appointment db.CreateAppointmentsRow) *barber.Appointments {
	return &barber.Appointments{
		Id:                  appointment.ID.String(),
		BarberShopId:        appointment.BarberShopID.String(),
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
			BarberShopId:        appointment.BarberShopID.String(),
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
