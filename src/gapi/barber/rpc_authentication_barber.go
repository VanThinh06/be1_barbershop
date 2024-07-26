package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/shared/helpers"
	"barbershop/src/shared/token"
	"barbershop/src/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/grpc/codes"
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

	nickName := utils.GenerateNickName(req.FullName)
	arg := db.InsertBarberWithDetailsParams{
		NickName: strings.ToLower(nickName),
		FullName: req.GetFullName(),
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		Phone: req.GetPhone(),
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	res, err := server.Store.InsertBarberWithDetails(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.ConstraintName {
			case "Barbers_email_key":
				return nil, utils.AlreadyExistsError(err, "This email has already existed")
			case "Barbers_phone_key":
				return nil, utils.AlreadyExistsError(err, "This phone has already existed")
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

	arg := db.GetBarberWithRoleAndShopParams{
		ID:           idBarber,
		BarberShopID: idBarberShop,
	}
	res, err := server.Store.GetBarberWithRoleAndShop(ctx, arg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
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
	contact := db.GetBarberByCredentialParams{
		TypeUsername: "phone",
		Email: sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != "",
		},
	}

	err := helpers.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.TypeUsername = "email"
		contact.Email = sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != "",
		}
	}

	res, err := server.Store.GetBarberByCredential(ctx, contact)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.NotFoundError(err, "you have not created an account yet")
		}
		return nil, utils.InvalidError(err, "username or password is incorrect")
	}

	if !res.HashedPassword.Valid {
		password, err := server.tokenMaker.DecodeAESString(req.GetPassword())
		if err != nil {
			return nil, utils.InvalidError(err, "username or password is incorrect")
		}
		if password != res.DefaultPasswordEncrypted.String {
			return nil, utils.InvalidError(err, "username or password is incorrect")
		}
	} else {
		err = helpers.CheckPassword(req.Password, res.HashedPassword.String)
		if err != nil {
			return nil, utils.InvalidError(err, "username or password is incorrect")
		}
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
		ID:           refreshPayload.ID,
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.NotFoundError(err, "not found")
		}

		return nil, utils.UnauthenticatedError(err)
	}

	if session.IsBlocked {
		err = fmt.Errorf("incorrect block user")
		return nil, utils.UnauthenticatedError(err)

	}

	if session.BarberID != payload.Barber.BarberID {
		err = fmt.Errorf("incorrect session user")

		return nil, utils.UnauthenticatedError(err)

	}

	if session.RefreshToken != req.RefreshToken {
		err = fmt.Errorf("incorrect session user")

		return nil, utils.UnauthenticatedError(err)

	}

	if time.Now().UTC().After(session.ExpiresAt) {
		err = fmt.Errorf("expired session")
		return nil, utils.UnauthenticatedError(err)
	}

	// mtdt := server.extractMetadata(ctx)
	// if session.ClientIp != mtdt.ClientIP {
	// 	_ = fmt.Errorf("incorrect clientIP")
	// 	return nil, utils.UnauthenticatedError(err)
	// }

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

// authenticate barber forgot password
func (server *Server) ForgotPasswordBarber(ctx context.Context, req *barber.ForgotPasswordBarberRequest) (*barber.ForgotPasswordBarberResponse, error) {
	contact := db.GetBarberByCredentialParams{
		TypeUsername: "phone",
		Email: sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != "",
		},
	}

	err := helpers.ValidatePhoneNumber(req.Username)
	if err != nil {
		contact.TypeUsername = "email"
		contact.Email = sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != "",
		}
	}

	user, err := server.Store.GetBarberByCredential(ctx, contact)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.NotFoundError(err, "account does not exist")
		}
		return nil, utils.InvalidError(err, "email is incorrect")
	}

	countOTP, err := server.Store.GetCountOTPsForBarberToday(ctx, user.ID)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	if countOTP >= 5 {
		return nil, utils.ReturnError(codes.ResourceExhausted, err, "too many requests. please try again tomorrow.")
	}

	otp := utils.GenerateOTP()
	argOTP := db.InsertNewOTPRequestParams{
		BarberID: user.ID,
		Otp:      otp,
	}
	if !user.Email.Valid {
		return nil, utils.ReturnError(codes.FailedPrecondition, err, "invalid email. please contact the administrator to add an email.")
	}

	err = utils.SendOTP(user.Email.String, otp)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	err = server.Store.InsertNewOTPRequest(ctx, argOTP)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.ForgotPasswordBarberResponse{
		Message: user.ID.String(),
	}
	return rsp, nil
}

// authenticate barber verify otp
func (server *Server) VerifyOtpBarber(ctx context.Context, req *barber.VerifyOtpBarberRequest) (*barber.VerifyOtpBarberResponse, error) {

	barberId, err := uuid.Parse(req.GetBarberId())
	if err != nil {
		return nil, utils.NotFoundError(err, "barber don't exist")
	}
	argOtpRequest := db.GetOTPRequestDetailsParams{
		BarberID: barberId,
		Otp:      req.Code,
	}
	resOtp, err := server.Store.GetOTPRequestDetails(ctx, argOtpRequest)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, utils.ReturnError(codes.ResourceExhausted, err, "invalid or expired OTP")
		}
		return nil, utils.ReturnError(codes.Internal, err, "an error occurred while processing OTP")
	}

	if time.Now().UTC().Sub(resOtp.RequestedAt) > time.Minute*5 {
		return nil, utils.ReturnError(codes.ResourceExhausted, err, "invalid or expired OTP")
	}

	argConfirmOTPRequest := db.UpdateOTPRequestConfirmationParams{
		IsConfirm: true,
		ID:        resOtp.ID,
	}
	err = server.Store.UpdateOTPRequestConfirmation(ctx, argConfirmOTPRequest)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.VerifyOtpBarberResponse{
		Message: resOtp.Otp,
	}
	return rsp, nil
}

// authenticate barber reset password
func (server *Server) ResetPasswordBarber(ctx context.Context, req *barber.ResetPasswordBarberRequest) (*barber.ResetPasswordBarberResponse, error) {

	err := helpers.ValidatePassword(req.GetPassword())
	if err != nil {
		return nil, utils.InvalidError(err, "invalid password format")
	}

	barberId, err := uuid.Parse(req.GetBarberId())
	if err != nil {
		return nil, utils.NotFoundError(err, "barber don't exist")
	}

	respBarber, err := server.Store.GetBarberById(ctx, barberId)
	if err != nil {
		return nil, utils.NotFoundError(err, "barber don't exist")
	}

	if respBarber.HashedPassword.Valid {
		newPasswordErr := helpers.CheckPassword(req.Password, respBarber.HashedPassword.String)
		if newPasswordErr == nil {
			return nil, utils.InvalidError(nil, "New password must not be the same as the old password")
		}
	}

	argOtpRequest := db.GetOTPRequestDetailsParams{
		BarberID: barberId,
		Otp:      req.Code,
	}
	resOtp, err := server.Store.GetOTPRequestDetails(ctx, argOtpRequest)
	if err != nil {
		return nil, utils.InternalError(err)
	}
	if !resOtp.IsConfirm {
		return nil, utils.ReturnError(codes.ResourceExhausted, err, "invalid OTP")
	}

	hashedPassword, err := helpers.HashPassword(req.GetPassword())
	if err != nil {
		return nil, utils.InternalError(err)
	}
	argUpdatePasswordBarber := db.UpdateBarberPasswordParams{
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		ID: barberId,
	}
	resBarber, err := server.Store.UpdateBarberPassword(ctx, argUpdatePasswordBarber)
	if err != nil {
		return nil, utils.InternalError(err)
	}

	rsp := &barber.ResetPasswordBarberResponse{
		Message: resBarber.NickName,
	}
	return rsp, nil
}
