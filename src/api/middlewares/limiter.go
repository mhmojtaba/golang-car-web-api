package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
)

func Limiter() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
			return
		}
		c.Next()
	}
}
