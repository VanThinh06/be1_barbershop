package gapi

import (
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GenerateCodeBarberShop(ctx context.Context, req *barber.CodeBarberShopRequest) (*barber.CodeBarberShopResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	res, err := server.Store.BarberGetIdShop(ctx, payload.Barber.BarberID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	resultString := fmt.Sprintf("%d + %s + %s", req.GetRole(), res.ShopID, payload.Barber.BarberID.String())
	codeBarberShop, err := token.GenerateAESString(server.config.AesKey, resultString)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.CodeBarberShopResponse{
		Code: codeBarberShop,
	}
	return rsp, nil
}

func validateGenerateCode(req *barber.CodeBarberShopRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value string, fieldName string, validateFunc func(string) error) {
		if value != "" && value != "0" {
			if err := validateFunc(value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(string(req.GetRole()), "role", utils.ValidateNumber)
	return validations
}
