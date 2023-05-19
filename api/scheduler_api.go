package api

import (
	db "barbershop/db/sqlc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type newSchedulerWorkParams struct {
	BarberID   uuid.UUID   `json:"barber_id" binding:"required"`
	UsersID    string      `json:"users_id" binding:"required"`
	Timerstart string      `json:"timerstart" binding:"required"`
	Timerend   string      `json:"timerend" binding:"required"`
	Service    []uuid.UUID `json:"serviced"`
	TotalPrice float32     `json:"total_price" binding:"required"`
	Status     int32       `json:"status" binding:"required"`
}

func (server *Server) newSchedulerWork(ctx *gin.Context) {
	var req newSchedulerWorkParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	timeStart, err := time.Parse(time.RFC3339, req.Timerstart)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	timeEnd, err := time.Parse(time.RFC3339, req.Timerend)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSchedulerWorkParams{
		BarberID:   req.BarberID,
		UsersID:    req.UsersID,
		Timerstart: timeStart,
		Timerend:   timeEnd,
		Service:    req.Service,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}

	schedulerWork, err := server.queries.CreateSchedulerWork(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulerWork)

}
