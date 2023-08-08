package api

import (
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"barbershop/token"
	"fmt"

	"github.com/gin-gonic/gin"
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
	// 	v.RegisterValidation("statusStore", ValidateStatusStore)
	// }
	server.setupRouter()
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// / Router
func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/barber", server.AuthRegister)
	router.POST("/barber/login", server.LoginBarber)
	router.GET("/barber/:id", server.GetBarber)
	router.POST("/token/refresh_access", server.ReNewAccessToken)

	// router.POST("/image", server.uploadImage)
	// router.Static("/assets", "./assets/upload")
	// router.POST("/users", server.createUser)

	authRoutes := router.Group("/").Use(AddMiddleWare(server.tokenMaker))
	authRoutes.POST("/store", server.NewStore)
	authRoutes.GET("/store/:id", server.GetStore)
	authRoutes.GET("/store", server.GetListStore)

	authRoutes.GET("/service_category/:id_store", server.GetListServiceCategorywithStore)
	authRoutes.POST("/service_category", server.CreateServiceCategory)
	authRoutes.DELETE("/service_category/:id", server.DeleteServiceCategory)

	authRoutes.GET("/service/:id", server.GetListServicewithCategory)
	authRoutes.POST("/service", server.CreateService)
	// router.GET("/sdkNotification", sdkFirebaseAdmin) // todo

	server.router = router
}

// ERROR
type ApiError struct {
	Field string
	Msg   string
}
