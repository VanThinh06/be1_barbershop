package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/utils"
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

func (server *Server) NewAppointment(ctx context.Context, req *customer.CreateAppointmentRequest) (*customer.CreateAppointmentResponse, error) {

	payload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	validations := validateCreateAppointment(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	tx, err := server.store.
	listServices, err := utils.ConvertStringListToUUIDList(req.GetServiceId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "services does not exist")
	}
	timer, err := server.store.GetTimerService(ctx, listServices)

	arg := db.InsertAppointmentAndGetInfoParams{
		BarbershopsID:       uuid.MustParse(req.GetBarbershopId()),
		CustomerID:          uuid.MustParse(req.GetCustomerId()),
		BarberID:            uuid.MustParse(req.GetBarberId()),
		AppointmentDatetime: req.GetAppointmentDatetime().AsTime(),
		Timer:               int32(timer),
		Status:              int32(utils.Pending),
	}
	resAppoinment, err := server.store.InsertAppointmentAndGetInfo(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "Appointments_barber_id_appointment_datetime_idx":
				return nil, status.Errorf(codes.AlreadyExists, "The calendar is already booked. Please choose another time.")
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	argServiceAppointment := db.InsertServicesForAppointmentParams{
		AppointmentsServiceID: resAppoinment.AppointmentID,
		ServicesID:            listServices,
	}
	_, err = server.store.InsertServicesForAppointment(ctx, argServiceAppointment)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "Services_Appointments_pkey":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	// firebase notification
	client, err := server.firebaseApp.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize")
		return nil, status.Errorf(codes.Internal, "internal")
	}

	if resAppoinment.BarberID == uuid.MustParse(req.BarberId) {
		registrationToken := resAppoinment.FcmDevice.String
		message := &messaging.Message{
			Data: map[string]string{
				"appointment_datetime": resAppoinment.AppointmentDatetime.String(),
			},
			Token: registrationToken,
			Notification: &messaging.Notification{
				Title: "Barber King",
				Body:  fmt.Sprintf("%s, Guests have just booked with you", resAppoinment.BarberNickName),
			},
		}
		response, err := client.Send(ctx, message)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("Successfully sent message:", response)
	}

	if payload.Customer.CustomerID == uuid.MustParse(req.CustomerId) {
		registrationToken := payload.Customer.FcmDevice
		message := &messaging.Message{
			Data: map[string]string{
				"appointment_datetime": resAppoinment.AppointmentDatetime.String(),
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

	rsp := &customer.CreateAppointmentResponse{
		Appointment: convertAppointment(resAppoinment),
	}
	return rsp, nil
}

func validateCreateAppointment(req *customer.CreateAppointmentRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value *string, fieldName string, validateFunc func(string) error) {
		if value != nil {
			if err := validateFunc(*value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(&req.BarberId, "barber_id", utils.ValidateId)
	validateField(&req.BarbershopId, "barbershop_id", utils.ValidateId)
	validateField(&req.CustomerId, "customer_id", utils.ValidateId)
	validateField(&req.ServiceId[0], "service_id", utils.ValidateId)

	return validations
}
