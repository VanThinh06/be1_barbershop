package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	customer.UnimplementedCustomerBarberShopServer
	config      utilities.Config
	store       db.StoreMain
	tokenMaker  token.CustomerMaker
	Router      *gin.Engine
	Payload     *token.CustomerPayload
	firebaseApp *firebase.App
}

func NewServer(config utilities.Config, store db.StoreMain, firebase db.FirebaseApp) (*Server, error) {
	tokenMaker, err := token.NewPasetoMakerCustomer(config.TokenSymmetricKeyCustomer)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:      config,
		store:       store,
		tokenMaker:  tokenMaker,
		firebaseApp: firebase.FirebaseApp,
	}
	return server, nil
}
