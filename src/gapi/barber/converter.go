package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// BarberRoles
func convertBarberRoles(res db.BarberRole) *barber.BarberRoles {
	return &barber.BarberRoles{
		Id:           res.BarberID.String(),
		BarberId:     res.BarberID.String(),
		BarberShopId: res.BarberShopID.String(),
		RoleId:       int32(res.RoleID),
	}
}

// / BarberShop
// / chains
func ConvertBarberShopChains(res db.BarberShopChain) *barber.BarberShopChains {
	return &barber.BarberShopChains{
		Id:      res.ID.String(),
		Name:    res.Name,
		Founder: res.Founder,
		Website: res.Website.String,
	}
}

// barbershops
func convertBarberShops(barberShop db.BarberShop) *barber.BarberShops {
	var barberShopChainId string
	if barberShop.BarberShopChainID.Valid {
		barberShopChainId = barberShop.BarberShopChainID.UUID.String()
	}
	return &barber.BarberShops{
		Id:                       barberShop.ID.String(),
		BarberShopChainId:        barberShopChainId,
		Name:                     barberShop.Name,
		ProvinceId:               int32(barberShop.ProvinceID),
		DistrictId:               int32(barberShop.DistrictID),
		WardId:                   int32(barberShop.ProvinceID),
		SpecificLocation:         barberShop.SpecificLocation,
		Phone:                    barberShop.Phone,
		Email:                    barberShop.Email,
		WebsiteUrl:               barberShop.WebsiteUrl.String,
		AvatarUrl:                barberShop.AvatarUrl,
		CoverPhotoUrl:            barberShop.CoverPhotoUrl,
		FacadePhotoUrl:           barberShop.FacadePhotoUrl,
		RepresentativeName:       barberShop.RepresentativeName,
		CitizenId:                barberShop.CitizenID,
		RepresentativePhone:      barberShop.RepresentativePhone,
		RepresentativeEmail:      barberShop.RepresentativeEmail,
		RepresentativePhoneOther: barberShop.RepresentativePhoneOther.String,
		StartTimeMonday:          barberShop.StartTimeMonday.Microseconds,
		EndTimeMonday:            barberShop.EndTimeMonday.Microseconds,
		StartTimeTuesday:         barberShop.StartTimeTuesday.Microseconds,
		EndTimeTuesday:           barberShop.EndTimeTuesday.Microseconds,
		StartTimeWednesday:       barberShop.StartTimeWednesday.Microseconds,
		EndTimeWednesday:         barberShop.EndTimeWednesday.Microseconds,
		StartTimeThursday:        barberShop.StartTimeThursday.Microseconds,
		EndTimeThursday:          barberShop.EndTimeThursday.Microseconds,
		StartTimeFriday:          barberShop.StartTimeFriday.Microseconds,
		EndTimeFriday:            barberShop.EndTimeFriday.Microseconds,
		StartTimeSaturday:        barberShop.StartTimeSaturday.Microseconds,
		EndTimeSaturday:          barberShop.EndTimeSaturday.Microseconds,
		StartTimeSunday:          barberShop.StartTimeSunday.Microseconds,
		EndTimeSunday:            barberShop.EndTimeSunday.Microseconds,
		IntervalScheduler:        int32(barberShop.IntervalScheduler),
		IsMainBranch:             barberShop.IsMainBranch,
		IsReputation:             barberShop.IsReputation,
		IsVerified:               barberShop.IsVerified,
		CreateAt:                 timestamppb.New(barberShop.CreateAt),
	}
}

// barbershops
func convertListBarberShops(res []db.ListBarberShopsRow) []*barber.BarberShopInList {
	var barberShops []*barber.BarberShopInList
	for _, barberShop := range res {
		barberShopPB := &barber.BarberShopInList{
			Id:                barberShop.ID.String(),
			BarberShopChainId: barberShop.BarberShopChainID.UUID.String(),
			Name:              barberShop.Name,
			ProvinceId:        int32(barberShop.ProvinceID),
			DistrictId:        int32(barberShop.DistrictID),
			WardId:            int32(barberShop.ProvinceID),
			SpecificLocation:  barberShop.SpecificLocation,
			Phone:             barberShop.Phone,
			Email:             barberShop.Email,
			WebsiteUrl:        barberShop.WebsiteUrl.String,
			AvatarUrl:         barberShop.AvatarUrl,
			CoverPhotoUrl:     barberShop.CoverPhotoUrl,
			FacadePhotoUrl:    barberShop.FacadePhotoUrl,
			ProvinceName:      barberShop.ProvinceName,
			DistrictName:      barberShop.DistrictName,
			WardName:          barberShop.WardName,
			IsMainBranch:      barberShop.IsMainBranch,
			IsReputation:      barberShop.IsReputation,
		}
		barberShops = append(barberShops, barberShopPB)
	}
	return barberShops
}

func ConvertSearchByNameBarberShops(res []db.SearchByNameBarberShopsRow) []*barber.BarberShops {
	var barberShops []*barber.BarberShops

	for _, barberShop := range res {
		// chainIDString := barberShop.BarberShopChainID.UUID.String()
		barberShopPB := &barber.BarberShops{
			Id: barberShop.ID.String(),
			// BarberShopChainId: &chainIDString,
			// BranchCount:       int32(barberShop.BranchNumber.Int16),
			// Distance: float32(barberShop.Distance),
			// Longitude:         wrapperspb.Double(barberShop.Longitude),
			// Latitude:          wrapperspb.Double(barberShop.Latitude),
			Name: barberShop.Name,
			// Coordinates:       barberShop.Coordinates,
			IsReputation: barberShop.IsReputation,
		}
		barberShops = append(barberShops, barberShopPB)
	}
	return barberShops
}

func ConvertListBarberShopsNearby(res []db.ListNearbyBarberShopsRow) []*barber.BarberShops {
	var barberShops []*barber.BarberShops
	for _, barberShop := range res {
		barberShopPB := &barber.BarberShops{
			Id:   barberShop.ID.String(),
			Name: barberShop.Name,
			// Coordinates: barberShop.Coordinates,
			// Distance: float32(barberShop.Distance),
			// Longitude:   wrapperspb.Double(barberShop.Longitude),
			// Latitude:    wrapperspb.Double(barberShop.Latitude),
		}

		barberShops = append(barberShops, barberShopPB)
	}

	return barberShops
}

// barbers

func convertCreateBarbers(res db.Barber) *barber.Barbers {
	return &barber.Barbers{
		Id:        res.ID.String(),
		GenderId:  int32(res.GenderID),
		Email:     res.Email.String,
		Phone:     res.Phone,
		NickName:  res.NickName,
		FullName:  res.FullName,
		Haircut:   res.Haircut,
		AvatarUrl: res.AvatarUrl.String,
	}
}
func convertBarbers(res db.GetBarbersRow) *barber.Barbers {
	return &barber.Barbers{
		Id:        res.ID.String(),
		GenderId:  int32(res.GenderID),
		Email:     res.Email.String,
		Phone:     res.Phone,
		NickName:  res.NickName,
		FullName:  res.FullName,
		Haircut:   res.Haircut,
		AvatarUrl: res.AvatarUrl.String,
	}
}

func convertBarberContact(res db.GetUserBarberRow) *barber.Barbers {
	return &barber.Barbers{
		Id:        res.ID.String(),
		GenderId:  int32(res.GenderID),
		Email:     res.Email.String,
		Phone:     res.Phone,
		NickName:  res.NickName,
		FullName:  res.FullName,
		Haircut:   res.Haircut,
		AvatarUrl: res.AvatarUrl.String,
	}
}

func convertBarberEmployees(res []db.GetBarberEmployeesRow) []*barber.BarberEmployees {
	var barberEmployees []*barber.BarberEmployees

	for _, item := range res {
		barberEmployee := &barber.BarberEmployees{
			Barber: &barber.Barbers{
				Id:        item.ID.String(),
				GenderId:  int32(item.GenderID),
				Email:     item.Email.String,
				Phone:     item.Phone,
				NickName:  item.NickName,
				FullName:  item.FullName,
				Haircut:   item.Haircut,
				AvatarUrl: item.AvatarUrl.String,
			},
			BarberRole: &barber.BarberRoles{
				Id:           item.BarberID.String(),
				BarberId:     item.BarberID.String(),
				BarberShopId: item.BarberShopID.String(),
				RoleId:       int32(item.RoleID),
			},
		}
		barberEmployees = append(barberEmployees, barberEmployee)
	}

	return barberEmployees
}

func convertBarberManagers(res db.BarberManager) *barber.BarberManagers {
	return &barber.BarberManagers{
		BarberId:  res.BarberID.String(),
		ManagerId: res.ManagerID.String(),
	}
}

func ConvertListBarberManagers(res []db.BarberManager) []*barber.BarberManagers {
	var barberManagers []*barber.BarberManagers

	for _, item := range res {
		barberManager := &barber.BarberManagers{

			BarberId:  item.BarberID.String(),
			ManagerId: item.ManagerID.String(),
		}
		barberManagers = append(barberManagers, barberManager)
	}

	return barberManagers
}

func convertBSServiceCategories(serviceCategory db.BarberShopServiceCategory) *barber.BarberShopServiceCategories {
	return &barber.BarberShopServiceCategories{
		Id: serviceCategory.ID.String(),
	}
}

func convertListBSServiceCategories(res []db.ListBarberShopServiceCategoriesRow) []*barber.BarberShopServiceCategories {
	var barbers []*barber.BarberShopServiceCategories
	for _, item := range res {
		barber := &barber.BarberShopServiceCategories{
			Id: item.ID.String(),
		}
		barbers = append(barbers, barber)
	}
	return barbers
}

func convertBarberShopServices(service db.BarberShopService) *barber.BarberShopServices {
	return &barber.BarberShopServices{
		Id:                   service.ID.String(),
		BarbershopCategoryId: service.BarbershopCategoryID.String(),
	}
}

func convertProvinces(res []db.Province) []*barber.Provinces {
	var provinces []*barber.Provinces

	for _, item := range res {
		province := &barber.Provinces{
			Id:   int32(item.ID),
			Name: item.Name,
		}
		provinces = append(provinces, province)
	}

	return provinces
}

func convertDistricts(res []db.District) []*barber.Districts {
	var districts []*barber.Districts

	for _, item := range res {
		district := &barber.Districts{
			Id:         int32(item.ID),
			Name:       item.Name,
			ProvinceId: int32(item.ProvinceID),
		}
		districts = append(districts, district)
	}

	return districts
}

func convertWards(res []db.Ward) []*barber.Wards {
	var wards []*barber.Wards

	for _, item := range res {
		ward := &barber.Wards{
			Id:         int32(item.ID),
			Name:       item.Name,
			DistrictId: int32(item.DistrictID),
		}
		wards = append(wards, ward)
	}

	return wards
}

func convertAppointments(appointment db.CreateAppointmentsRow) *barber.Appointments {
	return &barber.Appointments{
		Id:                  appointment.ID.String(),
		BarberShopId:        appointment.BarberShopID.String(),
		CustomerId:          appointment.CustomerID.String(),
		BarberId:            appointment.BarberID.String(),
		AppointmentDateTime: timestamppb.New(appointment.AppointmentDateTime),
		Status:              int32(appointment.Status),
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
			Status:              int32(appointment.Status),
			AppointmentDateTime: timestamppb.New(appointment.AppointmentDateTime),
			CreateAt:            timestamppb.New(appointment.CreateAt),
			UpdateAt:            timestamppb.New(appointment.UpdateAt),
			ServiceTimer:        int32(appointment.ServiceTimer),
		}
		appointments = append(appointments, appointment)
	}
	return appointments
}
