package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"barbershop/src/utils"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateBarberShop(ctx context.Context, req *barber.CreateBarberShopRequest) (*barber.CreateBarberShopResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	errTx := make(chan error)
	go func() {
		err := server.txCreateBarberShop(ctx, req, payload)
		errTx <- err
	}()
	err = <-errTx
	if err != nil {
		return nil, internalError(err)
	}

	var message = "Congratulations! You have successfully registered for " + req.Name + "."
	rsp := &barber.CreateBarberShopResponse{
		Message: message,
	}

	return rsp, nil
}

func (server *Server) txCreateBarberShop(ctx context.Context, req *barber.CreateBarberShopRequest, payload *token.BarberPayload) error {
	err := server.Store.ExecTx(ctx, func(q *db.Queries) error {

		var barberShopChainId uuid.NullUUID
		barberShopChainUUID, err := uuid.Parse(
			req.GetBarberShopChainId(),
		)
		if err != nil {
			barberShopChainId = uuid.NullUUID{
				UUID:  barberShopChainUUID,
				Valid: req.BarberShopChainId != nil,
			}
		}

		arg := db.InsertBarberShopParams{
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
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeMonday),
				Valid:        true,
			},
			EndTimeMonday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeMonday),
				Valid:        true,
			},
			StartTimeTuesday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeTuesday),
				Valid:        true,
			},
			EndTimeTuesday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeTuesday),
				Valid:        true,
			},
			StartTimeWednesday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeWednesday),
				Valid:        true,
			},
			EndTimeWednesday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeWednesday),
				Valid:        true,
			},
			StartTimeThursday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeThursday),
				Valid:        true,
			},
			EndTimeThursday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeThursday),
				Valid:        true,
			},
			StartTimeFriday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeFriday),
				Valid:        true,
			},
			EndTimeFriday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeFriday),
				Valid:        true,
			},
			StartTimeSaturday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeSaturday),
				Valid:        true,
			},
			EndTimeSaturday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeSaturday),
				Valid:        true,
			},
			StartTimeSunday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.StartTimeSunday),
				Valid:        true,
			},
			EndTimeSunday: pgtype.Time{
				Microseconds: utils.ConvertMinutesToMicroseconds(req.EndTimeSunday),
				Valid:        true,
			},
			IntervalScheduler: int16(req.IntervalScheduler),
			Longitude:         float64(req.Longitude),
			Latitude:          float64(req.Latitude),
		}
		resBarberShop, err := server.Store.InsertBarberShop(ctx, arg)
		if err != nil {
			return internalError(err)
		}

		argBarberRole := db.InsertBarberRoleParams{
			BarberID:     payload.Barber.BarberID,
			BarberShopID: resBarberShop.ID,
			RoleID:       int16(utilities.Admin),
		}
		_, err = server.Store.InsertBarberRole(ctx, argBarberRole)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		return nil
	})

	return err
}


func (server *Server) GetBarberShop(ctx context.Context, req *barber.GetBarberShopRequest) (*barber.GetBarberShopResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	barberShopId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetBarberShopById(ctx, barberShopId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershops don't exist")
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopResponse{
		BarberShop: convertGetBarberShop(res),
	}
	return rsp, nil
}

func (server *Server) ListBarberShops(ctx context.Context, req *barber.ListBarberShopsRequest) (*barber.ListBarberShopsResponse, error) {

	payload, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	res, err := server.Store.ListBarberShopsByBarberId(ctx, payload.Barber.BarberID)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListBarberShopsResponse{
		BarberShops: convertListBarberShops(res),
	}
	return rsp, nil
}


func (server *Server) UpdateBarberShop(ctx context.Context, req *barber.UpdateBarberShopRequest) (*barber.UpdateBarberShopResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	arg := db.UpdateBarberShopByIdParams{
		ID: uuid.MustParse(req.Id),
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
	}
	barberShop, err := server.Store.UpdateBarberShopById(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "barbershop don't exist")
		}
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.UpdateBarberShopResponse{
		BarberShop: convertBarberShop(barberShop),
	}
	return rsp, nil
}


func (server *Server) DeleteBarberShops(ctx context.Context, req *barber.DeleteBarberShopsRequest) (*barber.DeleteBarberShopsResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}


	var id = uuid.MustParse(req.Id)
	err = server.Store.DeleteBarberShopById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.DeleteBarberShopsResponse{
		Status: "success",
	}
	return rsp, nil
}
