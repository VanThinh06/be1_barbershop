package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarberShop(ctx context.Context, req *barber.UpdateBarberShopsRequest) (*barber.UpdateBarberShopsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	if payload.Barber.BarberRole != int32(utilities.Admin) {
		return nil, unauthenticatedError(err)
	}

	startTime := pgtype.Time{}
	endTime := pgtype.Time{}
	breakTime := pgtype.Time{}
	err = startTime.Set(time.Date(0, 1, 1, int(req.GetStartTime().Hours), int(req.GetStartTime().Minutes), 0, 0, time.UTC))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parameters")
	}
	err = endTime.Set(time.Date(0, 1, 1, int(req.GetEndTime().Hours), int(req.GetEndTime().Minutes), 0, 0, time.UTC))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parameters")
	}
	err = breakTime.Set(time.Date(0, 1, 1, int(req.GetBreakTime().Hours), int(req.GetBreakTime().Minutes), 0, 0, time.UTC))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid parameters")
	}

	arg := db.UpdateBarberShopParams{
		ID: uuid.MustParse(req.Id),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		// Longitude: sql.NullFloat64{
		// 	Float64: req.GetLongtude().GetValue(),
		// 	Valid:   req.Longtude != nil,
		// },
		// Latitude: sql.NullFloat64{
		// 	Float64: req.GetLatitude().GetValue(),
		// 	Valid:   req.Latitude != nil,
		// },
		// IsMainBranch: sql.NullBool{
		// 	Bool:  req.GetIsMainBranch(),
		// 	Valid: req.IsMainBranch != nil,
		// },
	}
	barberShop, err := server.Store.UpdateBarberShop(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershop don't exist")
		}
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.UpdateBarberShopsResponse{
		BarberShop: convertBarberShops(barberShop),
	}
	return rsp, nil
}
