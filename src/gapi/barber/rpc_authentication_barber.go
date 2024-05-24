package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/token"
	"barbershop/src/utils"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// authentication barber register
func (server *Server) CreateBarber(ctx context.Context, req *barber.CreateBarberRequest) (*barber.CreateBarberResponse, error) {

	validations := utils.ValidateBarber(req)
	if validations != nil {
		return nil, utils.InvalidArgumentError(validations)
	}

	hashedPassword, err := helpers.HashPassword(req.GetPassword())
	if err != nil {
		return nil, utils.InternalError(err)
	}

	arg := db.CreateBarberParams{
		NickName: strings.ToLower(req.GetNickName()),
		FullName: req.GetFullName(),
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		Phone: req.GetPhone(),
	}
	res, err := server.Store.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "Barbers_pkey":
				return nil, utils.AlreadyExistsError(err, "This nick name has already existed")
			case "Barbers_phone_key":
				return nil, utils.AlreadyExistsError(err, "This nick name has already existed")
			case "Barbers_nick_name_key":
				return nil, utils.AlreadyExistsError(err, "This nick name has already existed")
			}
		}
		return nil, utils.InternalError(err)
	}

	rsp := &barber.CreateBarberResponse{
		Barber: &barber.Barbers{
			Id:        res.ID.String(),
			GenderId:  int32(res.GenderID.Int16),
			Email:     res.Email.String,
			Phone:     res.Phone,
			NickName:  res.NickName,
			FullName:  res.FullName,
			Haircut:   res.Haircut,
			AvatarUrl: res.AvatarUrl.String,
		},
	}
	return rsp, nil
}

// authenticate barber detail
func (server *Server) GetBarber(ctx context.Context, req *barber.GetBarbersRequest) (*barber.GetBarbersResponse, error) {
	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, utils.UnauthenticatedError(err)
	}
	idBarber, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, utils.NotFoundError(err, "barber don't exist")
	}
	idBarberShop, err := uuid.Parse(req.GetBarberShopId())
	if err != nil {
		return nil, utils.NotFoundError(err, "barber don't exist")
	}

	arg := db.GetBarberParams{
		ID:           idBarber,
		BarberShopID: idBarberShop,
	}
	res, err := server.Store.GetBarber(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NotFoundError(err, "barber don't exist")
		}
		return nil, utils.InternalError(err)
	}

	rsp := &barber.GetBarbersResponse{
		BarberDetail: convertBarberEmployee(res),
	}
	return rsp, nil
}

// authenticate barber login
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
			return nil, utils.NotFoundError(err, "you have not created an account yet")
		}
		return nil, utils.InvalidError(err, "username or password is incorrect")
	}

	if !res.HashedPassword.Valid {
		return nil, utils.InvalidError(err, "username or password is incorrect")
	}
	err = helpers.CheckPassword(req.Password, res.HashedPassword.String)
	if err != nil {
		return nil, utils.InvalidError(err, "username or password is incorrect")
	}

	barberPayload := token.Barber{
		BarberID:  res.ID,
		Phone:     res.Phone,
		Email:     res.Email.String,
		FcmDevice: req.FcmDevice,
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barberPayload,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, utils.InternalError(err)
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
		return nil, utils.InternalError(err)
	}

	rsp := &barber.LoginBarberResponse{
		Barber: &barber.AuthenticationBarber{
			Id:                    res.ID.String(),
			GenderId:              int32(res.GenderID.Int16),
			Email:                 res.Email.String,
			Phone:                 res.Phone,
			NickName:              res.NickName,
			FullName:              res.FullName,
			AvatarUrl:             res.AvatarUrl.String,
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
			RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
			BarberRoleId:          int32(res.RoleID.Int16),
			BarberShopId:          res.BarberShopID.UUID.String(),
		},
	}
	return rsp, nil
}

// authenticate barber refresh token
func (server *Server) RefreshTokenBarber(ctx context.Context, req *barber.RefreshTokenBarberRequest) (*barber.RefreshTokenBarberResponse, error) {
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return nil, utils.UnauthenticatedError(err)
	}

	session, err := server.Store.GetSessionBarber(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.NotFoundError(err, "not found")
		}

		return nil, utils.UnauthenticatedError(err)
	}

	if session.IsBlocked {
		_ = fmt.Errorf("incorrect block user")
		return nil, utils.UnauthenticatedError(err)

	}

	if session.BarberID != payload.Barber.BarberID {
		_ = fmt.Errorf("incorrect session user")

		return nil, unauthenticatedError(err)

	}

	if session.RefreshToken != req.RefreshToken {
		_ = fmt.Errorf("incorrect session user")

		return nil, utils.UnauthenticatedError(err)

	}

	if time.Now().After(session.ExpiresAt) {
		_ = fmt.Errorf("expired session")
		return nil, utils.UnauthenticatedError(err)

	}

	if session.ClientIp != server.extractMetadata(ctx).ClientIP {
		err := fmt.Errorf("incorrect clientIP")
		if err != nil {
			return nil, utils.UnauthenticatedError(err)
		}
	}

	access_token, accessPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Barber,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		return nil, utils.InternalError(err)
	}

	refresh_token, refreshPayload, err := server.tokenMaker.RefreshToken(
		payload.ID,
		payload.Barber,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, utils.InternalError(err)
	}

	arg := db.UpdateSessionRefreshTokenParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
	}
	err = server.Store.UpdateSessionRefreshToken(ctx, arg)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.RefreshTokenBarberResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiresAt),
		RefreshToken:          refresh_token,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}
	return rsp, nil
}
