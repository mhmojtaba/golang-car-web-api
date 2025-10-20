package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/handlers"
	"github.com/mhmojtaba/golang-car-web-api/api/middlewares"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUsersHandler(cfg)

	r.POST("/send-otp/", middlewares.OtpLimiter(cfg), handler.SendOtp)
}
