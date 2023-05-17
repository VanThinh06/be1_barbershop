package api

import (
	db "barbershop/db/sqlc"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type newServiceParams struct {
	StoreID     uuid.UUID   `json:"store_id" binding:"required"`
	Work        string      `json:"work" binding:"required"`
	Timer       int32     `json:"timer" binding:"required"`
	Price       float32      `json:"price" binding:"required"`
	Description null.String `json:"description"`
	Image       null.String `json:"image"`
}

func (server *Server) createService(ctx *gin.Context) {
	var req newServiceParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	arg := db.CreateServiceParams{
		StoreID:     req.StoreID,
		Work:        req.Work,
		Timer:       req.Timer,
		Price:       req.Price,
		Description: req.Description,
		Image:       req.Image,
	}
	service, err := server.queries.CreateService(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}
