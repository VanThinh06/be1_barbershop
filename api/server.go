package api

import (
	db "barbershop/db/sqlc"
	"barbershop/helpers"
	"barbershop/token"
	"barbershop/utils"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	queries    db.StoreMain
	tokenMaker token.Maker
	Router     *gin.Engine
	payload    *token.Payload
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", helpers.ValidatePhoneNumber)
	}
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	barberRoutes := router.Group("/barber")
	barberRoutes.POST("/authentication", server.LoginBarber)
	barberRoutes.POST("/newBarber", server.AuthNewBarber)
	barberRoutes.POST("/newManager", server.AuthNewManager)
	barberRoutes.GET("/:id", server.GetBarber)

	barberRoutes.Use(server.AddMiddleWare(server.tokenMaker)).POST("/updateBarber", server.UpdateBarber)
	barberRoutes.POST("/token/refresh_access", server.ReNewAccessToken)

	barberShopRoutes := router.Group("/barberShop").Use(server.AddMiddleWare(server.tokenMaker))
	barberShopRoutes.POST("/newBarberShop", server.NewBarberShop)
	barberShopRoutes.GET("/:id", server.GetCodeBarberShop)

	// router.POST("/image", server.uploadImage)
	// router.Static("/assets", "./assets/upload")
	// router.POST("/users", server.createUser)

	// authRoutes.POST("/store", server.NewStore)
	// authRoutes.GET("/store/:id", server.GetStore)
	// authRoutes.GET("/store", server.GetListStore)
	// authRoutes.PUT("/store", server.UpdateStore)
	// authRoutes.DELETE("/store/:id", server.DeleteStore)

	// authRoutes.GET("/service_category/:id_store", server.GetListServiceCategorywithStore)
	// authRoutes.POST("/service_category", server.CreateServiceCategory)
	// authRoutes.DELETE("/service_category/:id", server.DeleteServiceCategory)

	// authRoutes.GET("/service/:id", server.GetListServicewithCategory)
	// authRoutes.POST("/service", server.CreateService)
	// authRoutes.PUT("/service", server.UpdateService)
	// authRoutes.DELETE("/service/:id", server.DeleteService)

	// router.GET("/sdkNotification", sdkFirebaseAdmin) // todo

	server.Router = router
}
