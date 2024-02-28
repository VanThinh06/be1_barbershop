package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/helpers"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ChangePasswordCustomer(ctx context.Context, req *customer.ChangePasswordCustomerRequest) (*customer.ChangePasswordCustomerResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	validations := validateChangePasswordCustomer(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	if authPayload.Customer.CustomerID.String() != req.IdCustomer {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
	}

	cus, err := server.store.GetCustomer(ctx, authPayload.Customer.CustomerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}
	err = helpers.CheckPassword(req.GetCurrentPassword(), cus.HashedPassword.String)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	// Hash the password provided in the request
	hashedPassword, err := helpers.HashPassword(req.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	arg := db.ChangePasswordCustomerParams{
		ID:             uuid.MustParse(req.IdCustomer),
		HashedPassword: hashedPassword,
	}

	_, err = server.store.ChangePasswordCustomer(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "barber not found")
		}
		return nil, status.Errorf(codes.PermissionDenied, "failed to update account")
	}

	rsp := &customer.ChangePasswordCustomerResponse{
		Status: true,
	}
	return rsp, nil
}

func validateChangePasswordCustomer(req *customer.ChangePasswordCustomerRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value *string, fieldName string, validateFunc func(string) error) {
		if value != nil {
			if err := validateFunc(*value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(&req.CurrentPassword, "current_password", helpers.ValidatePassword)
	validateField(&req.NewPassword, "new_password", helpers.ValidatePassword)
	return validations
}
