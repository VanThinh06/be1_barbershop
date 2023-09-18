package api

import (
	db "barbershop/db/sqlc"
	"barbershop/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type barberResponse struct {
	Username   string        `json:"username"`
	FullName   string        `json:"full_name"`
	Email      string        `json:"email"`
	Avatar     null.String   `json:"avatar"`
	Role       int32         `json:"role"`
	Status     null.Int      `json:"status"`
	StoreWork  uuid.NullUUID `json:"store_work"`
	TypeBarber int32         `json:"type_barber"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdateAt   null.Time     `json:"update_at"`
}

// * api get barber
func (server *Server) GetBarber(ctx *gin.Context) {
	id := ctx.Param("id")
	barber, err := server.queries.GetBarber(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.MessageResponse("This account does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}
	response := barberResponse{
		Username:  barber.Username,
		FullName:  barber.FullName,
		Email:     barber.Email,
		Avatar:    barber.Avatar,
		Role:      barber.Role,
		Status:    barber.Status,
		StoreWork: barber.StoreWork,
		CreatedAt: barber.CreatedAt,
	}
	ctx.JSON(http.StatusOK, response)
}

// * auth register
// create barber
type authNewBarberParams struct {
	Username   string        `json:"username" binding:"required"`
	FullName   string        `json:"full_name" binding:"required"`
	Email      string        `json:"email" binding:"email,required"`
	Password   string        `json:"password" binding:"required,min=6"`
	TypeBarber int32         `json:"type_barber" binding:"required"`
	StoreWork  uuid.NullUUID `json:"store_work"`
	Role       int32         `json:"role" binding:"required"`
}

func newBarberResponse(barber db.Barber) barberResponse {
	return barberResponse{
		Username:   barber.Username,
		FullName:   barber.FullName,
		Email:      barber.Email,
		CreatedAt:  barber.CreatedAt,
		TypeBarber: barber.TypeBarber,
		Role:       barber.Role,
		Avatar:     barber.Avatar,
		Status:     barber.Status,
		StoreWork:  barber.StoreWork,
		UpdateAt:   barber.UpdateAt,
	}
}

// * auth register
func (server *Server) AuthRegister(ctx *gin.Context) {
	var req authNewBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := util.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, util.MessageResponse("The request was invalid"))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	arg := db.CreateBarberParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		StoreWork:      req.StoreWork,
		Role:           req.Role,
		TypeBarber:     req.TypeBarber,
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
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}

// * auth login
// login accout params
type authBarberParams struct {
	Username string `json:"username" binding:"required"`
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
	var req authBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := util.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, util.MessageResponse("Login information is incorrect"))
		return
	}

	// get barber with user name
	barber, err := server.queries.GetBarber(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.MessageResponse("Incorrect account or password"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	// check password
	err = util.CheckPassword(req.Password, barber.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, util.MessageResponse("Incorrect account or password"))
		return
	}

	// create token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barber.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	// create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barber.Username,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
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
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
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

// update
type updateBarberParams struct {
	Username          string        `json:"username"`
	Status            null.Int      `json:"status"`
	FullName          string        `json:"full_name"`
	Role              int32         `json:"role"`
	StoreWork         uuid.NullUUID `json:"store_work"`
	TypeBarber        int32         `json:"type_barber"`
	Email             string        `json:"email"`
	Avatar            null.String   `json:"avatar"`
	PasswordChangedAt time.Time     `json:"password_changed_at"`
	UpdateAt          null.Time     `json:"update_at"`
}

func (server *Server) UpdateBarber(ctx *gin.Context) {
	var req updateBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "Update invalid information")
		return
	}

	params := db.UpdateBarberParams{
		Username:   req.Username,
		Status:     req.Status,
		FullName:   req.FullName,
		Role:       req.Role,
		StoreWork:  req.StoreWork,
		TypeBarber: req.TypeBarber,
		Email:      req.Email,
		Avatar:     req.Avatar,
		UpdateAt:   null.TimeFrom(time.Now()),
	}
	barber, err := server.queries.UpdateBarber(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.MessageResponse("Update invalid information"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}
