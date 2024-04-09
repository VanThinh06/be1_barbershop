package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest) (*barber.CreateBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberRoleType != string(utilities.Administrator) {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	validations := validateCreateBarberEmployee(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	defaultPassword, err := server.Store.GetDefaultPasswordEmployee(ctx, barberShopID)
	if err != nil {
		return nil, internalError(err)
	}

	var hashedPassword string
	if defaultPassword.Valid {
		hashedPassword, err = helpers.HashPassword(defaultPassword.String)
		if err != nil {
			return nil, internalError(err)
		}
	}

	newUUID := uuid.New()
	uuidPrefix := newUUID.String()[:8]
	trimNickName := strings.ReplaceAll(req.BarberEmployee.NickName, " ", "")
	combinedNickName := strings.TrimSpace(trimNickName) + uuidPrefix

	var hashedPasswordValid bool = hashedPassword != ""
	arg := db.CreateBarberEmployeeParams{
		Phone: req.BarberEmployee.GetPhone(),
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  hashedPasswordValid,
		},
		NickName: combinedNickName,
		FullName: req.BarberEmployee.GetFullName(),
	}

	errTx := make(chan error)
	message := make(chan string)

	go func() {
		str, err := server.txCreateBarberEmployee(ctx, req, arg)
		errTx <- err
		message <- str
	}()

	err = <-errTx
	response := <-message
	

	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberEmployeeResponse{
		Message: response,
	}
	return rsp, nil
}

func (server *Server) txCreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest, arg db.CreateBarberEmployeeParams) (string, error) {

	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		resBarber, err := server.Store.CreateBarberEmployee(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pgconn.PgError); ok {
				switch pqErr.ConstraintName {
				case "Barbers_pkey":
					return returnError(codes.AlreadyExists, "This id has already existed", err)
				case "Barbers_email_key":
					return returnError(codes.AlreadyExists, "This email has already existed", err)
				case "Barbers_phone_key":
					return returnError(codes.AlreadyExists, "This phone has already existed", err)
				case "Barbers_nick_name_key":
					return returnError(codes.AlreadyExists, "This nick name has already existed", err)
				}
			}
			return internalError(err)
		}

		barberShopID, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "barbershops don't exist")
		}
		argBarberRole := db.CreateBarberRolesParams{
			BarberID:     resBarber.ID,
			BarberShopID: barberShopID,
			RoleID:       int16(req.RoleId),
		}
		_, err = server.Store.CreateBarberRoles(ctx, argBarberRole)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		return nil
	})
	return "", err
}

func validateCreateBarberEmployee(req *barber.CreateBarberEmployeeRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	if err := helpers.ValidatePhoneNumber(req.BarberEmployee.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidateFullName(req.BarberEmployee.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}
