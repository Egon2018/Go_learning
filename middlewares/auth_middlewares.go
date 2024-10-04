package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exchangeapp/utils"
)

func Authmiddlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error1": "missing auth header"})
			ctx.Abort()
			return
		}
		username, err := utils.ParseJWt(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error1": err})
			ctx.Abort()
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
