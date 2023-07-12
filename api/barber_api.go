package api

import (
	db "barbershop/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type newBarberParams struct {
	NameID       string        `json:"name_id" binding:"required"`
	StoreID      uuid.NullUUID `json:"store_id" `
	StoreManager []uuid.UUID   `json:"store_manager"`
}

func (server *Server) newBarber(ctx *gin.Context) {
	var req newBarberParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBarberParams{
		NameID:       req.NameID,
		StoreID:      req.StoreID,
		StoreManager: req.StoreManager,
	}

	barber, err := server.queries.CreateBarber(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, barber)

}
