package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateBarberShop(ctx context.Context, req *barber.UpdateBarberShopRequest) (*barber.UpdateBarberShopResponse, error) {

	payload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	if payload.Barber.Role != int32(utils.Manager) {
		return nil, status.Errorf(codes.PermissionDenied, "unauthenticated")
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
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Longitude: sql.NullFloat64{
			Float64: req.GetLongtude().GetValue(),
			Valid:   req.Longtude != nil,
		},
		Latitude: sql.NullFloat64{
			Float64: req.GetLatitude().GetValue(),
			Valid:   req.Latitude != nil,
		},
		Address: sql.NullString{
			String: req.GetAddress(),
			Valid:  req.Address != nil,
		},
		Image: sql.NullString{
			String: req.GetImage(),
			Valid:  req.Image != nil,
		},

		Facility: sql.NullInt32{
			Int32: req.GetFacility(),
			Valid: req.Facility != nil,
		},
		ShopID:    uuid.MustParse(req.GetShopId()),
		StartTime: startTime,
		EndTime:   endTime,
		BreakTime: breakTime,
		Status: sql.NullInt32{
			Int32: req.GetStatus(),
			Valid: req.Status != nil,
		},
		IntervalScheduler: sql.NullInt32{
			Int32: req.GetIntervalSheduler(),
			Valid: req.IntervalSheduler != nil,
		},
	}
	barberShop, err := server.Store.UpdateBarberShop(ctx, arg)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.UpdateBarberShopResponse{
		BarberShop: convertBarberShop(barberShop),
	}
	return rsp, nil
}
