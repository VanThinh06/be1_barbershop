package api

// import (
// 	db "barbershop/db/sqlc"
// 	"barbershop/utils"
// 	"database/sql"
// 	"strconv"
// 	"time"

// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"gopkg.in/guregu/null.v4"
// )

// // * api create barber
// type newStoreParams struct {
// 	NameID         string      `json:"name_id"  binding:"required"`
// 	NameStore      string      `json:"name_store" binding:"required"`
// 	Location       float32     `json:"location" binding:"required"`
// 	Address        string      `json:"address" binding:"required"`
// 	Image          null.String `json:"image"`
// 	ListImageStore []string    `json:"list_image_store" `
// 	ManagerID      []string    `json:"manager_id"`
// 	EmployeeID     []string    `json:"employee_id"`
// 	Status         int32       `json:"status" binding:"required"`
// }

// func (server *Server) NewStore(ctx *gin.Context) {
// 	var req newStoreParams

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		err := utils.CatchErrorParams(err)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, err)
// 			return
// 		}
// 		ctx.JSON(http.StatusBadRequest, utils.MessageResponse("The request was invalid"))
// 		return
// 	}
// 	response := db.CreateStoreParams{
// 		NameID:     req.NameID,
// 		NameStore:  req.NameStore,
// 		Location:   req.Location,
// 		Address:    req.Address,
// 		Image:      req.Image,
// 		ListImage:  req.ListImageStore,
// 		ManagerID:  []string{server.payload.Username},
// 		EmployeeID: req.EmployeeID,
// 		Status:     req.Status,
// 		Boss:       server.payload.Username,
// 	}
// 	store, err := server.queries.CreateStore(ctx, response)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, store)
// }

// // * api Get Store
// type storeResponse struct {
// 	NameID     string      `json:"name_id"`
// 	NameStore  string      `json:"name_store"`
// 	Location   float32     `json:"location"`
// 	Address    string      `json:"address"`
// 	Image      null.String `json:"image"`
// 	ListImage  []string    `json:"list_image"`
// 	ManagerID  []string    `json:"manager_id"`
// 	EmployeeID []string    `json:"employee_id"`
// 	Status     int32       `json:"status"`
// }

// func (server *Server) GetStore(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	store, err := server.queries.GetStore(ctx, uuid.MustParse(id))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	response := storeResponse{
// 		NameID:     store.NameID,
// 		NameStore:  store.NameStore,
// 		Location:   store.Location,
// 		Address:    store.Address,
// 		Image:      store.Image,
// 		ListImage:  store.ListImage,
// 		ManagerID:  store.ManagerID,
// 		EmployeeID: store.EmployeeID,
// 		Status:     store.Status,
// 	}
// 	ctx.JSON(http.StatusOK, response)
// }

// type listStoreRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) GetListStore(ctx *gin.Context) {
// 	var req listStoreRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	pageId, errPageID := strconv.Atoi(ctx.Query("page_id"))
// 	pageSize, errPageSize := strconv.Atoi(ctx.Query("page_size"))
// 	if errPageID != nil {
// 		return
// 	}
// 	if errPageSize != nil {
// 		return
// 	}
// 	arg := db.GetListStoreParams{
// 		Limit:     int32(pageSize),
// 		Offset:    int32((pageId - 1) * pageSize),
// 		ManagerID: server.payload.Username,
// 	}

// 	listBarber, err := server.queries.GetListStore(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, listBarber)
// }

// // update
// type updateStoreParams struct {
// 	ID         uuid.UUID   `json:"id" binding:"required"`
// 	NameID     string      `json:"name_id"`
// 	NameStore  string      `json:"name_store"`
// 	Location   float32     `json:"location"`
// 	Address    string      `json:"address"`
// 	Image      null.String `json:"image"`
// 	ListImage  []string    `json:"list_image"`
// 	ManagerID  []string    `json:"manager_id"`
// 	EmployeeID []string    `json:"employee_id"`
// 	Status     int32       `json:"status" `
// }

// func (server *Server) UpdateStore(ctx *gin.Context) {
// 	var req updateStoreParams
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	response := db.UpdateStoreParams{
// 		ID:         req.ID,
// 		NameID:     req.NameID,
// 		NameStore:  req.NameStore,
// 		Location:   req.Location,
// 		Address:    req.Address,
// 		Image:      req.Image,
// 		ListImage:  req.ListImage,
// 		ManagerID:  req.ManagerID,
// 		EmployeeID: req.EmployeeID,
// 		Status:     req.Status,
// 		UpdateAt:   null.TimeFrom(time.Now()),
// 	}
// 	store, err := server.queries.UpdateStore(ctx, response)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, store)
// }

// func (server *Server) DeleteStore(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	store, err := server.queries.GetStore(ctx, uuid.MustParse(id))
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, "Store not found")
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	if utils.IsAvailable(store.ManagerID, server.payload.Username) {
// 		store, err := server.queries.DeleteStore(ctx, uuid.MustParse(id))
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				ctx.JSON(http.StatusNotFound, errorResponse(err))
// 				return
// 			}
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, "Delete success store "+store.NameStore)
// 		return
// 	}
// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
// }
