package middlewares

import (
	"net/http"

	"web-service-gin/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireAuthCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the required cookie exists
		cookie, err := c.Request.Cookie("auth_token")
		if err != nil || cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Continue to the next middleware or route handler
		c.Next()
	}
}
