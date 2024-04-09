package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
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

	if authPayload.Barber.BarberID.String() != req.Id {
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

		barberId, err := uuid.Parse(req.Id)
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
		Haircut: pgtype.Bool{
			Bool:  req.GetHaircut(),
			Valid: req.Haircut != nil,
		},
		GenderID: pgtype.Int2{
			Int16: int16(req.GetGenderId()),
			Valid: req.GenderId != nil,
		},
		WorkStatus: pgtype.Bool{
			Bool:  req.GetWorkStatus(),
			Valid: req.WorkStatus != nil,
		},

		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		AvatarUrl: sql.NullString{
			String: req.GetAvatarUrl(),
			Valid:  req.AvatarUrl != nil,
		},
	}

	res, err := server.Store.UpdateBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
		return nil, returnError(codes.PermissionDenied, "failed to update account", err)
	}

	rsp := &barber.UpdateBarberResponse{
		Barber: convertCreateBarbers(res),
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

	validateField(req.GetEmail(), "email", helpers.ValidateEmail)
	validateField(req.GetPhone(), "phone", helpers.ValidatePhoneNumber)
	validateField(req.GetNickname(), "nickname", helpers.ValidateNickName)
	validateField(req.GetFullName(), "full_name", helpers.ValidateFullName)

	return validations
}
