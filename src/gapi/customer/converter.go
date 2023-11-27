package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCustomer(res db.Customer) *customer.Customer {
	return &customer.Customer{
		Phone:             res.Phone,
		Email:             res.Email,
		Gender:            res.Gender,
		Name:              res.Name,
		Avatar:            res.Avatar.String,
		CustomerId:        res.CustomerID.String(),
		CreatedAt:         timestamppb.New(res.CreateAt),
		UpdateAt:          timestamppb.New(res.UpdateAt.Time),
		PasswordChangedAt: timestamppb.New(res.PasswordChangedAt),
	}
}

func convertAppointment(appointment db.Appointment) *customer.Appointment {
	return &customer.Appointment{
		AppointmentId:       appointment.AppointmentID.String(),
		CustomerId:          appointment.CustomerID.String(),
		BarberId:            appointment.BarberID.String(),
		ServiceId:           appointment.ServiceID.String(),
		AppointmentDatetime: timestamppb.New(appointment.AppointmentDatetime),
		Status:              appointment.Status,
		CreatedAt:           timestamppb.New(appointment.CreatedAt),
		UpdateAt:            timestamppb.New(appointment.UpdateAt.Time),
	}
}
