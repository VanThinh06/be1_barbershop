package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	customer.UnimplementedCustomerBarberShopServer
	config     utils.Config
	store      db.StoreMain
	tokenMaker token.CustomerMaker
	Router     *gin.Engine
	Payload    *token.CustomerPayload
	serviceApp db.ServiceApp
}

func NewServer(config utils.Config, store db.StoreMain, serviceApp db.ServiceApp) (*Server, error) {
	tokenMaker, err := token.NewPasetoMakerCustomer(config.TokenSymmetricKeyCustomer)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		serviceApp: serviceApp,
	}
	return server, nil
}
