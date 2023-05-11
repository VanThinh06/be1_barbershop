package api

import (
	db "barbershop/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	queries db.StoreMain
	router  *gin.Engine
}

func NewServer(queries db.StoreMain) *Server {
	server := &Server{
		queries: queries,
	}
	router := gin.Default()
	router.POST("/barber", server.newBarber)
	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
