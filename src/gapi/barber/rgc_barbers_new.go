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

func (server *Server) CreateBarber(ctx context.Context, req *barber.CreateBarberRequest) (*barber.CreateBarberResponse, error) {

	validations := validateCreateBarber(req)
	if validations != nil {
		return nil, InValidArgumentError(validations)
	}

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	var managerId uuid.NullUUID
	var barberShopId uuid.NullUUID
	var roleBarber = int(req.GetRole())
	if req.Role != int32(utils.Manager) {
		if req.Role == int32(utils.HairStylist) {
			req.Haircut = true
		}
		var code, err = token.DecodeAESString(server.config.AesKey, req.CodeBarberShop)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "information is incorrect")
		}
		var listCode = splitExpression(code)

		strRoleBarber, err := utils.GetValueAtIndex(listCode, 0)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "information is incorrect")
		}
		roleBarber, err = utils.ConvertStringToInt(strRoleBarber)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "information is incorrect")
		}

		valueIdBarberShop, err := utils.GetValueAtIndex(listCode, 1)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "information is incorrect")
		}
		if valueIdBarberShop != "" {
			barberShopId = uuid.NullUUID{
				UUID:  uuid.MustParse(valueIdBarberShop),
				Valid: valueIdBarberShop != "",
			}
		}

		valueIdManager, err := utils.GetValueAtIndex(listCode, 2)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "information is incorrect")
		}
		if valueIdManager != "" {
			managerId = uuid.NullUUID{
				UUID:  uuid.MustParse(valueIdManager),
				Valid: valueIdManager != "",
			}
		}
	}

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

	res, err := server.Store.CreateBarber(ctx, arg)
	if err != nil {

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

		return nil, status.Errorf(codes.Internal, "internal")
	}

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

	parts := strings.Split(expression, "+")

	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
}
