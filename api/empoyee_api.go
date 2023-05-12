package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type NameEmployee struct {
// 	Name string `uri:"name" binding:"required" fieldexcludes:"a"`
// }

// func (sever *Server) GetEmployeeWithName(ctx *gin.Context) {
// 	var req NameEmployee
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	employee, err := sever.queries.GetEmployeeWithName(ctx, req.Name)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"employee": employee})

// }
