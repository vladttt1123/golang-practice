package middlewares

import (
	"eventBooking/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token must not be empty"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
