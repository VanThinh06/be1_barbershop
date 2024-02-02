package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func (server *Server) CreateBarber(ctx context.Context, req *barber.CreateBarbersRequest) (*barber.CreateBarbersResponse, error) {
	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if payload.Barber.BarberRoleType != string(utilities.Administrator) || payload.Barber.BarberRole != int32(utilities.Manager) {
		return nil, noPermissionError(err)
	}

	if payload.Barber.BarberRole == int32(utilities.Manager) {
		if req.RoleId != int32(utilities.Barber) {
			return nil, noPermissionError(err)
		}
	}

	if payload.Barber.BarberRole == int32(utilities.Admin) {
		if req.RoleId == int32(utilities.SuperAdmin) {
			return nil, noPermissionError(err)
		}
	}

	var barberShopId uuid.NullUUID
	if payload.Barber.BarberRole != int32(utilities.SuperAdmin) {
		barberShopId = uuid.NullUUID{
			UUID:  uuid.MustParse(req.GetBarberShopId()),
			Valid: true,
		}
	}
	arg := db.GetBarberRolesParams{
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopId,
	}
	_, err = server.Store.GetBarberRoles(ctx, arg)
	if err != nil {
		return nil, noPermissionError(err)
	}

	validations := validateCreateBarber(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	errTx := make(chan error)
	resultBarber := make(chan db.Barber)

	go func() {
		result, err := server.TxCreateBarber(ctx, req)
		errTx <- err
		resultBarber <- result
	}()

	err = <-errTx
	res := <-resultBarber

	if err != nil {
		return nil, internalError(err)
	}

	rsp := &barber.CreateBarbersResponse{
		Barber: convertCreateBarbers(res),
	}
	return rsp, nil
}

func (server *Server) TxCreateBarber(ctx context.Context, req *barber.CreateBarbersRequest) (db.Barber, error) {
	var resBarber db.Barber

	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {

		hashedPassword, err := utilities.HashPassword(req.GetPassword())
		if err != nil {
			return internalError(err)
		}

		arg := db.CreateBarbersParams{
			NickName:       req.GetNickName(),
			HashedPassword: hashedPassword,
			Phone:          req.GetPhone(),
			FullName:       req.GetEmail(),
			GenderID:       req.GetGenderId(),
			Email:          req.GetEmail(),
		}
		resBarber, err = server.Store.CreateBarbers(ctx, arg)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				switch pqErr.Code.Name() {
				case "unique_violation":
					if pqErr.Constraint == "Barbers_pkey" {
						return returnError(codes.AlreadyExists, "This account has already existed", err)
					}
					if pqErr.Constraint == "Barbers_phone_key" {
						return returnError(codes.AlreadyExists, "This phone has already existed", err)
					}
					if pqErr.Constraint == "Barbers_email_key" {
						return returnError(codes.AlreadyExists, "This email has already existed", err)
					}
				}
			}
			return internalError(err)
		}

		argBarberManager := db.CreateBarberManagersParams{
			BarberID:  resBarber.ID,
			ManagerID: uuid.MustParse(req.GetBarberManagerId()),
		}

		_, err = q.CreateBarberManagers(ctx, argBarberManager)
		if err != nil {
			return internalError(err)
		}

		barberShopId, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return notFoundError(err, "barberShop")
		}
		argBarberRole := db.CreateBarberRolesParams{
			RoleID:   req.GetRoleId(),
			BarberID: resBarber.ID,
			BarberShopID: uuid.NullUUID{
				UUID:  barberShopId,
				Valid: true,
			},
		}
		_, err = q.CreateBarberRoles(ctx, argBarberRole)
		if err != nil {
			return internalError(err)
		}
		return nil
	})
	return resBarber, err
}

func validateCreateBarber(req *barber.CreateBarbersRequest) (validations []*errdetails.BadRequest_FieldViolation) {
	if err := helpers.ValidateEmail(req.Email); err != nil {
		validations = append(validations, FieldValidation("email", err))
	}

	if err := helpers.ValidatePhoneNumber(req.Phone); err != nil {
		validations = append(validations, FieldValidation("phone", err))
	}

	if err := helpers.ValidatePassword(req.Password); err != nil {
		validations = append(validations, FieldValidation("password", err))
	}

	if err := helpers.ValidateNickname(req.NickName); err != nil {
		validations = append(validations, FieldValidation("nickname", err))
	}

	if err := helpers.ValidateFullName(req.FullName); err != nil {
		validations = append(validations, FieldValidation("full_name", err))
	}

	return validations
}
