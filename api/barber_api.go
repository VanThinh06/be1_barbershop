package api

import (
	db "barbershop/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type newBarberParams struct {
	Username       string        `json:"username" binding:"required"`
	FullName       string        `json:"full_name" binding:"required"`
	Email          string        `json:"email" validate:"email, required"`
	HashedPassword string        `json:"hashed_password" binding:"required"`
	Avatar         null.String   `json:"avatar"`
	Role           null.String   `json:"role"`
	Status         null.String   `json:"status"`
	StoreID        uuid.NullUUID `json:"store_id"`
	StoreManager   []uuid.UUID   `json:"store_manager"`
}

func (server *Server) NewBarber(ctx *gin.Context) {
	var req newBarberParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBarberParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: req.HashedPassword,
		Avatar:         req.Avatar,
		StoreID:        req.StoreID,
		StoreManager:   req.StoreManager,
	}

	barber, err := server.queries.CreateBarber(ctx, arg)
	if err != nil {
		// if pqErr, ok := err.(*pq.Error); ok {
		// 	switch pqErr.Code.Name() {
		// 	case "foreign_key_violation", "unique_violation":
		// 		ctx.JSON(http.StatusForbidden, errorResponse(err))
		// 		return
		// 	}
		// }
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, barber)
}

func (server *Server) GetBarber(ctx *gin.Context) {
	id := ctx.Param("id")
	barber, err := server.queries.GetBarber(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, barber)
}
