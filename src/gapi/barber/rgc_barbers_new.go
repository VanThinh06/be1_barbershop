package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatxeBarber(ctx context.Context, req *barber.CreateBarberRequest) (*barber.CreateBarberResponse, error) {

	validations := validateCreateBarber(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	// Hash the password provided in the request
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, "failed to password")
	}

	var managerId uuid.NullUUID
	var barberShopId uuid.NullUUID
	var roleBarber = int(req.GetRole())
	if req.Role != int32(utils.Manager) {
		var code, err = token.DecodeAESString(server.config.AesKey, req.CodeBarberShop)
		if err != nil {
			return nil, status.Error(codes.Unimplemented, "failed to password")
		}
		var listCode = splitExpression(code)

		strRoleBarber, err := utils.GetValueAtIndex(listCode, 0)
		if err != nil {
			return nil, status.Error(codes.Unimplemented, "failed to password")
		}
		roleBarber, err = utils.ConvertStringToInt(strRoleBarber)
		if err != nil {
			return nil, status.Error(codes.Unimplemented, "failed to password")
		}

		valueIdBarberShop, err := utils.GetValueAtIndex(listCode, 1)
		if err != nil {
			return nil, status.Error(codes.Unimplemented, "failed to password")
		}
		if valueIdBarberShop != "" {
			barberShopId = uuid.NullUUID{
				UUID:  uuid.MustParse(valueIdBarberShop),
				Valid: valueIdBarberShop != "",
			}
		}
		valueIdManager, err := utils.GetValueAtIndex(listCode, 2)
		if err != nil {
			return nil, status.Error(codes.Unimplemented, "failed to password")
		}
		if valueIdManager != "" {
			managerId = uuid.NullUUID{
				UUID:  uuid.MustParse(valueIdManager),
				Valid: valueIdManager != "",
			}
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
		ShopID:         barberShopId,
		Role:           int32(roleBarber),
		ManagerID:      managerId,
		Haircut:        req.Haircut,
	}
	// Create the barber using the store
	res, err := server.Store.CreateBarber(ctx, arg)
	if err != nil {
		// If there's an error creating the barber, handle specific error cases
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Barbers_pkey" {
					return nil, status.Errorf(codes.AlreadyExists, "This account has already existed")
				}
				if pqErr.Constraint == "Barbers_phone_key" {
					return nil, status.Errorf(codes.AlreadyExists, "This phone has already existed")
				}
				if pqErr.Constraint == "Barbers_email_key" {
					return nil, status.Errorf(codes.AlreadyExists, "This email has already existed")
				}
			}
		}

		// If the error is not a specific case, return a forbidden error
		return nil, status.Errorf(codes.Internal, "internal")
	}

	// Prepare the response with the newly created barber
	rsp := &barber.CreateBarberResponse{
		Barber: convertBarber(res),
	}
	return rsp, nil
}

func validateCreateBarber(req *barber.CreateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateEmail(req.Email); err != nil {
		validations = append(validations, FieldValidation("email", err))
	}

	if err := utils.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := utils.ValidatePassword(req.Password); err != nil {
		validations = append(validations, FieldValidation("password", err))
	}

	if err := utils.ValidateNickname(req.Nickname); err != nil {
		validations = append(validations, FieldValidation("nickname", err))
	}

	if err := utils.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}

func splitExpression(expression string) []string {
	// Tách chuỗi bằng dấu cộng
	parts := strings.Split(expression, "+")

	// Loại bỏ khoảng trắng xung quanh mỗi phần tử
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
}
