package middlewares

import (
	"kurbankan/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.ErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token")
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			utils.ErrorResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", "Invalid token payload")
			c.Abort()
			return
		}

		c.Set("user_id", uint(userID))
		c.Set("email", claims["email"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
