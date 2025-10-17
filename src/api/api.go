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
	"github.com/mhmojtaba/golang-car-web-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	r := gin.New()

	// set validators for tags
	RegisterValidators()
	// set middlewares
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery() /* middlewares.NewTestMiddleware(),*/, middlewares.Limiter())

	RegisterRouter(r)
	RegisterSwagger(r, cfg)

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterRouter(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("/test")

		routers.TestRouter(test)
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.ValidateMobile)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang Car web api"
	docs.SwaggerInfo.Description = "car application web api via golang"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
