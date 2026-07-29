package middlewares

import (
	"net/http"
	"rest-api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	//checking for incoming request to have the valid token
	//extracting the token
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// we are getting and extracting thre bearer token from the user before verifying
	token = strings.TrimPrefix(token, "Bearer ")

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId",userId)

	context.Next()

}
