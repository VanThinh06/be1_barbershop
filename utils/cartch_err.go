package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func messageErrorParams(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "min":
		return "Password must be more than 6 characters"
	case "email":
		return "Invalid email account"
	case "phone":
		return "Invalid phone account"
	}
	return ""
}

func MessageResponse(err string) gin.H {
	return gin.H{"error": err}
}

var MessageInternalServer = gin.H{"error": "Internal Server Error"}

func CatchErrorParams(err error) map[string]any {
	var validateError validator.ValidationErrors
	if errors.As(err, &validateError) {
		for _, fe := range validateError {
			return gin.H{"error": messageErrorParams(fe.Tag())}
		}
	}
	return nil
}
