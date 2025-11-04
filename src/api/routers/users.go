package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/handlers"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

func User(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUsersHandler(cfg)

	r.POST("/send-otp/" /*middlewares.OtpLimiter(cfg),*/, handler.SendOtp)
	r.POST("/login-by-mobile/", handler.RegisterLoginByMobileNumber)
	r.POST("/register-by-username/", handler.RegisterByUsername)
	r.POST("/login-by-username/", handler.LoginByUsername)
}
