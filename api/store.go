package api

import (
	db "barbershop/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type Enum interface {
	IsValid() bool
}

type Status int

const (
	Active Status = iota + 1 // add + 1 otherwise validation won't work for 0
	Shutdown
	Stop
)

func (s Status) IsValid() bool {
	switch s {
	case Active, Shutdown, Stop:
		return true
	}
	return false
}

func (s Status) NameStatus() int32 {
	switch s {
	case Active:
		return 1
	case Shutdown:
		return 2
	case Stop:
		return 3
	}
	return 1
}

type newUserParams struct {
	NameID     string        `json:"name_id" binding:"required"`
	NameStore  string        `json:"name_store" binding:"required"`
	Location   null.Int      `json:"location"`
	Image      null.String   `json:"image"`
	ManagerID  uuid.NullUUID `json:"manager_id"`
	EmployeeID []uuid.UUID   `json:"employee_id"`
	Status     Status        `json:"status" binding:"required,enum"`
}

func (server *Server) newStore(ctx *gin.Context) {
	var req newUserParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStoreParams{
		NameID:     req.NameID,
		NameStore:  req.NameStore,
		Location:   req.Location,
		Image:      req.Image,
		ManagerID:  req.ManagerID,
		EmployeeID: req.EmployeeID,
		Status:     req.Status.NameStatus(),
	}

	store, err := server.queries.CreateStore(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			log.Println(ok)
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, store)

}
