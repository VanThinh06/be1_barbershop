package utils

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

func InvalidArgumentError(validations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: validations}
	msgError := "invalid " + validations[0].Field + " format"
	statusInvalid := status.New(codes.InvalidArgument, msgError)

	statusDetail, err := statusInvalid.WithDetails(badRequest)
	
	if err != nil {
		return statusInvalid.Err()
	}

	return statusDetail.Err()
}

func UnauthenticatedError(err error) error {
	return status.Error(codes.Unauthenticated, "unauthenticated")
}

func InternalError(err error) error {
	return status.Error(codes.Internal, "internal")
}
func NoPermissionError(err error) error {
	return status.Error(codes.PermissionDenied, "no permission")
}

func NotFoundError(err error, strError string) error {
	return status.Errorf(codes.NotFound, "not found %v", strError)
}

func InvalidError(err error, strError string) error {
	return status.Errorf(codes.InvalidArgument, strError)
}

func AlreadyExistsError(err error, strError string) error {
	return status.Error(codes.AlreadyExists, strError)
}

func ReturnError(code codes.Code, err error, strError string) error {
	return status.Error(code, strError)
}

func ForeignKeyError(err error, strError string) error {
	return status.Error(codes.FailedPrecondition, strError)
}