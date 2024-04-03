package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest) (*barber.CreateBarberEmployeeResponse, error) {

	// TODO
	_, err := server.authorizeBarber(ctx)
	if err != nil {
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

	hashedPassword, err := helpers.HashPassword(defaultPassword)
	if err != nil {
		return nil, internalError(err)
	}

	newUUID := uuid.New()
	uuidPrefix := newUUID.String()[:8]
	combinedNickName := req.NickName + uuidPrefix

	arg := db.CreateBarberEmployeeParams{
		Phone:          req.GetPhone(),
		HashedPassword: hashedPassword,
		NickName:       combinedNickName,
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  true,
		},
	}

	errTx := make(chan error)
	resultBarber := make(chan db.Barber)
	resultBarberRole := make(chan db.BarberRole)

	go func() {
		barber, role, err := server.txCreateBarberEmployee(ctx, req, arg)
		errTx <- err
		resultBarber <- barber
		resultBarberRole <- role
	}()

	err = <-errTx
	rspBarber := <-resultBarber
	_ = <-resultBarberRole

	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.CreateBarberEmployeeResponse{
		Barber: convertCreateBarbers(rspBarber),
	}
	return rsp, nil
}

func (server *Server) txCreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest, arg db.CreateBarberEmployeeParams) (db.Barber, db.BarberRole, error) {
	var resBarber db.Barber
	var resBarberRole db.BarberRole

	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
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
		resBarberRole, err = server.Store.CreateBarberRoles(ctx, argBarberRole)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		return nil
	})
	return resBarber, resBarberRole, err
}

func validateCreateBarberEmployee(req *barber.CreateBarberEmployeeRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}
