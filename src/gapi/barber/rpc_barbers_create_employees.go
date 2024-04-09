package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberEmployees(ctx context.Context, req *barber.CreateBarberEmployeesRequest) (*barber.CreateBarberEmployeesResponse, error) {
	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberRoleType != string(utilities.Administrator) {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	defaultPassword, err := server.Store.GetDefaultPasswordEmployee(ctx, barberShopID)
	if err != nil {
		return nil, internalError(err)
	}
	var hashedPassword string
	if defaultPassword.Valid {
		hashedPassword, err = helpers.HashPassword(defaultPassword.String)
		if err != nil {
			return nil, internalError(err)
		}
	}

	var responses string = ""
	for _, b := range req.BarberEmployees {

		validations := validateBarberEmployee(b)
		if validations != nil {
			return nil, inValidArgumentError(validations)
		}

		newUUID := uuid.New()
		uuidPrefix := newUUID.String()[:8]
		combinedNickName := b.NickName + uuidPrefix

		var hashedPasswordValid bool = hashedPassword != ""
		arg := db.CreateBarberEmployeeParams{
			Phone: b.GetPhone(),
			HashedPassword: sql.NullString{
				String: hashedPassword,
				Valid:  hashedPasswordValid,
			},
			NickName: combinedNickName,
			FullName: b.GetFullName(),
		}

		barberResponse, err := server.createSingleBarber(ctx, req, arg)
		if err != nil {
			return nil, status.Error(codes.Internal, "internal")
		}
		responses += barberResponse + " "
	}

	return &barber.CreateBarberEmployeesResponse{
		Message: responses,
	}, nil
}

func (server *Server) createSingleBarber(ctx context.Context, req *barber.CreateBarberEmployeesRequest,  arg db.CreateBarberEmployeeParams) (string, error) {
	var resBarber db.Barber
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		resBarber, err = server.Store.CreateBarberEmployee(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pgconn.PgError); ok {
				switch pqErr.ConstraintName {
				case "Barbers_pkey":
					return returnError(codes.AlreadyExists, "This id has already existed", err)
				case "Barbers_email_key":
					return returnError(codes.AlreadyExists, "This email has already existed", err)
				case "Barbers_phone_key":
					return returnError(codes.AlreadyExists, "This phone has already existed", err)
				case "Barbers_nick_name_key":
					return returnError(codes.AlreadyExists, "This nick name has already existed", err)
				}
			}
			return internalError(err)
		}

		barberShopID, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "barbershops don't exist")
		}
		argBarberRole := db.CreateBarberRolesParams{
			BarberID:     resBarber.ID,
			BarberShopID: barberShopID,
			RoleID:       int16(req.RoleId),
		}
		_, err = server.Store.CreateBarberRoles(ctx, argBarberRole)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}

		return nil
	})
	if err != nil {
		return "", internalError(err)
	}

	return resBarber.NickName, nil
}
func validateBarberEmployee(req *barber.BarberEmployeeRequest) (validations []*errdetails.BadRequest_FieldViolation) {

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}
