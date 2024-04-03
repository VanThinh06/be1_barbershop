package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginBarber(ctx context.Context, req *barber.LoginBarberRequest) (*barber.LoginBarberResponse, error) {
	contact := db.GetUserBarberParams{
		TypeUsername: "phone",
		Email: sql.NullString{
			String: req.GetUsername(),
			Valid:  true,
		},
	}

	err := helpers.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.TypeUsername = "email"
		contact.Email = sql.NullString{
			String: req.GetUsername(),
			Valid:  true,
		}
	}

	res, err := server.Store.GetUserBarber(ctx, contact)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "you have not created an account yet")
		}
		return nil, status.Error(codes.InvalidArgument, "username or password is incorrect")
	}

	err = helpers.CheckPassword(req.Password, res.HashedPassword)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "username or password is incorrect")
	}

	barberPayload := token.Barber{
		BarberID:       res.ID,
		BarberRole:     int32(res.BarberRole.Int16),
		BarberRoleType: utilities.MapRoleToRoleType[int32(res.BarberRole.Int16)],
		Phone:          res.Phone,
		Email:          res.Email.String,
		FcmDevice:      req.FcmDevice,
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	mtdt := server.extractMetadata(ctx)
	_, err = server.Store.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
		BarberID:     refreshPayload.Barber.BarberID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		FcmDevice:    req.FcmDevice,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.LoginBarberResponse{
		Barber:                convertBarbersEmail(res),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}
	return rsp, nil
}
