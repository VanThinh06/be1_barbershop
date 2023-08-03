package api

import (
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type barberResponse struct {
	Username          string        `json:"username"`
	FullName          string        `json:"full_name"`
	Email             string        `json:"email"`
	Avatar            null.String   `json:"avatar"`
	Role              null.String   `json:"role"`
	Status            null.String   `json:"status"`
	StoreWork         uuid.NullUUID `json:"store_work"`
	StoreManager      []uuid.UUID   `json:"store_manager"`
	PasswordChangedAt time.Time     `json:"password_changed_at"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdateAt          null.Time     `json:"update_at"`
}

// * api get barber
func (server *Server) GetBarber(ctx *gin.Context) {
	id := ctx.Param("id")
	barber, err := server.queries.GetBarber(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := barberResponse{
		Username:          barber.Username,
		FullName:          barber.FullName,
		Email:             barber.Email,
		Avatar:            barber.Avatar,
		Role:              barber.Role,
		Status:            barber.Status,
		StoreWork:         barber.StoreWork,
		StoreManager:      barber.StoreManager,
		PasswordChangedAt: barber.PasswordChangedAt,
		CreatedAt:         barber.CreatedAt,
		UpdateAt:          barber.UpdateAt,
	}
	ctx.JSON(http.StatusOK, response)
}

// * auth register
// create barber
type newBarberParams struct {
	Username  string        `json:"username" binding:"required,alphanum"`
	FullName  string        `json:"full_name" binding:"required"`
	Email     string        `json:"email" binding:"email,required"`
	Password  string        `json:"password" binding:"required,min=6"`
	Avatar    null.String   `json:"avatar"`
	Role      null.String   `json:"role"`
	Status    null.String   `json:"status"`
	StoreWork uuid.NullUUID `json:"store_work"`
}

func newBarberResponse(barber db.Barber) barberResponse {
	return barberResponse{
		Username:          barber.Username,
		FullName:          barber.FullName,
		Email:             barber.Email,
		PasswordChangedAt: barber.PasswordChangedAt,
		CreatedAt:         barber.CreatedAt,
	}
}

// * auth register
func (server *Server) NewBarber(ctx *gin.Context) {
	var req newBarberParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validateError validator.ValidationErrors
		if errors.As(err, &validateError) {
			for _, fe := range validateError {

				if fe.Field() == "Email" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": util.MsgForTag(fe.Field())})
					return
				}
				if fe.Field() == "Password" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": util.MsgForTag(fe.Field())})
					return
				}
				if fe.Field() == "UserName" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Username is only for alphanumeric input, no special characters"})
					return
				}
			}
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateBarberParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Avatar:         req.Avatar,
		StoreWork:      req.StoreWork,
	}

	barber, err := server.queries.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "barber_pkey" {
					ctx.JSON(http.StatusForbidden, "This account has already existed")
					return
				}
				if pqErr.Constraint == "barber_email_key" {
					ctx.JSON(http.StatusForbidden, "This email has already existed")
					return
				}
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}

// * auth login
// login accout params
type LoginAccoutBarberParams struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
type BarberLoginResponse struct {
	SessionId             uuid.UUID      `json:"session_id"`
	AccessToken           string         `json:"access_token"`
	AccessTokenExpiresAt  time.Time      `json:"access_token_expires_at"`
	RefreshToken          string         `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time      `json:"refresh_token_expires_at"`
	Barber                barberResponse `json:"barber"`
}

// * auth login
func (server *Server) LoginBarber(ctx *gin.Context) {
	// check login
	var req LoginAccoutBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validateError validator.ValidationErrors
		if errors.As(err, &validateError) {
			for _, fe := range validateError {
				if fe.Field() == "Password" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Password must be more than 6 characters"})
					return
				}
				if fe.Field() == "UserName" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Username is only for alphanumeric input, no special characters"})
					return
				}
			}
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get barber with user name
	barber, err := server.queries.GetBarber(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// check password
	err = util.CheckPassword(req.Password, barber.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// create token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barber.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barber.Username,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.queries.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
		ID:           refreshPayload.ID,
		Username:     refreshPayload.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false, // chặn session
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := BarberLoginResponse{
		SessionId:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Barber:                newBarberResponse(barber),
	}
	ctx.JSON(http.StatusOK, response)
}
