package api

// import (
// 	db "barbershop/db/sqlc"
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"gopkg.in/guregu/null.v4"
// )

// type Status int

// const (
// 	Work    Status = 1
// 	Off     Status = 2
// 	Destroy Status = 3
// )

// type createStoreParams struct {
// 	NameStore  string `json:"name_store" binding:"required"`
// 	ManagerID  string `json:"manager_id"`
// 	EmployeeID string `json:"employee_id"`
// 	Location   int64  `json:"location"`
// 	Status     string `json:"status" binding:"required,oneof=Work Off Destroy"`
// }

// func (server *Server) NewStore(ctx *gin.Context) {
// 	var req createStoreParams
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateStoreParams{
// 		NameStore:  null.StringFrom(req.NameStore),
// 		ManagerID:  uuid.MustParse(req.ManagerID),
// 		EmployeeID: uuid.MustParse(req.EmployeeID),
// 		Location:   null.IntFrom(req.Location),
// 		Status:     null.StringFrom(req.Status),
// 	}
// 	store, err := server.queries.CreateStore(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, store)
// }

// func nillUUID(a sql.NullString) string {
// 	if len(a.String) == 0 {
// 		return uuid.Nil.String()
// 	}
// 	return a.String
// }
