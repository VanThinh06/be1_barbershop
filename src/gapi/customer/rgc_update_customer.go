package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateCustomer(ctx context.Context, req *customer.UpdateCustomerRequest) (*customer.UpdateCustomerResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, UnauthenticatedError(err)
	}

	validations := validateUpdateCustomer(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	if authPayload.Customer.CustomerID.String() != req.CustomerId {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
	}

	arg := db.UpdateCustomerParams{
		CustomerID: uuid.MustParse(req.CustomerId),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
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
		UpdateAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
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

	res, err := server.store.UpdateCustomer(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "barber not found")
		}
		return nil, status.Errorf(codes.PermissionDenied, "failed to update account")
	}

	rsp := &customer.UpdateCustomerResponse{
		Customer: convertCustomer(res),
	}
	return rsp, nil
}

func validateUpdateCustomer(req *customer.UpdateCustomerRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value *string, fieldName string, validateFunc func(string) error) {
		if value != nil {
			if err := validateFunc(*value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(req.Email, "email", utils.ValidateEmail)
	validateField(req.Phone, "phone", utils.ValidatePhoneNumber)
	validateField(req.Password, "password", utils.ValidatePassword)
	validateField(req.Name, "name", utils.ValidateFullName)

	return validations
}
