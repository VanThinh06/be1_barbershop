package gapi

import (
	db "barbershop/db/sqlc"
	"barbershop/pb"
	"barbershop/token"
	"barbershop/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server servers gRPC requests for our barbershop
type Server struct {
	pb.UnimplementedBarberShopServer
	config     utils.Config
	queries    db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	Payload    *token.Payload
}

func NewServer(config utils.Config, queries db.StoreMain) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		queries:    queries,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
