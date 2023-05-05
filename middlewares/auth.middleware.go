package middlewares

import "github.com/gin-gonic/gin"

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "token is null"})
			// ctx.Abort()
			return
		}

	}
}
