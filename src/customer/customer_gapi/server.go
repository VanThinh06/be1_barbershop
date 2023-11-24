package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"
	"barbershop/src/shared/token"
	"barbershop/src/shared/utils"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	pb.UnimplementedAuthCustomerBarberShopServer
	config     utils.Config
	store      db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	Payload    *token.Payload
	firebaseApp   *firebase.App
}

func NewServer(config utils.Config, store db.StoreMain, firebase db.FirebaseApp) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		firebaseApp: firebase.FirebaseApp,
	}
	return server, nil
}
