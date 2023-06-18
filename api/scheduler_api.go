package api

import (
	db "barbershop/db/sqlc"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

		//
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				if fe.Field() == "Fcm_Device" {
					ctx.JSON(http.StatusBadRequest, gin.H{"errors": msgForTag(fe.Field())})
					return
				}
			}

		}
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				if pqErr.Constraint == "schedulerwork_barber_id_timerstart_idx" {
					ctx.JSON(http.StatusForbidden, "Working time already exists")
					return
				}
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, schedulerWork)

}
