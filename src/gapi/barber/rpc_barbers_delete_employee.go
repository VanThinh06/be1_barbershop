package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteBarberEmployee(ctx context.Context, req *barber.DeleteBarberEmployeeRequest) (*barber.DeleteBarberEmployeeResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	if payload.Barber.BarberID.String() != req.Id {
		barberShopId, err := uuid.Parse(req.BarberShopId)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
		}
		argBarberRoleAdmin := db.GetBarberRolesParams{
			BarberID:     payload.Barber.BarberID,
			BarberShopID: barberShopId,
		}
		roleAdmin, err := server.Store.GetBarberRoles(ctx, argBarberRoleAdmin)
		if err != nil {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update employee")
		}
		if roleAdmin.RoleID != int16(utilities.Admin) {
			return nil, status.Errorf(codes.PermissionDenied, "failed to no permission to update employee")
		}
	}

	barberId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "employee don't exist")
	}
	barberShopId, err := uuid.Parse(req.BarberShopId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}
	argBarber := db.GetBarberEmployeeParams{
		ID:           barberId,
		BarberShopID: barberShopId,
	}
	resBarber, err := server.Store.GetBarberEmployee(ctx, argBarber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, returnError(codes.NotFound, "the employee does not belong to this barber shop ", err)
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
		Message: "Nhân viên " + resBarber.NickName + " đã được xóa thành công",
	}

	if !resBarber.HashedPassword.Valid {
		err = server.Store.DeleteBarber(ctx, barberId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal")
		}
	}
	return rsp, nil
}
