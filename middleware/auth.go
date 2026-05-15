package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// APIKeyAuth is a middleware that checks the X-API-Key header against the API_KEY env variable
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		expectedAPIKey := os.Getenv("API_KEY")
		if expectedAPIKey == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "API_KEY not configured on server"})
			return
		}

		providedAPIKey := c.GetHeader("X-API-Key")

		if providedAPIKey != expectedAPIKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or missing X-API-Key header"})
			return
		}

		c.Next()
	}
}
