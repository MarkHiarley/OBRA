package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Authorization")
		tokenString, err := ExtractTokenFromHeader(authHeader)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token não fornecido",
			})
			ctx.Abort()
			return
		}

		claims, err := ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token inválido ou expirado",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
