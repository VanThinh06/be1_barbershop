package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarber(ctx context.Context, req *barber.UpdateBarbersRequest) (*barber.UpdateBarbersResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	validations := validateUpdateBarber(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	if authPayload.Barber.BarberID.String() != req.Id {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
	}

	arg := db.UpdateBarbersParams{
		ID: uuid.MustParse(req.GetId()),
		NickName: sql.NullString{
			String: req.GetNickname(),
			Valid:  req.Nickname != nil,
		},

		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},

		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		// GenderID: sql.NullInt16{
		// 	Int16: int16(req.GetGenderId()),
		// 	Valid: req.GenderId != nil,
		// },
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		AvatarUrl: sql.NullString{
			String: req.GetAvatarUrl(),
			Valid:  req.AvatarUrl != nil,
		},
	}

	res, err := server.Store.UpdateBarbers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
		return nil, returnError(codes.PermissionDenied, "failed to update account", err)
	}

	rsp := &barber.UpdateBarbersResponse{
		Barber: convertCreateBarbers(res),
	}
	return rsp, nil
}

func validateUpdateBarber(req *barber.UpdateBarbersRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value, fieldName string, validateFunc func(string) error) {
		if value != "" {
			if err := validateFunc(value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(req.GetEmail(), "email", helpers.ValidateEmail)
	validateField(req.GetPhone(), "phone", helpers.ValidatePhoneNumber)
	validateField(req.GetNickname(), "nickname", helpers.ValidateNickName)
	validateField(req.GetFullName(), "full_name", helpers.ValidateFullName)

	return validations
}
