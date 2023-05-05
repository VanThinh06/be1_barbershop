package api

import (
	db "barbershop/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

type IdEmployee struct {
	ID string `uri:"id" binding:"required"`
}

func NewServer(queries *db.Queries) *Server {
	server := &Server{
		queries: queries,
	}
	router := gin.Default()
	router.DELETE("/employee/:id", server.deleteEmployee)

	router.POST("/store", server.NewStore)

	server.router = router
	return server
}

func (sever *Server) deleteEmployee(ctx *gin.Context) {
	var req IdEmployee
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := sever.queries.DeleteEmployee(ctx, uuid.MustParse(req.ID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "success")

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
