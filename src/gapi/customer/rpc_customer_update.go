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

func (server *Server) UpdateCustomer(ctx context.Context, req *customer.UpdateCustomerRequest) (*customer.UpdateCustomerResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	validations := validateUpdateCustomer(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	if authPayload.Customer.CustomerID.String() != req.CustomerId {
		return nil, status.Errorf(codes.PermissionDenied, "permissionDenied")
	}

	arg := db.UpdateCustomerParams{
		ID: uuid.MustParse(req.CustomerId),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},

		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		// GenderID: sql.NullInt16{
		// 	Int16: int16(req.GetGender()),
		// 	Valid: req.Gender != nil,
		// },
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Avatar: sql.NullString{
			String: req.GetAvatar(),
			Valid:  req.Avatar != nil,
		},
	}

	res, err := server.store.UpdateCustomer(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "information is incorrect")
		}
		return nil, status.Errorf(codes.Internal, "internal")
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

	validateField(req.Email, "email", helpers.ValidateEmail)
	validateField(req.Phone, "phone", helpers.ValidatePhoneNumber)
	validateField(req.Password, "password", helpers.ValidatePassword)
	validateField(req.Name, "name", helpers.ValidateFullName)

	return validations
}
