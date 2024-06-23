package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/utilities"
	"barbershop/src/utils"
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// barber employee create single
func (server *Server) CreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest) (*barber.CreateBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, utils.UnauthenticatedError(err)
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, utils.NotFoundError(err, "barber shop don't exist")
	}

	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ManageEmployee),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, utils.NoPermissionError(err)
	}

	validations := utils.ValidateBarberEmployee(req.BarberEmployee)
	if validations != nil {
		return nil, utils.InvalidArgumentError(validations)
	}

	defaultPassword, err := server.Store.GetDefaultPasswordEmployee(ctx, barberShopID)
	if err != nil {
		return nil, utils.InternalError(err)
	}
	var hashedPassword string
	if defaultPassword.Valid {
		hashedPassword, err = helpers.HashPassword(defaultPassword.String)
		if err != nil {
			return nil, utils.InternalError(err)
		}
	}

	nickName := utils.GenerateNickName(req.BarberEmployee.FullName)

	var hashedPasswordValid bool = hashedPassword != ""
	arg := db.CreateBarberEmployeeParams{
		Phone: req.BarberEmployee.GetPhone(),
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  hashedPasswordValid,
		},
		NickName: strings.ToLower(nickName),
		FullName: req.BarberEmployee.GetFullName(),
	}

	errTx := make(chan error)

	go func() {
		err := server.txCreateBarberEmployee(ctx, req, arg)
		errTx <- err
	}()

	err = <-errTx

	if err != nil {
		return nil, err
	}

	rsp := &barber.CreateBarberEmployeeResponse{
		Message: "Barber added successfully.",
	}
	return rsp, nil
}

func (server *Server) txCreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest, arg db.CreateBarberEmployeeParams) error {

	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {

		resBarber, err := server.Store.CreateBarberEmployee(ctx, arg)
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
		argBarberRole := db.CreateBarberRoleParams{
			BarberID:     resBarber.ID,
			BarberShopID: barberShopID,
			RoleID:       int16(req.RoleId),
		}
		_, err = server.Store.CreateBarberRole(ctx, argBarberRole)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

// barber employee create multi
func (server *Server) CreateMultiBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeesRequest) (*barber.CreateBarberEmployeesResponse, error) {
	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ManageEmployee),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, noPermissionError(err)
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

	err = server.txCreateMultiBarberEmployee(ctx, req, hashedPassword)
	if err != nil {
		return nil, err
	}

	return &barber.CreateBarberEmployeesResponse{
		Message: "Barber added successfully.",
	}, nil
}

func (server *Server) txCreateMultiBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeesRequest, hashedPassword string) error {
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		for _, b := range req.BarberEmployees {

			validations := utils.ValidateBarberEmployee(b)
			if validations != nil {
				return inValidArgumentError(validations)
			}

			nickName := utils.GenerateNickName(b.FullName)
			var hashedPasswordValid bool = hashedPassword != ""
			arg := db.CreateBarberEmployeeParams{
				Phone: b.GetPhone(),
				HashedPassword: sql.NullString{
					String: hashedPassword,
					Valid:  hashedPasswordValid,
				},
				NickName: strings.ToLower(nickName),
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
			argBarberRole := db.CreateBarberRoleParams{
				BarberID:     resBarber.ID,
				BarberShopID: barberShopID,
				RoleID:       int16(req.RoleId),
			}
			_, err = server.Store.CreateBarberRole(ctx, argBarberRole)
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}

// barber employee get list
func (server *Server) GetBarberEmployees(ctx context.Context, req *barber.GetBarberEmployeeRequest) (*barber.GetBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ViewEmployeeInformation),
		BarberID:     payload.Barber.BarberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return nil, noPermissionError(err)
	}

	arg := db.ListEmployeesParams{
		BarberShopID: barberShopID,
		Limit:        req.Limit,
		Offset:       req.Limit * (req.Page - 1),
	}
	barberEmployees, err := server.Store.ListEmployees(ctx, arg)
	if err != nil {
		return nil, internalError(err)
	}

	var totalEmployees int16 = 0
	if len(barberEmployees) > 0 {
		totalEmployees = int16(barberEmployees[0].TotalEmployees)
	}
	rsp := &barber.GetBarberEmployeeResponse{
		BarberDetails:  convertBarberEmployees(barberEmployees),
		TotalEmployees: int32(totalEmployees),
	}
	return rsp, nil
}

// barber employee delete
func (server *Server) DeleteBarberEmployee(ctx context.Context, req *barber.DeleteBarberEmployeeRequest) (*barber.DeleteBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	if payload.Barber.BarberID.String() != req.Id {
		argCheckPermission := db.CheckBarberRolePermissionParams{
			ID:           int16(utilities.ManageEmployee),
			BarberID:     payload.Barber.BarberID,
			BarberShopID: barberShopId,
		}
		isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
		if !isPermission {
			return nil, noPermissionError(err)
		}
	}

	barberId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "employee don't exist")
	}
	argBarber := db.GetBarberEmployeeParams{
		ID:           barberId,
		BarberShopID: barberShopId,
	}
	resBarber, err := server.Store.GetBarberEmployee(ctx, argBarber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "the employee does not belong to this barber shop", err)
		}
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to delete employee")
	}
	if !resBarber.ID_2.Valid {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to delete employee")
	}

	err = server.Store.DeleteBarberRole(ctx, resBarber.ID_2.UUID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberEmployeeResponse{
		Message: "Employee " + resBarber.NickName + " has been successfully deleted",
	}

	if !resBarber.HashedPassword.Valid {
		err = server.Store.DeleteBarber(ctx, barberId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal")
		}
	}
	return rsp, nil
}

// barber employee update
func (server *Server) UpdateBarber(ctx context.Context, req *barber.UpdateBarberRequest) (*barber.UpdateBarberResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	validations := utils.ValidateBarberUpdate(req)
	if validations != nil {
		return nil, inValidArgumentError(validations)
	}

	if authPayload.Barber.BarberID.String() != req.Barber.Id {
		barberShopId, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "barbershop don't exist")
		}

		argCheckPermission := db.CheckBarberRolePermissionParams{
			ID:           int16(utilities.ManageEmployee),
			BarberID:     authPayload.Barber.BarberID,
			BarberShopID: barberShopId,
		}
		isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
		if !isPermission {
			return nil, noPermissionError(err)
		}

		barberId, err := uuid.Parse(req.Barber.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "barber don't exist")
		}
		argBarberRole := db.GetBarberRoleParams{
			BarberID:     barberId,
			BarberShopID: barberShopId,
		}
		_, err = server.Store.GetBarberRole(ctx, argBarberRole)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
		}
	}

	arg := db.UpdateBarberParams{
		ID: uuid.MustParse(req.Barber.GetId()),

		NickName: sql.NullString{
			String: strings.ToLower(req.Barber.GetNickname()),
			Valid:  req.Barber.Nickname != nil,
		},

		FullName: sql.NullString{
			String: req.Barber.GetFullName(),
			Valid:  req.Barber.FullName != nil,
		},

		Phone: sql.NullString{
			String: req.Barber.GetPhone(),
			Valid:  req.Barber.Phone != nil,
		},
		Haircut: pgtype.Bool{
			Bool:  req.Barber.GetHaircut(),
			Valid: req.Barber.Haircut != nil,
		},
		GenderID: pgtype.Int2{
			Int16: int16(req.Barber.GetGenderId()),
			Valid: req.Barber.GenderId != nil,
		},
		WorkStatus: pgtype.Bool{
			Bool:  req.Barber.GetWorkStatus(),
			Valid: req.Barber.WorkStatus != nil,
		},

		Email: sql.NullString{
			String: strings.ToLower(req.Barber.GetEmail()),
			Valid:  req.Barber.Email != nil,
		},
		AvatarUrl: sql.NullString{
			String: req.Barber.GetAvatarUrl(),
			Valid:  req.Barber.AvatarUrl != nil,
		},
	}

	_, err = server.Store.UpdateBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "barber not found", err)
		}
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "Barbers_pkey":
				return nil, returnError(codes.AlreadyExists, "This id has already existed", err)
			case "Barbers_email_key":
				return nil, returnError(codes.AlreadyExists, "This email has already existed", err)
			case "Barbers_phone_key":
				return nil, returnError(codes.AlreadyExists, "This phone has already existed", err)
			case "Barbers_nick_name_key":
				return nil, returnError(codes.AlreadyExists, "This nick name has already existed", err)
			}
		}
		return nil, internalError(err)
	}

	rsp := &barber.UpdateBarberResponse{
		Message: "Your information has been successfully updated.",
	}
	return rsp, nil
}

// barber get permissions from barber shop
func (server *Server) GetPermissionFromBarberShop(ctx context.Context, req *barber.GetPermissionFromBarberShopRequest) (*barber.GetPermissionFromBarberShopResponse, error) {
	authPayload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershop don't exist")
	}

	barberId, err := uuid.Parse(req.GetBarberId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barber don't exist")
	}
	if authPayload.Barber.BarberID.String() != req.GetBarberId() {
		return nil, status.Errorf(codes.InvalidArgument, "barber don't exist")
	}

	argPermission := db.GetPermissionFromBarberShopParams{
		ID:           barberId,
		BarberShopID: barberShopId,
	}

	permission, err := server.Store.GetPermissionFromBarberShop(ctx, argPermission)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.GetPermissionFromBarberShopResponse{
		Permission: convertListPermission(permission),
	}
	return rsp, nil
}
