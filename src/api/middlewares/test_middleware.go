package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewTestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		api_key := c.GetHeader("api_key")
		if api_key == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
