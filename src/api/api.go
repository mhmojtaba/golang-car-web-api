package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mhmojtaba/golang-car-web-api/api/middlewares"
	"github.com/mhmojtaba/golang-car-web-api/api/routers"
	"github.com/mhmojtaba/golang-car-web-api/api/validation"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	// set validators for tags
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.ValidateMobile)
	}

	// set middlewares
	r.Use(gin.Logger(), gin.Recovery(), middlewares.NewTestMiddleware(), middlewares.Limiter())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("/test")

		routers.TestRouter(test)
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
