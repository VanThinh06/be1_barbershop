package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FieldValidation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}

func InValidArgumentError(validations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: validations}
	statusInvalid := status.New(codes.InvalidArgument, "invalid parameters")

	statusDetail, err := statusInvalid.WithDetails(badRequest)

	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetail.Err()
}

func unauthenticatedError(err error) error {
	return status.Error(codes.Unauthenticated, "unauthenticated")
}

func internalError(err error) error {
	return status.Error(codes.Internal, "internal")
}

func returnError(code codes.Code, strError string, err error) error {
	return status.Error(code, strError)
}
