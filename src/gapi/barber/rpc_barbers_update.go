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
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarber(ctx context.Context, req *barber.UpdateBarberRequest) (*barber.UpdateBarberResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	validations := validateUpdateBarber(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	if authPayload.Barber.BarberID.String() != req.Barber.Id {
		barberShopId, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
		}
		argBarberRoleAdmin := db.GetBarberRolesParams{
			BarberID:     authPayload.Barber.BarberID,
			BarberShopID: barberShopId,
		}
		barberRoleAmin, err := server.Store.GetBarberRoles(ctx, argBarberRoleAdmin)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
		}
		if barberRoleAmin.RoleID != int16(utilities.Admin) {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
		}

		barberId, err := uuid.Parse(req.Barber.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "barber don't exist")
		}
		argBarberRole := db.GetBarberRolesParams{
			BarberID:     barberId,
			BarberShopID: barberShopId,
		}
		_, err = server.Store.GetBarberRoles(ctx, argBarberRole)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
		}
	}

	arg := db.UpdateBarberParams{
		ID: uuid.MustParse(req.Barber.GetId()),

		NickName: sql.NullString{
			String: strings.ToLower(req.Barber.GetNickname()),
			Valid:  req.Barber.Nickname != nil,
		},

		FullName: sql.NullString{
			String: req.Barber.GetFullName(),
			Valid:  req.Barber.FullName != nil,
		},

		Phone: sql.NullString{
			String: req.Barber.GetPhone(),
			Valid:  req.Barber.Phone != nil,
		},
		Haircut: pgtype.Bool{
			Bool:  req.Barber.GetHaircut(),
			Valid: req.Barber.Haircut != nil,
		},
		GenderID: pgtype.Int2{
			Int16: int16(req.Barber.GetGenderId()),
			Valid: req.Barber.GenderId != nil,
		},
		WorkStatus: pgtype.Bool{
			Bool:  req.Barber.GetWorkStatus(),
			Valid: req.Barber.WorkStatus != nil,
		},

		Email: sql.NullString{
			String: strings.ToLower(req.Barber.GetEmail()),
			Valid:  req.Barber.Email != nil,
		},
		AvatarUrl: sql.NullString{
			String: req.Barber.GetAvatarUrl(),
			Valid:  req.Barber.AvatarUrl != nil,
		},
	}

	_, err = server.Store.UpdateBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
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

	rsp := &barber.UpdateBarberResponse{
		Message: "Your information has been successfully updated.",
	}
	return rsp, nil
}

func validateUpdateBarber(req *barber.UpdateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value, fieldName string, validateFunc func(string) error) {
		if value != "" {
			if err := validateFunc(value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(req.Barber.GetEmail(), "email", helpers.ValidateEmail)
	validateField(req.Barber.GetPhone(), "phone", helpers.ValidatePhoneNumber)
	validateField(req.Barber.GetNickname(), "nickname", helpers.ValidateNickName)
	validateField(req.Barber.GetFullName(), "full_name", helpers.ValidateFullName)

	return validations
}
