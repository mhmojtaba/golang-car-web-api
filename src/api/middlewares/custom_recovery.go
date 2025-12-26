package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, int(helper.CustomRecovery), err, err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithError(nil, false, int(helper.CustomRecovery), nil, "Internal Server Error")
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
