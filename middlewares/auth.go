package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-project.com/go-project/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		utils.Unauthorized(context, "Unauthorized")
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		utils.Unauthorized(context, "Unauthorized")
		return
	}

	// set userId in the request context
	context.Set("userId", userId)

	context.Next()
}
