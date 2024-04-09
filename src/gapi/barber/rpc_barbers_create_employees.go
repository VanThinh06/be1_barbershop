package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"
	"strings"

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

	response, err := server.createMultiBarber(ctx, req, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &barber.CreateBarberEmployeesResponse{
		Message: response,
	}, nil
}

func (server *Server) createMultiBarber(ctx context.Context, req *barber.CreateBarberEmployeesRequest, hashedPassword string) (string, error) {
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		for _, b := range req.BarberEmployees {

			validations := validateBarberEmployee(b)
			if validations != nil {
				return inValidArgumentError(validations)
			}

			newUUID := uuid.New()
			uuidPrefix := newUUID.String()[:8]
			trimNickName := strings.ReplaceAll(b.NickName, " ", "")
			combinedNickName := strings.TrimSpace(trimNickName) + uuidPrefix
			if err = helpers.ValidateNickName(combinedNickName); err != nil {
				validations = append(validations, FieldValidation("nick_name", err))
				return inValidArgumentError(validations)
			}

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

			resBarber, err := server.Store.CreateBarberEmployee(ctx, arg)
			if err != nil {
				if pqErr, ok := err.(*pgconn.PgError); ok {
					strErr := ""
					switch pqErr.ConstraintName {
					case "Barbers_pkey":
						strErr = arg.FullName + " 's id has already existed"
						return returnError(codes.AlreadyExists, strErr, err)
					case "Barbers_email_key":
						strErr = arg.FullName + " 's email already exists"
						return returnError(codes.AlreadyExists, strErr, err)
					case "Barbers_phone_key":
						strErr = arg.FullName + " 's phone number already exists"
						return returnError(codes.AlreadyExists, strErr, err)
					case "Barbers_nick_name_key":
						strErr = arg.FullName + " 's nick name already exists"
						return returnError(codes.AlreadyExists, strErr, err)
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
				return err
			}
		}
		return err
	})

	return "", err
}
func validateBarberEmployee(req *barber.CreateBarberEmployee) (validations []*errdetails.BadRequest_FieldViolation) {

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}
