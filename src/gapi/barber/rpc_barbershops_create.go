package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShops(ctx context.Context, req *barber.CreateBarberShopsRequest) (*barber.CreateBarberShopsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	var barberShopChainId = uuid.NullUUID{
		UUID:  uuid.MustParse(req.GetBarberShopChainId()),
		Valid: req.BarberShopChainId != nil,
	}
	arg := db.CreateBarberShopParams{
		BarberShopChainID: barberShopChainId,
		Name:              req.Name,
		ProvinceID:        int16(req.ProvinceId),
		DistrictID:        int16(req.DistrictId),
		WardID:            int16(req.WardId),
		SpecificLocation:  req.SpecificLocation,
		Phone:             req.Phone,
		Email:             req.Email,
		WebsiteUrl: sql.NullString{
			String: req.GetWebsiteUrl(),
			Valid:  req.WebsiteUrl != nil,
		},
		AvatarUrl:           req.AvatarUrl,
		CoverPhotoUrl:       req.CoverPhotoUrl,
		FacadePhotoUrl:      req.FacadePhotoUrl,
		RepresentativeName:  req.RepresentativeName,
		CitizenID:           req.CitizenId,
		RepresentativePhone: req.RepresentativePhone,
		RepresentativeEmail: req.RepresentativeEmail,
		RepresentativePhoneOther: sql.NullString{
			String: req.GetRepresentativePhoneOther(),
			Valid:  req.RepresentativePhoneOther != nil,
		},
		StartTimeMonday: pgtype.Time{
			Microseconds: req.StartTimeMonday,
		},
		EndTimeMonday: pgtype.Time{
			Microseconds: req.EndTimeMonday,
		},
		StartTimeTuesday: pgtype.Time{
			Microseconds: req.StartTimeTuesday,
		},
		EndTimeTuesday: pgtype.Time{
			Microseconds: req.EndTimeTuesday,
		},
		StartTimeWednesday: pgtype.Time{
			Microseconds: req.StartTimeWednesday,
		},
		EndTimeWednesday: pgtype.Time{
			Microseconds: req.EndTimeWednesday,
		},
		StartTimeThursday: pgtype.Time{
			Microseconds: req.StartTimeThursday,
		},
		EndTimeThursday: pgtype.Time{
			Microseconds: req.EndTimeThursday,
		},
		StartTimeFriday: pgtype.Time{
			Microseconds: req.StartTimeFriday,
		},
		EndTimeFriday: pgtype.Time{
			Microseconds: req.EndTimeFriday,
		},
		StartTimeSaturday: pgtype.Time{
			Microseconds: req.StartTimeSaturday,
		},
		EndTimeSaturday: pgtype.Time{
			Microseconds: req.EndTimeSaturday,
		},
		StartTimeSunday: pgtype.Time{
			Microseconds: req.StartTimeSunday,
		},
		EndTimeSunday: pgtype.Time{
			Microseconds: req.EndTimeSunday,
		},
		IntervalScheduler: int16(req.IntervalScheduler),
		Longitude: req.Longitude.Value,
		Latitude:  req.Latitude.Value,
	}

	barberShop, err := server.Store.CreateBarberShop(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	argBarberRole := db.CreateBarberRolesParams{
		BarberID: payload.Barber.BarberID,
		BarberShopID: barberShop.ID,
		RoleID: int16(utilities.Manager),
	}

	rsp := &barber.CreateBarberShopsResponse{
		BarberShop: convertBarberShops(barberShop),
	}
	return rsp, nil
}
