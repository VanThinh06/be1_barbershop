package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"

	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCustomer(ctx context.Context, req *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {

	validations := validateCreateCustomer(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid password")
	}

	if req.IsSocialAuth {
		hashedPassword = ""
		socialEmail, err := server.authVerifyJWTGG(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, "internal")
		}
		if socialEmail.Email != req.Email {
			return nil, status.Error(codes.Unauthenticated, "information is incorrect")
		}
	}

	arg := db.CreateCustomerParams{
		Name: req.GetName(),
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  req.Phone != nil,
		},
		Gender:         req.GetGender(),
		Email:          req.GetEmail(),
		IsSocialAuth:   req.GetIsSocialAuth(),
		HashedPassword: sql.NullString{String: hashedPassword, Valid: hashedPassword != ""},
	}

	res, err := server.store.CreateCustomer(ctx, arg)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Customers_pkey" {
					return nil, status.Errorf(codes.AlreadyExists, "this account has already existed")
				}
				if pqErr.Constraint == "Customers_phone_key" {
					return nil, status.Errorf(codes.AlreadyExists, "this phone has already existed")
				}
				if pqErr.Constraint == "Customers_email_key" {
					return nil, status.Errorf(codes.AlreadyExists, "this email has already existed")
				}
			}
		}
		return nil, status.Error(codes.InvalidArgument, "account creation failed. Please try again")
	}

	rsp := &customer.CreateCustomerResponse{
		Customer: convertCustomer(res),
	}
	return rsp, nil
}

func validateCreateCustomer(req *customer.CreateCustomerRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateEmail(req.Email); err != nil {
		validations = append(validations, FieldValidation("email", err))
	}

	if req.Phone != nil && *req.Phone != "" {
		if err := utils.ValidatePhoneNumber(req.GetPhone()); err != nil {
			validations = append(validations, FieldValidation("phone", err))
		}
	}

	if req.IsSocialAuth == false {
		if err := utils.ValidatePassword(req.GetPassword()); err != nil {
			validations = append(validations, FieldValidation("password", err))
		}
	}

	if err := utils.ValidateFullName(req.Name); err != nil {
		validations = append(validations, FieldValidation("name", err))
	}

	return validations
}
