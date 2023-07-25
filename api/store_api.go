package api

import (
	db "barbershop/db/sqlc"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"gopkg.in/guregu/null.v4"
)

type StatusStore interface {
	IsValid() bool
}

type Status int

const (
	Active Status = iota + 1
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

func (s Status) NameStatusStore() int32 {
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
func ValidateStatusStore(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(StatusStore)
	return value.IsValid()
}

type newStoreParams struct {
	NameID     string          `json:"name_id" binding:"required"`
	NameStore  string          `json:"name_store" binding:"required"`
	Location   sql.NullFloat64 `json:"location"`
	Image      null.String     `json:"image"`
	ManagerID  uuid.NullUUID   `json:"manager_id"`
	EmployeeID []uuid.UUID     `json:"employee_id"`
	Status     Status          `json:"status" binding:"required,statusStore"`
}

func (server *Server) newStore(ctx *gin.Context) {
	var req newStoreParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag())}
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errorResponse(err)})
		return
	}

	arg := db.CreateStoreParams{
		NameID:     req.NameID,
		NameStore:  req.NameStore,
		Location:   req.Location,
		Image:      req.Image,
		ManagerID:  req.ManagerID,
		EmployeeID: req.EmployeeID,
		Status:     req.Status.NameStatusStore(),
	}

	store, err := server.queries.CreateStore(ctx, arg)
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
	ctx.JSON(http.StatusOK, store)
}
