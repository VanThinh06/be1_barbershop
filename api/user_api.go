package api

// import (
// 	db "barbershop/db/sqlc"
// 	"barbershop/util"
// 	"database/sql"
// 	"errors"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/google/uuid"
// 	"github.com/lib/pq"
// 	"gopkg.in/guregu/null.v4"
// )

// type createUsersParams struct {
// 	Username   string      `json:"username" binding:"required,alphanum"`
// 	FullName   string      `json:"full_name"`
// 	Email      string      `json:"email" binding:"required,email"`
// 	Password   string      `json:"password" binding:"required,min=6"`
// 	Image      null.String `json:"image"`
// 	Role       null.String `json:"role"`
// 	Fcm_Device string      `json:"fcm" binding:"required"`
// }
// type userResponse struct {
// 	Username          string      `json:"username"`
// 	FullName          string      `json:"full_name"`
// 	Email             string      `json:"email"`
// 	PasswordChangedAt time.Time   `json:"password_changed_at"`
// 	CreatedAt         time.Time   `json:"created_at"`
// 	Image             null.String `json:"image"`
// 	Role              null.String `json:"role"`
// }

// func newUserResponse(user db.User) userResponse {
// 	return userResponse{
// 		Username:          user.Username,
// 		FullName:          user.FullName,
// 		Email:             user.Email,
// 		PasswordChangedAt: user.PasswordChangedAt,
// 		CreatedAt:         user.CreatedAt,
// 	}
// }

// func _msgForUserParams(tag string) string {
// 	switch tag {
// 	case "Email":
// 		return "Invalid email"
// 	case "Password":
// 		return "Invalid fmc"
// 	}
// 	return ""
// }
// func (server *Server) createUser(ctx *gin.Context) {
// 	var req createUsersParams
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		var ve validator.ValidationErrors
// 		if errors.As(err, &ve) {
// 			for _, fe := range ve {
// 				if fe.Field() == "email" {
// 					ctx.JSON(http.StatusBadRequest, gin.H{"errors": _msgForUserParams(fe.Field())})
// 					return
// 				}
// 				if fe.Field() == "password" {
// 					ctx.JSON(http.StatusBadRequest, gin.H{"errors": _msgForUserParams(fe.Field())})
// 					return
// 				}
// 			}
// 		}
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	hashedPassword, err := util.HashPassword(req.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}
// 	arg := db.CreateUsersParams{
// 		Username:       req.Username,
// 		FullName:       req.FullName,
// 		Email:          req.Email,
// 		HashedPassword: hashedPassword,
// 		Image:          req.Image,
// 		Role:           req.Role,
// 		FcmDevice:      req.Fcm_Device,
// 	}

// 	user, err := server.queries.CreateUsers(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "unique_violation":
// 				if pqErr.Constraint == "users_email_key" {
// 					ctx.JSON(http.StatusForbidden, "This email has already existed")
// 					return
// 				}
// 				if pqErr.Constraint == "users_pkey" {
// 					ctx.JSON(http.StatusForbidden, "This account has already existed")
// 					return
// 				}
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := newUserResponse(user)

// 	ctx.JSON(http.StatusOK, rsp)
// }

// // login
// type loginUserRequest struct {
// 	Username string `json:"username" binding:"required,alphanum"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type loginUserResponse struct {
// 	SessionId             uuid.UUID    `json:"session_id"`
// 	AccessToken           string       `json:"access_token"`
// 	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
// 	RefreshToken          string       `json:"refresh_token"`
// 	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
// 	User                  userResponse `json:"user"`
// }

// func (server *Server) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	users, err := server.queries.GetUsers(ctx, req.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	err = util.CheckPassword(req.Password, users.HashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
// 		return
// 	}

// 	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
// 		users.Username,
// 		server.config.AccessTokenDuration,
// 	)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
// 		users.Username,
// 		server.config.RefreshTokenDuration,
// 	)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	session, err := server.queries.CreateSession(ctx, db.CreateSessionParams{
// 		ID:           refreshPayload.ID,
// 		Username:     refreshPayload.Username,
// 		RefreshToken: refreshToken,
// 		UserAgent:    ctx.Request.UserAgent(),
// 		ClientIp:     ctx.ClientIP(),
// 		IsBlocked:    false,  // chặn session
// 		ExpiresAt:    refreshPayload.ExpiredAt,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	rsp := loginUserResponse{
// 		SessionId:             session.ID,
// 		AccessToken:           accessToken,
// 		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
// 		RefreshToken:          refreshToken,
// 		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
// User:                  newUserResponse(users),
// 	}
// 	ctx.JSON(http.StatusOK, rsp)
// }
