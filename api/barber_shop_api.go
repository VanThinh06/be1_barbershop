package api

import (
	db "barbershop/db/sqlc"
	"barbershop/util"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// * api get barber
func (server *Server) GetBarberShop(ctx *gin.Context) {
	id := ctx.Param("id")
	barberShop, err := server.queries.GetBarberShop(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.MessageResponse("This account does not exist"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}
	response := barberShop
	ctx.JSON(http.StatusOK, response)
}

type barberShopsParams struct {
	CodeBarberShop string  `json:"code_barber_shop" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Location       float32 `json:"location" binding:"required"`
	Address        string  `json:"address" binding:"required"`
}

func (server *Server) NewBarberShop(ctx *gin.Context) {
	var req barberShopsParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := util.CatchErrorParams(err)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusBadRequest, util.MessageResponse("The request was invalid"))
		return
	}
	fmt.Println(server.payload.BarberID)
	arg := db.CreateBarberShopsParams{
		CodeBarberShop: req.CodeBarberShop,
		OwnerID:        server.payload.BarberID,
		Name:           req.Name,
		Location:       req.Location,
		Address:        req.Address,
	}

	barberShop, err := server.queries.CreateBarberShops(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	requestBarber := db.UpdateIDShopManagerParams{BarberID: server.payload.BarberID, ShopID: uuid.NullUUID{UUID: barberShop.ShopID, Valid: true}}
	_, errUpdateBarber := server.queries.UpdateIDShopManager(ctx, requestBarber)
	if errUpdateBarber != nil {
		ctx.JSON(http.StatusInternalServerError, util.MessageInternalServer)
		return
	}

	response := barberShop
	ctx.JSON(http.StatusOK, response)
}
