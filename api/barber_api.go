package api

import (
	db "barbershop/db/sqlc"
	"barbershop/utils"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	Avatar    sql.NullString   `json:"avatar"`
	Status    sql.NullInt32      `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdateAt  sql.NullTime     `json:"update_at"`
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
	id := ctx.Param("id")
	barber, err := server.queries.GetEmailBarber(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.MessageResponse("This account does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
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
		err := utils.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse("Login information is incorrect"))
		return
	}

	// get barber with user name
	barber, err := server.queries.GetEmailBarber(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.MessageResponse("Incorrect account or password"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}

	// check password
	err = utils.CheckPassword(req.Password, barber.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.MessageResponse("Incorrect account or password"))
		return
	}

	// create token
	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		barber.BarberID,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}

	// create refresh token
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		barber.BarberID,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
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
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
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

// * auth register
// create barber
type newBarberParams struct {
	NickName       string `json:"nickname" binding:"required"`
	FullName       string `json:"full_name" binding:"required"`
	Phone          string `json:"phone" binding:"phone,required"`
	Email          string `json:"email" binding:"email,required"`
	Gender         int    `json:"gender" binding:"required"`
	Password       string `json:"password" binding:"required,min=6"`
	CodeBarberShop string `json:"code_barber_shop" binding:"required"`
}

// * auth register barber
func (server *Server) AuthNewBarber(ctx *gin.Context) {
	var req newBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := utils.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse("The request was invalid"))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}

	shop_id, err := server.queries.GetCodeBarberShop(ctx, req.CodeBarberShop)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.MessageResponse("Not found barber shop"))
		return
	}

	arg := db.CreateBarberParams{
		ShopID: uuid.NullUUID{
			UUID:  shop_id,
			Valid: true,
		},
		NickName:       req.NickName,
		FullName:       req.FullName,
		Phone:          req.Phone,
		Email:          strings.ToLower(req.Email),
		Gender:         int32(req.Gender),
		Role:           int32(utils.HairStylist),
		HashedPassword: hashedPassword,
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
				if pqErr.Constraint == "Barbers_phone_key" {
					ctx.JSON(http.StatusForbidden, "This phone has already existed")
					return
				}
				if pqErr.Constraint == "Barbers_email_key" {
					ctx.JSON(http.StatusForbidden, "This email has already existed")
					return
				}
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}

type newManagerParams struct {
	NickName string `json:"nickname" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"phone,required"`
	Email    string `json:"email" binding:"email,required"`
	Gender   int    `json:"gender" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// * auth register manager
func (server *Server) AuthNewManager(ctx *gin.Context) {
	var req newManagerParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := utils.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, utils.MessageResponse("The request was invalid"))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}

	arg := db.CreateBarberParams{
		NickName:       req.NickName,
		FullName:       req.FullName,
		Phone:          req.Phone,
		Email:          req.Email,
		Gender:         int32(req.Gender),
		Role:           int32(utils.Manager),
		HashedPassword: hashedPassword,
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
				if pqErr.Constraint == "Barbers_phone_key" {
					ctx.JSON(http.StatusForbidden, "This phone has already existed")
					return
				}
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
		return
	}
	response := newBarberResponse(barber)
	ctx.JSON(http.StatusOK, response)
}

// update
// type updateBarberParams struct {
// 	BarberID uuid.UUID     `json:"barber_id" binding:"required"`
// 	ShopID   uuid.NullUUID `json:"shop_id"`
// 	NickName string        `json:"nick_name"`
// 	FullName string        `json:"full_name"`
// 	Phone    string        `json:"phone" `
// 	Email    string        `json:"email" `
// 	Gender   int           `json:"gender" `
// 	Role     int           `json:"role" `
// 	Avatar   sql.NullString   `json:"avatar"`
// 	Status   sql.NullInt32      `json:"status"`
// 	UpdateAt sql.NullTime     `json:"update_at"`
// }

// func (server *Server) UpdateBarber(ctx *gin.Context) {
// 	var req updateBarberParams
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, "Update invalid information")
// 		return
// 	}

// 	params := db.UpdateBarberParams{
// 		BarberID: req.BarberID,
// 		ShopID:   req.ShopID,
// 		NickName: req.NickName,
// 		Status:   req.Status,
// 		FullName: req.FullName,
// 		Role:     int32(req.Role),
// 		Phone:    req.Phone,
// 		Gender:   int32(req.Gender),
// 		Email:    req.Email,
// 		Avatar:   req.Avatar,
// 		UpdateAt: null.TimeFrom(time.Now()),
// 	}
// 	barber, err := server.queries.UpdateBarber(ctx, params)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, utils.MessageResponse("Update invalid information"))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, utils.MessageInternalServer)
// 		return
// 	}
// 	response := newBarberResponse(barber)
// 	ctx.JSON(http.StatusOK, response)
// }
