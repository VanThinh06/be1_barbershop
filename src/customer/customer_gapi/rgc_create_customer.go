package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"barbershop/src/shared/utils"
	"context"
	"net/http"

	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {

	validations := validateCreateCustomer(req)
	if validations != nil {
		return nil, pb.InValidArgumentError(validations)
	}

	// Hash the password provided in the request
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		// If there's an error hashing the password, return an unimplemented error
		return nil, status.Errorf(codes.Unimplemented, "failed to password %s", err)
	}

	// Prepare the arguments for creating a new barber
	arg := db.CreateCustomerParams{
		Name:           req.GetName(),
		Phone:          req.GetPhone(),
		Gender:         req.GetGender(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
	}

	// Create the barber using the store
	customer, err := server.store.CreateCustomer(ctx, arg)
	if err != nil {
		// If there's an error creating the barber, handle specific error cases
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Customers_pkey" {
					// If the barber account already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This account has already existed")
				}
				if pqErr.Constraint == "Customers_phone_key" {
					// If the phone number already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This phone has already existed")
				}
				if pqErr.Constraint == "Customers_email_key" {
					// If the email already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This email has already existed")
				}
			}
		}
		// If the error is not a specific case, return a forbidden error
		return nil, status.Errorf(http.StatusForbidden, "failed to create account %s", err)
	}

	// Prepare the response with the newly created barber
	rsp := &pb.CreateCustomerResponse{
		Customer: convertCustomer(customer),
	}
	return rsp, nil
}

func validateCreateCustomer(req *pb.CreateCustomerRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateEmail(req.Email); err != nil {
		validations = append(validations, pb.FieldValidation("email", err))
	}

	if err := utils.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, pb.FieldValidation("phone", err))
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		validations = append(validations, pb.FieldValidation("password", err))
	}

	if err := utils.ValidateFullName(req.Name); err != nil {
		validations = append(validations, pb.FieldValidation("name", err))
	}

	return validations
}
