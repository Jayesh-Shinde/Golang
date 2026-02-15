package middlewares

import (
	"net/http"

	"example.org/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticateRequest(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	id, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized.." + err.Error()})
		return
	}

	context.Set("userId", id)
	context.Next()
}
