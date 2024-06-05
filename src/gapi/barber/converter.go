package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertListPermission(res []db.Permission) []string {
	permissions := make([]string, len(res))
	for i, item := range res {
		permissions[i] = item.Name
	}
	return permissions
}


// BarberRoles
func convertBarberRoles(res db.BarberRole) *barber.BarberRoles {
	return &barber.BarberRoles{
		Id:           res.BarberID.String(),
		BarberId:     res.BarberID.String(),
		BarberShopId: res.BarberShopID.String(),
		RoleId:       int32(res.RoleID),
	}
}

// chains
func ConvertBarberShopChains(res db.BarberShopChain) *barber.BarberShopChains {
	return &barber.BarberShopChains{
		Id:      res.ID.String(),
		Name:    res.Name,
		Founder: res.Founder,
		Website: res.Website.String,
	}
}

// barbershops
func convertBarberShop(barberShop db.BarberShop) *barber.BarberShop {
	var barberShopChainId string
	if barberShop.BarberShopChainID.Valid {
		barberShopChainId = barberShop.BarberShopChainID.UUID.String()
	}
	return &barber.BarberShop{
		Id:                barberShop.ID.String(),
		BarberShopChainId: barberShopChainId,
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
		IsMainBranch:      barberShop.IsMainBranch,
		IsReputation:      barberShop.IsReputation,
	}
}

func convertGetBarberShop(barberShop db.GetBarberShopRow) *barber.BarberShop {
	return &barber.BarberShop{
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
		IsMainBranch:      barberShop.IsMainBranch,
		IsReputation:      barberShop.IsReputation,
		ProvinceName:      barberShop.ProvinceName,
		DistrictName:      barberShop.DistrictName,
		WardName:          barberShop.WardName,
	}
}

func convertListBarberShops(res []db.ListBarberShopsRow) []*barber.BarberShop {
	var barberShops []*barber.BarberShop
	for _, barberShop := range res {
		barberShopPB := &barber.BarberShop{
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
			IsMainBranch:      barberShop.IsMainBranch,
			IsReputation:      barberShop.IsReputation,
			ProvinceName:      barberShop.ProvinceName,
			DistrictName:      barberShop.DistrictName,
			WardName:          barberShop.WardName,
			RoleId:            int32(barberShop.RoleID.Int16),
		}

		barberShops = append(barberShops, barberShopPB)
	}
	return barberShops
}

func ConvertSearchByNameBarberShops(res []db.SearchByNameBarberShopsRow) []*barber.BarberShop {
	var barberShops []*barber.BarberShop

	for _, barberShop := range res {
		// chainIDString := barberShop.BarberShopChainID.UUID.String()
		barberShopPB := &barber.BarberShop{
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

func ConvertListBarberShopsNearby(res []db.ListNearbyBarberShopsRow) []*barber.BarberShop {
	var barberShops []*barber.BarberShop
	for _, barberShop := range res {
		barberShopPB := &barber.BarberShop{
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
		GenderId:  int32(res.GenderID.Int16),
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
		GenderId:  int32(res.GenderID.Int16),
		Email:     res.Email.String,
		Phone:     res.Phone,
		NickName:  res.NickName,
		FullName:  res.FullName,
		Haircut:   res.Haircut,
		AvatarUrl: res.AvatarUrl.String,
	}
}
func convertBarberEmployee(res db.GetBarberRow) *barber.BarberDetail {

	barberEmployee := &barber.BarberDetail{
		Barber: &barber.Barbers{
			Id:         res.ID.String(),
			GenderId:   int32(res.GenderID.Int16),
			Email:      res.Email.String,
			Phone:      res.Phone,
			NickName:   res.NickName,
			FullName:   res.FullName,
			Haircut:    res.Haircut,
			WorkStatus: res.WorkStatus,
			AvatarUrl:  res.AvatarUrl.String,
		},
		BarberRole: &barber.BarberRoles{
			Id:             res.BarberID.String(),
			BarberId:       res.BarberID.String(),
			BarberShopId:   res.BarberShopID.String(),
			RoleId:         int32(res.RoleID),
			BarberShopName: res.BarberShopName,
		},
	}
	return barberEmployee
}

func convertBarberEmployees(res []db.ListEmployeesRow) []*barber.BarberDetail {
	var barberEmployees []*barber.BarberDetail

	for _, item := range res {
		barberEmployee := &barber.BarberDetail{
			Barber: &barber.Barbers{
				Id:         item.ID.String(),
				GenderId:   int32(item.GenderID.Int16),
				Email:      item.Email.String,
				Phone:      item.Phone,
				NickName:   item.NickName,
				FullName:   item.FullName,
				Haircut:    item.Haircut,
				WorkStatus: item.WorkStatus,
				AvatarUrl:  item.AvatarUrl.String,
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



func convertServiceCategory(serviceCategory db.ServiceCategory) *barber.ServiceCategory {
	return &barber.ServiceCategory{
		Id:           int32(serviceCategory.ID),
		Name:         serviceCategory.Name,
		BarberShopId: serviceCategory.BarberShopID.UUID.String(),
	}
}

func convertListServiceCategory(res []db.ServiceCategory) []*barber.ServiceCategory {
	var serviceCategories []*barber.ServiceCategory
	for _, item := range res {
		serviceCategory := &barber.ServiceCategory{
			Id:           int32(item.ID),
			Name:         item.Name,
			BarberShopId: item.BarberShopID.UUID.String(),
		}
		serviceCategories = append(serviceCategories, serviceCategory)
	}
	return serviceCategories
}



func convertBarberShopService(res db.BarberShopService) *barber.BarberShopService {
	return &barber.BarberShopService{
		Id:                res.ID.String(),
		CategoryId:        int32(res.CategoryID),
		BarberShopId:      res.BarberShopID.String(),
		GenderId:          int32(res.GenderID),
		Name:              res.Name,
		Timer:             int32(res.Timer),
		Price:             res.Price,
		Description:       res.Description.String,
		ImageUrl:          res.ImageUrl.String,
		ComboServices:     res.ComboServices,
		DiscountPrice:     &res.DiscountPrice.Float32,
		DiscountStartTime: timestamppb.New(res.DiscountStartTime.Time),
		DiscountEndTime:   timestamppb.New(res.DiscountEndTime.Time),
		IsActive:          res.IsActive,
	}
}

// func convertListBarberShopServices(res []db.ListBarberShopServicesRow) []*barber.ListItemBarberShopService {
// 	var services []*barber.ListItemBarberShopService
// 	for _, item := range res {
// 		service := &barber.ListItemBarberShopService{
// 			Id:           item.ID.String(),
// 			Name:         item.Name,
// 			CategoryName: item.CategoryName,
// 		}
// 		services = append(services, service)
// 	}
// 	return services
// }

func convertListService(res []db.ListServicesByCategoryRow) []*barber.BarberShopService {
	var services []*barber.BarberShopService
	for _, item := range res {
		service := &barber.BarberShopService{
			Id:                item.ServiceID.UUID.String(),
			CategoryId:        int32(item.CategoryID),
			GenderId:          int32(item.GenderID.Int16),
			Name:              item.ServiceName.String,
			Timer:             int32(item.Timer.Int16),
			Price:             item.Price.Float32,
			Description:       item.Description.String,
			ImageUrl:          item.ImageUrl.String,
			ComboServices:     item.ComboServices,
			IsActive:          item.IsActive.Bool,
			CategoryName:      item.CategoryName,
			DiscountPrice:     &item.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(item.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(item.DiscountEndTime.Time),
		}
		services = append(services, service)
	}
	return services
}

func convertListComboService(res []db.ListComboServicesRow) []*barber.BarberShopService {
	var services []*barber.BarberShopService
	for _, item := range res {
		service := &barber.BarberShopService{
			Id:                item.ServiceID.String(),
			GenderId:          int32(item.GenderID),
			Name:              item.ServiceName,
			Timer:             int32(item.Timer),
			Price:             item.Price,
			Description:       item.Description.String,
			ImageUrl:          item.ImageUrl.String,
			ComboServices:     item.ComboServices,
			IsActive:          item.IsActive,
			DiscountPrice:     &item.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(item.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(item.DiscountEndTime.Time),
		}
		services = append(services, service)
	}
	return services
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
