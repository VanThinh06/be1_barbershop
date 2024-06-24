package utils

import (
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func ValidateBarber(req *barber.CreateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidatePassword(req.Password); err != nil {
		validations = append(validations, FieldValidation("password", err))
	}

	if err := helpers.ValidateEmail(req.Email); err != nil {
		validations = append(validations, FieldValidation("email", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}
	return validations
}


func ValidateBarberEmployee(req *barber.BarberEmployee) (validations []*errdetails.BadRequest_FieldViolation) {

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}


func ValidateBarberUpdate(req *barber.UpdateBarberRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	validateField := func(value, fieldName string, validateFunc func(string) error) {
		if value != "" {
			if err := validateFunc(value); err != nil {
				validations = append(validations, FieldValidation(fieldName, err))
			}
		}
	}

	validateField(req.Barber.GetEmail(), "email", helpers.ValidateEmail)
	validateField(req.Barber.GetPhone(), "phone", helpers.ValidatePhoneNumber)
	validateField(req.Barber.GetFullName(), "full_name", helpers.ValidateFullName)

	return validations
}
