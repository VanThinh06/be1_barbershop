package api

import (
	db "barbershop/db/sqlc"
	"barbershop/util"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ReNewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	ClientIP     string `json:"client_ip" binding:"required"`
}

type ReNewAccessTokenResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (server *Server) ReNewAccessToken(ctx *gin.Context) {
	var req ReNewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := util.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, util.MessageResponse("The request was invalid"))
		return
	}
	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.queries.GetSession(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.Username != payload.Username {
		err := fmt.Errorf("incorect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("incorect session user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if session.ClientIp != req.ClientIP {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	access_token, accessPayload, err := server.tokenMaker.CreateToken(
		payload.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refresh_token, refreshPayload, err := server.tokenMaker.CreateToken(
		payload.Username,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ssUpdate, err := server.queries.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
		ID:           payload.ID,
		RefreshToken: refresh_token,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	rsp := ReNewAccessTokenResponse{
		AccessToken:           access_token,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          ssUpdate.RefreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
