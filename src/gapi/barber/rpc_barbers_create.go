package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"

	"github.com/jackc/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func (server *Server) CreateBarber(ctx context.Context, req *barber.CreateBarbersRequest) (*barber.CreateBarbersResponse, error) {

	validations := validateCreateBarberAccountBarberShop(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	hashedPassword, err := utilities.HashPassword(req.GetPassword())
	if err != nil {
		return nil, internalError(err)
	}

	arg := db.CreateBarbersParams{
		NickName:       req.GetNickName(),
		HashedPassword: hashedPassword,
		Phone:          req.GetPhone(),
		GenderID:       req.GetGenderId(),
		Email:          req.GetEmail(),
	}
	res, err := server.Store.CreateBarbers(ctx, arg)
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

	rsp := &barber.CreateBarbersResponse{
		Barber: convertCreateBarbers(res),
	}
	return rsp, nil
}

func validateCreateBarberAccountBarberShop(req *barber.CreateBarbersRequest) (validations []*errdetails.BadRequest_FieldViolation) {
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
		validations = append(validations, FieldValidation("nick_name", err))
	}

	return validations
}
