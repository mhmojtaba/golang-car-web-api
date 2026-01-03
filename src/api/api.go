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
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {

	r := gin.New()

	// set validators for tags
	RegisterValidators()
	// set middlewares
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler) /* middlewares.NewTestMiddleware(),, middlewares.Limiter()*/)

	RegisterRouter(r, cfg)
	RegisterSwagger(r, cfg)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		logger.Error(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRouter(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test := v1.Group("/test")
		users := v1.Group("/users")
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"Admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"Admin"}))
		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"Admin", "default"}))

		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"Admin"}))
		properties := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"Admin", "default"}))

		companies := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		colors := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		years := v1.Group("/years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		carTypes := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModels := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelYears := v1.Group("/car-model-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelPriceHistories := v1.Group("/car-model-price-histories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelImages := v1.Group("/car-model-images", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelProperties := v1.Group("/car-model-properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelComments := v1.Group("/car-model-comments", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin", "default"}))

		routers.User(users, cfg)
		routers.Country(countries, cfg)
		routers.City(cities, cfg)
		routers.File(files, cfg)

		routers.Company(companies, cfg)
		routers.Color(colors, cfg)
		routers.Year(years, cfg)

		routers.CarType(carTypes, cfg)
		routers.Gearbox(gearboxes, cfg)
		routers.CarModel(carModels, cfg)
		routers.CarModelColor(carModelColors, cfg)
		routers.CarModelYear(carModelYears, cfg)
		routers.CarModelPriceHistory(carModelPriceHistories, cfg)
		routers.CarModelImage(carModelImages, cfg)
		routers.CarModelProperty(carModelProperties, cfg)
		routers.CarModelComment(carModelComments, cfg)

		routers.PropertyCategory(propertyCategories, cfg)
		routers.Property(properties, cfg)

		routers.TestRouter(test)
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validation.ValidateMobile, true)
		if err != nil {
			logger.Error(logging.Validation, logging.Startup, err.Error(), nil)
		}
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
