package middleware

import (
	"net/http"
	"strings"

	"github.com/Somaycon/api-go/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeadle := ctx.GetHeader("Authorization")
		if authHeadle == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Autorization header required"})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeadle, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString := parts[1]

		token, err := handler.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired Token,", "details": err.Error()})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			ctx.Abort()
			return
		}

		userId, ok := claims["user_id"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", uint(userId))
		ctx.Next()
	}
}
