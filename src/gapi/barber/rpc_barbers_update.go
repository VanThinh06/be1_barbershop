package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarber(ctx context.Context, req *barber.UpdateBarberRequest) (*barber.UpdateBarberResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, UnauthenticatedError(err)
	}

	validations := validateUpdateBarber(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	if authPayload.Barber.BarberID.String() != req.BarberId {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
	}

	arg := db.UpdateBarberParams{
		BarberID: uuid.MustParse(req.GetBarberId()),
		NickName: sql.NullString{
			String: req.GetNickname(),
			Valid:  req.Nickname != nil,
		},
		Status: sql.NullInt32{
			Int32: req.GetStatus(),
			Valid: req.Status != nil,
		},
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},

		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Gender: sql.NullInt32{
			Int32: req.GetGender(),
			Valid: req.Gender != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Avatar: sql.NullString{
			String: req.GetAvatar(),
			Valid:  req.Avatar != nil,
		},
		UpdatedAt: time.Now(),
	}

	// update shop id
	if req.ShopId != nil {
		arg.ShopID = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetShopId()),
			Valid: true,
		}
	}

	// Hash the password provided in the request
	if req.Password != nil {
		hashedPassword, err := utils.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password")
		}

		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}

		arg.PasswordChangedAt = sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
	}

	res, err := server.Store.UpdateBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "barber not found")
		}
		return nil, status.Errorf(codes.PermissionDenied, "failed to update account")
	}

	rsp := &barber.UpdateBarberResponse{
		Barber: convertBarber(res),
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

	validateField(req.GetEmail(), "email", utils.ValidateEmail)
	validateField(req.GetPhone(), "phone", utils.ValidatePhoneNumber)
	validateField(req.GetPassword(), "password", utils.ValidatePassword)
	validateField(req.GetNickname(), "nickname", utils.ValidateNickname)
	validateField(req.GetFullName(), "full_name", utils.ValidateFullName)

	return validations
}
