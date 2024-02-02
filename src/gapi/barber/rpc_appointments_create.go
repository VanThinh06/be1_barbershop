package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateAppointments(ctx context.Context, req *barber.CreateAppointmentsRequest) (*barber.CreateAppointmentsResponse, error) {

	payload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	validations := validateCreateAppointment(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	errTx := make(chan error)
	resultAppointment := make(chan db.CreateAppointmentsRow)

	go func() {
		result, err := server.TxCreateAppointment(ctx, req)
		errTx <- err
		resultAppointment <- result
	}()

	err = <-errTx
	appointment := <-resultAppointment

	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	client, err := server.firebase.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize")
		return nil, status.Errorf(codes.Internal, "internal")
	}

	if appointment.BarberID == uuid.MustParse(req.BarberId) {
		registrationToken := appointment.FcmDevice.String
		message := &messaging.Message{
			Data: map[string]string{
				"appointment_date_time": appointment.AppointmentDateTime.String(),
			},
			Token: registrationToken,
			Notification: &messaging.Notification{
				Title: "Barber King",
				Body:  fmt.Sprintf("%s, Guests have just booked with you", appointment.BarberNickName),
			},
		}
		response, err := client.Send(ctx, message)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("Successfully sent message:", response)
	}

	if payload.Barber.BarberID == uuid.MustParse(req.BarberId) {
		registrationToken := payload.Barber.FcmDevice
		message := &messaging.Message{
			Data: map[string]string{
				"appointment_date_time": appointment.AppointmentDateTime.String(),
			},
			Token: registrationToken,
			Notification: &messaging.Notification{
				Title: "Barber King",
				Body:  "Guests have just booked with you",
			},
		}
		response, err := client.Send(ctx, message)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("Successfully sent message:", response)
	}

	rsp := &barber.CreateAppointmentsResponse{
		Appointment: convertAppointments(appointment),
		Services:    req.ServiceId,
	}
	return rsp, nil
}

func (server *Server) TxCreateAppointment(ctx context.Context, req *barber.CreateAppointmentsRequest) (db.CreateAppointmentsRow, error) {
	var resAppointment db.CreateAppointmentsRow
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		listServices, err := helpers.ConvertStringListToUUIDList(req.GetServiceId())
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "services do not exist")
		}

		arg := db.CreateAppointmentsParams{
			CustomerID:          uuid.MustParse(req.GetCustomerId()),
			BarberID:            uuid.MustParse(req.GetBarberId()),
			AppointmentDateTime: req.GetAppointmentDateTime().AsTime(),
		}

		resAppointment, err = q.CreateAppointments(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pgconn.PgError); ok {
				switch pqErr.ConstraintName {
				case "Appointments_barber_id_appointment_date_time_idx":
					return status.Errorf(codes.AlreadyExists, "The calendar is already booked. Please choose another time.")
				}
			}
			return status.Errorf(codes.Internal, "internal")
		}

		argServiceAppointment := db.CreateServicesForAppointmentsParams{
			AppointmentsServiceID: resAppointment.ID,
			ServicesID:            listServices,
		}

		_, err = q.CreateServicesForAppointments(ctx, argServiceAppointment)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "Services_Appointments_pkey":
					return status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
				}
			}
			return status.Errorf(codes.Internal, "internal")
		}

		return nil
	})
	return resAppointment, err
}

func validateCreateAppointment(req *barber.CreateAppointmentsRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	validateField := func(value *string, fieldName string, validateFunc func(string) error) {
		if value != nil {
			if err := validateFunc(*value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(&req.BarberId, "barber_id", helpers.ValidateId)
	validateField(&req.BarbershopId, "barbershop_id", helpers.ValidateId)
	validateField(&req.CustomerId, "customer_id", helpers.ValidateId)
	validateField(&req.ServiceId[0], "service_id", helpers.ValidateId)

	return validations
}
