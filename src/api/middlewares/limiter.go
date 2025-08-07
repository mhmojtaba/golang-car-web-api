package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}
