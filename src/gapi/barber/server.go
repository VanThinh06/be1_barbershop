package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	barber.UnimplementedBarberServiceServer
	customer.UnimplementedCustomerBarberShopServer
	config     utilities.Config
	Store      db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	Payload    *token.BarberPayload
	firebase   *firebase.App
}

func NewServer(config utilities.Config, store db.StoreMain, firebase db.FirebaseApp) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		Store:      store,
		tokenMaker: tokenMaker,
		firebase:   firebase.FirebaseApp,
	}
	return server, nil
}
