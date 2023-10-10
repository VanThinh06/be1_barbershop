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
	BarberID  uuid.UUID     `json:"barber_id"`
	ShopID    uuid.NullUUID `json:"shop_id"`
	NickName  string        `json:"nick_name"`
	FullName  string        `json:"full_name"`
	Phone     string        `json:"phone"`
	Email     string        `json:"email"`
	Gender    int32         `json:"gender"`
	Role      int32         `json:"role"`
	Avatar    null.String   `json:"avatar"`
	Status    null.Int      `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdateAt  null.Time     `json:"update_at"`
}

func newBarberResponse(barber db.Barber) barberResponse {
	return barberResponse{
		BarberID:  barber.BarberID,
		ShopID:    barber.ShopID,
		NickName:  barber.NickName,
		FullName:  barber.FullName,
		Phone:     barber.Phone,
		Email:     barber.Email,
		Gender:    barber.Gender,
		Role:      barber.Role,
		Avatar:    barber.Avatar,
		Status:    barber.Status,
		CreatedAt: barber.CreatedAt,
		UpdateAt:  barber.UpdateAt,
	}
}

// * api get barber
func (server *Server) GetBarber(ctx *gin.Context) {
	id := ctx.Param("email")
	barber, err := server.queries.GetEmailBarber(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.MessageResponse("This account does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}

// * auth register
// create barber
type authNewBarberParams struct {
	ShopID   uuid.NullUUID `json:"shop_id"`
	NickName string        `json:"nickname" binding:"required"`
	FullName string        `json:"full_name" binding:"required"`
	Phone    string        `json:"phone" binding:"required"`
	Email    string        `json:"email" binding:"email,required"`
	Gender   int           `json:"gender" binding:"required"`
	Role     int           `json:"role" binding:"required"`
	Password string        `json:"password" binding:"required,min=6"`
	Avatar   null.String   `json:"avatar"`
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
		ShopID:         req.ShopID,
		NickName:       req.NickName,
		FullName:       req.FullName,
		Phone:          req.Phone,
		Email:          req.Email,
		Gender:         int32(req.Gender),
		Role:           int32(req.Role),
		HashedPassword: hashedPassword,
		Avatar:         req.Avatar,
	}

	barber, err := server.queries.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "Barbers_pkey" {
					ctx.JSON(http.StatusForbidden, "This account has already existed")
					return
				}
				if pqErr.Constraint == "Barbers_email_key" {
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
// login account params
type authBarberParams struct {
	Email    string `json:"email" binding:"email,required"`
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
	barber, err := server.queries.GetEmailBarber(ctx, req.Email)
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
		barber.BarberID,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	// create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barber.BarberID,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	session, err := server.queries.CreateSessionBarber(ctx, db.CreateSessionBarberParams{
		ID:           refreshPayload.ID,
		BarberID:     refreshPayload.BarberID,
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
	BarberID uuid.UUID     `json:"barber_id" binding:"required"`
	ShopID   uuid.NullUUID `json:"shop_id"`
	NickName string        `json:"nick_name"`
	FullName string        `json:"full_name"`
	Phone    string        `json:"phone" `
	Email    string        `json:"email" `
	Gender   int           `json:"gender" `
	Role     int           `json:"role" `
	Avatar   null.String   `json:"avatar"`
	Status   null.Int      `json:"status"`
	UpdateAt null.Time     `json:"update_at"`
}

func (server *Server) UpdateBarber(ctx *gin.Context) {
	var req updateBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, "Update invalid information")
		return
	}

	params := db.UpdateBarberParams{
		BarberID: req.BarberID,
		ShopID:   req.ShopID,
		NickName: req.NickName,
		Status:   req.Status,
		FullName: req.FullName,
		Role:     int32(req.Role),
		Phone:    req.Phone,
		Gender:   int32(req.Gender),
		Email:    req.Email,
		Avatar:   req.Avatar,
		UpdateAt: null.TimeFrom(time.Now()),
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
