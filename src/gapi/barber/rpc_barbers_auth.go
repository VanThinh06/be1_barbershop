package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginBarber(ctx context.Context, req *barber.LoginBarberRequest) (*barber.LoginBarberResponse, error) {

	res, err := server.Store.GetBarbersEmail(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "you have not created an account yet")
		}
		return nil, status.Error(codes.InvalidArgument, "username or password is incorrect")
	}

	err = utilities.CheckPassword(req.Password, res.HashedPassword)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "username or password is incorrect")
	}

	barberPayload := token.BarberPayload{
		BarberID:       res.ID,
		BarberRole:     res.BarberRole,
		BarberRoleType: utilities.MapRoleToRoleType[res.BarberRole],
		Phone:          res.Phone,
		Email:          res.Email,
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
	session, err := server.Store.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
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
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp, nil
}
