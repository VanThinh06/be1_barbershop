package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) NewAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {

	// payload, err := server.AuthorizeUser(ctx)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	// }

	// err = server.IsManager(ctx, payload)
	// if err != nil {
	// 	return nil, status.Errorf(codes.PermissionDenied, "No permission")
	// }

	arg := db.CreateAppointmentParams{
		CustomerID:          uuid.MustParse(req.GetCustomerId()),
		BarberID:            uuid.MustParse(req.GetBarberId()),
		ServiceID:           uuid.MustParse(req.GetServiceId()),
		AppointmentDatetime: req.GetAppointmentDatetime().AsTime(),
		Status:              req.GetStatus(),
	}

	appointment, err := server.store.CreateAppointment(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "BarberShops_code_barber_shop_key":
				return nil, status.Errorf(codes.AlreadyExists, "This barber shop has already existed")
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create barber shop")
	}

	// Khởi tạo client Firebase Messaging
	client, err := server.firebaseApp.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Khởi tạo client Firebase Messaging thất bại: %v", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := "cwwkNuugTkWWRozOhJS9XK:APA91bHSIHVFiZmKkp5tTuRaSSpLRpSutkpyPoWOvLhZedk5zvBRxuyMk0WeQlog7HaV4K2EcKzPQgX0CGjlr48lwWbwfb_CMO8XSvWw9MdWiTVROCP9KtXwkllbMwBZICKvIhjY-pnz"

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"score":                "850",
			"time":                 "2:45",
			"message":              "Hello Work",
			"appointment_datetime": appointment.AppointmentDatetime.Format("2006-01-02T15:04:05Z"),
		},
		Token: registrationToken,
		Notification: &messaging.Notification{
			Title: "HelloWork",
			Body:  "Ok la nha",
		},
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

	rsp := &pb.CreateAppointmentResponse{
		Appointment: convertAppointment(appointment),
	}
	return rsp, nil
}
