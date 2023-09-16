package api

import (
	db "barbershop/db/sqlc"
	"barbershop/token"
	"barbershop/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	queries    db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	payload    *token.Payload
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
	return server.Router.Run(address)
}

// / Router
func (server *Server) setupRouter() {
	router := gin.Default()

	barberRoutes := router.Group("/barber")
	barberRoutes.POST("/authentication", server.LoginBarber)
	barberRoutes.POST("/newBarber", server.AuthRegister)
	barberRoutes.GET("/barber/:id", server.GetBarber)
	barberRoutes.POST("/token/refresh_access", server.ReNewAccessToken)

	// router.POST("/image", server.uploadImage)
	// router.Static("/assets", "./assets/upload")
	// router.POST("/users", server.createUser)

	authRoutes := router.Group("/").Use(server.AddMiddleWare(server.tokenMaker))
	authRoutes.POST("/store", server.NewStore)
	authRoutes.GET("/store/:id", server.GetStore)
	authRoutes.GET("/store", server.GetListStore)
	authRoutes.PUT("/store", server.UpdateStore)
	authRoutes.DELETE("/store/:id", server.DeleteStore)

	authRoutes.GET("/service_category/:id_store", server.GetListServiceCategorywithStore)
	authRoutes.POST("/service_category", server.CreateServiceCategory)
	authRoutes.DELETE("/service_category/:id", server.DeleteServiceCategory)

	authRoutes.GET("/service/:id", server.GetListServicewithCategory)
	authRoutes.POST("/service", server.CreateService)
	authRoutes.PUT("/service", server.UpdateService)
	authRoutes.DELETE("/service/:id", server.DeleteService)

	// router.GET("/sdkNotification", sdkFirebaseAdmin) // todo

	server.Router = router
}
