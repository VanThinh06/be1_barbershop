package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCustomer(res db.Customer) *customer.Customer {
	return &customer.Customer{
		Phone:             res.Phone.String,
		Email:             res.Email,
		Gender:            res.Gender,
		Name:              res.Name,
		Avatar:            res.Avatar.String,
		CustomerId:        res.CustomerID.String(),
		CreatedAt:         timestamppb.New(res.CreatedAt),
		UpdateAt:          timestamppb.New(res.UpdatedAt.Time),
		PasswordChangedAt: timestamppb.New(res.PasswordChangedAt),
		IsSocialAuth:      res.IsSocialAuth.Bool,
	}
}

func convertAppointment(appointment db.Appointment) *customer.Appointment {
	return &customer.Appointment{
		AppointmentId:       appointment.AppointmentID.String(),
		CustomerId:          appointment.CustomerID.String(),
		BarberId:            appointment.BarberID.String(),
		AppointmentDatetime: timestamppb.New(appointment.AppointmentDatetime),
		Status:              appointment.Status,
		CreatedAt:           timestamppb.New(appointment.CreatedAt),
		UpdateAt:            timestamppb.New(appointment.UpdatedAt.Time),
	}
}

func convertGetAppointmentByDate(res []db.Appointment) []*customer.Appointment {
	var appointments []*customer.Appointment
	for _, appointment := range res {
		appointment := &customer.Appointment{
			AppointmentId:       appointment.AppointmentID.String(),
			CustomerId:          appointment.CustomerID.String(),
			BarberId:            appointment.BarbershopsID.String(),
			Status:              appointment.Status,
			AppointmentDatetime: timestamppb.New(appointment.AppointmentDatetime),
			CreatedAt:           timestamppb.New(appointment.CreatedAt),
			UpdateAt:            timestamppb.New(appointment.UpdatedAt.Time),
		}
		appointments = append(appointments, appointment)
	}
	return appointments
}

func convertGetBarberInShop(res []db.Barber) []*customer.Barbers {
	var barbers []*customer.Barbers
	for _, item := range res {
		barber := &customer.Barbers{
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
