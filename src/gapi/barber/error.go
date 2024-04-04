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

func inValidArgumentError(validations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: validations}
	msgError := "invalid " + validations[0].Field + " format"
	statusInvalid := status.New(codes.InvalidArgument, msgError)

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

func noPermissionError(err error) error {
	return status.Error(codes.PermissionDenied, "no permission")
}

func invalidInformationError(err error) error {
	return status.Errorf(codes.InvalidArgument, "invalid information")
}

func notFoundError(err error, strError string) error {
	return status.Errorf(codes.NotFound, "not found %v", strError)
}
