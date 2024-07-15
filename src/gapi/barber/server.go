package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	barber.UnimplementedBarberServiceServer
	customer.UnimplementedCustomerBarberShopServer
	config     utils.Config
	Store      db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	Payload    *token.BarberPayload
	ServiceApp db.ServiceApp
}

func NewServer(config utils.Config, store db.StoreMain, serviceApp db.ServiceApp) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		Store:      store,
		tokenMaker: tokenMaker,
		ServiceApp: serviceApp,
	}
	return server, nil
}
