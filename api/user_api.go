package api

import (
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type createUsersParams struct {
	Username string      `json:"username" binding:"required,alphanum"`
	FullName null.String `json:"full_name"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required,min=6"`
	Image    null.String `json:"image"`
	Role     null.String `json:"role"`
}
type createUserResponse struct {
	Username string      `json:"username"`
	FullName null.String `json:"full_name"`
	Email    string      `json:"email"`

	PasswordChangedAt time.Time   `json:"password_changed_at"`
	CreatedAt         time.Time   `json:"created_at"`
	Image             null.String `json:"image"`
	Role              null.String `json:"role"`
}

func (server *Server) createUser(ctx *gin.Context) {

	var req createUsersParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUsersParams{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Image:          req.Image,
		Role:           req.Role,
	}

	user, err := server.queries.CreateUsers(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {

			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	userResponse := createUserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		Image:             user.Image,
		Role:              user.Image,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, userResponse)
}
