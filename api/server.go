package api

import (
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"barbershop/token"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	queries    db.StoreMain
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, queries db.StoreMain) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		queries:    queries,
		tokenMaker: tokenMaker,
	}
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//     v.RegisterValidation("currency", validCurrency)
	// }
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("enum", ValidateEnum)
	}
	server.setupRouter()

	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ValidateEnum(fl validator.FieldLevel) bool {
	value := fl.Field().Interface().(Enum)
	return value.IsValid()
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users/login", server.loginUser)
	router.POST("/tokent/refresh_access", server.reNewAccessToken)
	router.POST("/users", server.createUser)

	authRoutes := router.Group("/").Use(addMiddleWare(server.tokenMaker))
	authRoutes.POST("/barber", server.newBarber)
	authRoutes.POST("/store", server.newStore)

	server.router = router
}
