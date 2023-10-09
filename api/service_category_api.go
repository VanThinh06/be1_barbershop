package api

// import (
// 	db "barbershop/db/sqlc"
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"gopkg.in/guregu/null.v4"
// )

// type newServiceCategoryParams struct {
// 	StoreID     uuid.UUID   `json:"store_id" binding:"required"`
// 	Work        string      `json:"work" binding:"required"`
// 	Description null.String `json:"description"`
// }

// func (server *Server) CreateServiceCategory(ctx *gin.Context) {
// 	var req newServiceCategoryParams

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	response := db.CreateServiceCategoryParams{
// 		StoreID:     req.StoreID,
// 		Work:        req.Work,
// 		Description: req.Description,
// 	}
// 	serviceCategory, err := server.queries.CreateServiceCategory(ctx, response)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, serviceCategory)
// }

// func (server *Server) GetListServiceCategorywithStore(ctx *gin.Context) {
// 	// get with id
// 	id := ctx.Param("id_store")

// 	listServiceCategory, err := server.queries.GetListServiceCategorywithStore(ctx, uuid.MustParse(id))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, listServiceCategory)
// }

// func (server *Server) DeleteServiceCategory(ctx *gin.Context) {
// 	// get with id
// 	id := ctx.Param("id")
// 	err := server.queries.DeleteServicewithStoreCategory(ctx, uuid.MustParse(id))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	serviceCategory, err := server.queries.DeleteServiceCategory(ctx, uuid.MustParse(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, "Delete serviceCategory "+serviceCategory.Work+" successfully")
// }
