package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"context"
	"database/sql"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func (server *Server) CreateBarber(ctx context.Context, req *barber.CreateBarberRequest) (*barber.CreateBarberResponse, error) {

	validations := validateCreateBarberAccountBarberShop(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	hashedPassword, err := helpers.HashPassword(req.GetPassword())
	if err != nil {
		return nil, internalError(err)
	}

	arg := db.CreateBarberParams{
		NickName: strings.ToLower(req.GetNickName()),
		FullName: req.GetFullName(),
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		Phone:    req.GetPhone(),
		GenderID: int16(req.GetGenderId()),
		Email: sql.NullString{
			String: strings.ToLower(req.GetEmail()),
			Valid:  true,
		},
	}
	res, err := server.Store.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "Barbers_pkey":
				return nil, returnError(codes.AlreadyExists, "This id has already existed", err)
			case "Barbers_email_key":
				return nil, returnError(codes.AlreadyExists, "This email has already existed", err)
			case "Barbers_phone_key":
				return nil, returnError(codes.AlreadyExists, "This phone has already existed", err)
			case "Barbers_nick_name_key":
				return nil, returnError(codes.AlreadyExists, "This nick name has already existed", err)
			}
		}
		return nil, internalError(err)
	} 

	rsp := &barber.CreateBarberResponse{
		Barber: convertCreateBarbers(res),
	}
	return rsp, nil
}

func validateCreateBarberAccountBarberShop(req *barber.CreateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := helpers.ValidateEmail(req.Email); err != nil {
		validations = append(validations, FieldValidation("email", err))
	}

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidatePassword(req.Password); err != nil {
		validations = append(validations, FieldValidation("password", err))
	}

	if err := helpers.ValidateNickName(req.NickName); err != nil {
		validations = append(validations, FieldValidation("nick name", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full name", err))
	}
	return validations
}
