package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
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
				Valid:        true,
			},
			EndTimeMonday: pgtype.Time{
				Microseconds: req.EndTimeMonday,
				Valid:        true,
			},
			StartTimeTuesday: pgtype.Time{
				Microseconds: req.StartTimeTuesday,
				Valid:        true,
			},
			EndTimeTuesday: pgtype.Time{
				Microseconds: req.EndTimeTuesday,
				Valid:        true,
			},
			StartTimeWednesday: pgtype.Time{
				Microseconds: req.StartTimeWednesday,
				Valid:        true,
			},
			EndTimeWednesday: pgtype.Time{
				Microseconds: req.EndTimeWednesday,
				Valid:        true,
			},
			StartTimeThursday: pgtype.Time{
				Microseconds: req.StartTimeThursday,
				Valid:        true,
			},
			EndTimeThursday: pgtype.Time{
				Microseconds: req.EndTimeThursday,
				Valid:        true,
			},
			StartTimeFriday: pgtype.Time{
				Microseconds: req.StartTimeFriday,
				Valid:        true,
			},
			EndTimeFriday: pgtype.Time{
				Microseconds: req.EndTimeFriday,
				Valid:        true,
			},
			StartTimeSaturday: pgtype.Time{
				Microseconds: req.StartTimeSaturday,
				Valid:        true,
			},
			EndTimeSaturday: pgtype.Time{
				Microseconds: req.EndTimeSaturday,
				Valid:        true,
			},
			StartTimeSunday: pgtype.Time{
				Microseconds: req.StartTimeSunday,
				Valid:        true,
			},
			EndTimeSunday: pgtype.Time{
				Microseconds: req.EndTimeSunday,
				Valid:        true,
			},
			IntervalScheduler: int16(req.IntervalScheduler),
			Longitude:         float64(req.Longitude),
			Latitude:          float64(req.Latitude),
		}
		resBarberShop, err := server.Store.CreateBarberShop(ctx, arg)
		if err != nil {
			return internalError(err)
		}

		argBarberRole := db.CreateBarberRoleParams{
			BarberID:     payload.Barber.BarberID,
			BarberShopID: resBarberShop.ID,
			RoleID:       int16(utilities.Admin),
		}
		_, err = server.Store.CreateBarberRole(ctx, argBarberRole)
		if err != nil {
			return status.Errorf(codes.Internal, err.Error())
		}
		return nil
	})

	return err
}
