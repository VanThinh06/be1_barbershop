package api

import (
	"barbershop/token"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (server *Server) AddMiddleWare(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provider")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fileds := strings.Fields(authorizationHeader)
		if len(fileds) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fileds[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fileds[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			err := errors.New("error! An error occurred. Please try again later")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		server.payload = payload
		
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
