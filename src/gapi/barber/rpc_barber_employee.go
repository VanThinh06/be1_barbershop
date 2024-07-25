package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
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

// barber employee create multi
func (server *Server) CreateBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest) (*barber.CreateBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopID, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	err = server.checkPermissionManageEmployee(ctx, barberShopID, payload.Barber.BarberID)
	if err != nil {
		return nil, noPermissionError(err)
	}

	valPassword := utils.ValidatePassword(req.DefaultPassword)
	if valPassword != nil {
		return nil, err
	}
	hasDefaultPassword, err := server.tokenMaker.CreateAESString(req.GetDefaultPassword())
	if err != nil {
		return nil, internalError(err)
	}

	err = server.txCreateMultiBarberEmployee(ctx, req, hasDefaultPassword)
	if err != nil {
		return nil, err
	}

	return &barber.CreateBarberEmployeeResponse{
		Message: "Barber added successfully.",
	}, nil
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

	arg := db.ListBarberEmployeesParams{
		BarberShopID: barberShopID,
		Limit:        req.Limit,
		Offset:       req.Limit * (req.Page - 1),
	}
	barberEmployees, err := server.Store.ListBarberEmployees(ctx, arg)
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
	argBarber := db.GetBarberWithOptionalRoleParams{
		ID:           barberId,
		BarberShopID: barberShopId,
	}
	resBarber, err := server.Store.GetBarberWithOptionalRole(ctx, argBarber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "the employee does not belong to this barber shop", err)
		}
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to delete employee")
	}
	if !resBarber.ID_2.Valid {
		return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to delete employee")
	}

	err = server.Store.DeleteBarberRoleById(ctx, resBarber.ID_2.UUID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberEmployeeResponse{
		Message: "Employee " + resBarber.NickName + " has been successfully deleted",
	}

	if !resBarber.HashedPassword.Valid {
		err = server.Store.DeleteBarberById(ctx, barberId)
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
		argBarberRole := db.SelectBarberRoleByIdsParams{
			BarberID:     barberId,
			BarberShopID: barberShopId,
		}
		_, err = server.Store.SelectBarberRoleByIds(ctx, argBarberRole)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update barber")
		}
	}

	arg := db.UpdateBarberDetailsParams{
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

	_, err = server.Store.UpdateBarberDetails(ctx, arg)
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

	argPermission := db.GetPermissionsByBarberShopParams{
		ID:           authPayload.Barber.BarberID,
		BarberShopID: barberShopId,
	}

	permission, err := server.Store.GetPermissionsByBarberShop(ctx, argPermission)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.GetPermissionFromBarberShopResponse{
		Permission: convertListPermission(permission),
	}
	return rsp, nil
}

func (server *Server) checkPermissionManageEmployee(ctx context.Context, barberShopID uuid.UUID, barberID uuid.UUID) error {
	argCheckPermission := db.CheckBarberRolePermissionParams{
		ID:           int16(utilities.ManageEmployee),
		BarberID:     barberID,
		BarberShopID: barberShopID,
	}
	isPermission, err := server.Store.CheckBarberRolePermission(ctx, argCheckPermission)
	if !isPermission {
		return noPermissionError(err)
	}
	return nil
}

func (server *Server) txCreateMultiBarberEmployee(ctx context.Context, req *barber.CreateBarberEmployeeRequest, hasDefaultPassword string) error {
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {

		for _, b := range req.BarberEmployees {

			validations := utils.ValidateBarberEmployee(b)
			if validations != nil {
				return inValidArgumentError(validations)
			}

			nickName := utils.GenerateNickName(b.FullName)
			arg := db.InsertBarberParams{
				Phone: b.GetPhone(),
				DefaultPasswordEncrypted: pgtype.Text{
					String: hasDefaultPassword,
					Valid:  true,
				},
				NickName: strings.ToLower(nickName),
				FullName: b.GetFullName(),
			}

			resBarber, err := q.InsertBarber(ctx, arg)
			if err != nil {
				if pqErr, ok := err.(*pgconn.PgError); ok {
					var strErr string
					switch pqErr.ConstraintName {
					case "Barbers_pkey":
						strErr = arg.FullName + "'s ID already exists"
					case "Barbers_phone_key":
						strErr = arg.FullName + "'s phone number already exists"
					case "Barbers_nick_name_key":
						strErr = arg.FullName + "'s nick name already exists"
					default:
						strErr = "Unexpected error occurred"
					}
					return returnError(codes.AlreadyExists, strErr, err)
				}
				return utils.InternalError(err)
			}
			barberShopID := uuid.MustParse(req.BarberShopId)

			argBarberRole := db.InsertBarberRoleParams{
				BarberID:     resBarber.ID,
				BarberShopID: barberShopID,
				RoleID:       int16(req.RoleId),
			}
			_, err = q.InsertBarberRole(ctx, argBarberRole)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
