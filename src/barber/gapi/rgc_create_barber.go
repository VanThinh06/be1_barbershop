package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"barbershop/src/shared/utils"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarber(ctx context.Context, req *pb.CreateBarberRequest) (*pb.CreateBarberResponse, error) {

	validations := validateCreateBarber(req)
	if validations != nil {
		return nil, pb.InValidArgumentError(validations)
	}

	// Hash the password provided in the request
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		// If there's an error hashing the password, return an unimplemented error
		return nil, status.Errorf(codes.Unimplemented, "failed to password %s", err)
	}

	// Create a channel to receive the result of fetching the codeBarberShop
	resultChan := make(chan uuid.NullUUID)
	var codeBarberShop uuid.NullUUID

	// Fetch the codeBarberShop asynchronously and send the result to the channel
	if req.Role != int32(utils.Manager) {
		go func() {
			uuidBarberShop, err := server.store.GetCodeBarberShop(ctx, req.CodeBarberShop)
			if err != nil {
				resultChan <- uuid.NullUUID{Valid: false}
			} else {
				resultChan <- uuid.NullUUID{Valid: true, UUID: uuidBarberShop}
			}
		}()
		codeBarberShop = <-resultChan

		if !codeBarberShop.Valid {
			// If there's an error fetching the codeBarberShop, return an internal server error
			return nil, status.Errorf(codes.NotFound, "code barber shop not found")
		}
	}

	// Prepare the arguments for creating a new barber
	arg := db.CreateBarberParams{
		NickName:       req.GetNickname(),
		HashedPassword: hashedPassword,
		Phone:          req.GetPhone(),
		FullName:       req.GetEmail(),
		Gender:         req.GetGender(),
		Email:          req.GetEmail(),
		ShopID:         codeBarberShop,
		Role:           req.GetRole(),
	}
	// Create the barber using the store
	barber, err := server.store.CreateBarber(ctx, arg)
	if err != nil {
		// If there's an error creating the barber, handle specific error cases
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Barbers_pkey" {
					// If the barber account already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This account has already existed")
				}
				if pqErr.Constraint == "Barbers_phone_key" {
					// If the phone number already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This phone has already existed")
				}
				if pqErr.Constraint == "Barbers_email_key" {
					// If the email already exists, return an already exists error
					return nil, status.Errorf(codes.AlreadyExists, "This email has already existed")
				}
			}
		}
		// If the error is not a specific case, return a forbidden error
		return nil, status.Errorf(codes.Internal, "failed to create account %s", err)
	}

	// Prepare the response with the newly created barber
	rsp := &pb.CreateBarberResponse{
		Barber: convertBarber(barber),
	}
	return rsp, nil
}

func validateCreateBarber(req *pb.CreateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateEmail(req.Email); err != nil {
		validations = append(validations, pb.FieldValidation("email", err))
	}

	if err := utils.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, pb.FieldValidation("phone", err))
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		validations = append(validations, pb.FieldValidation("password", err))
	}

	if err := utils.ValidateNickname(req.Nickname); err != nil {
		validations = append(validations, pb.FieldValidation("nickname", err))
	}

	if err := utils.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, pb.FieldValidation("full_name", err))
	}

	return validations
}
