package api

import (
	db "barbershop/db/sqlc"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type NewServiceParams struct {
	ServiceCategoryID uuid.UUID   `json:"service_category_id" binding:"required"`
	Work              string      `json:"work" binding:"required"`
	Timer             null.Int    `json:"timer"`
	Price             float32     `json:"price" binding:"required"`
	Description       null.String `json:"description"`
	Image             null.String `json:"image"`
}

func (server *Server) CreateService(ctx *gin.Context) {
	var req NewServiceParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateServiceParams{
		ServiceCategoryID: req.ServiceCategoryID,
		Work:              req.Work,
		Timer:             req.Timer,
		Price:             req.Price,
		Description:       req.Description,
		Image:             req.Image,
	}
	service, err := server.queries.CreateService(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, service)
}

// * get list service with service category
type listServiceRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) GetListServicewithCategory(ctx *gin.Context) {
	// get with id
	id := ctx.Param("id")
	pageId, errPageID := strconv.Atoi(ctx.Query("page_id"))
	pageSize, errPageSize := strconv.Atoi(ctx.Query("page_size"))
	if errPageID != nil && errPageSize != nil {
		service, err := server.queries.GetService(ctx, uuid.MustParse(id))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusOK, service)
		return
	}

	// get list
	var req listServiceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetListServicewithCategoryParams{
		ServiceCategoryID: uuid.MustParse(id),
		Limit:             int32(pageId),
		Offset:            int32((pageId - 1) * pageSize),
	}

	listService, err := server.queries.GetListServicewithCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listService)
}
