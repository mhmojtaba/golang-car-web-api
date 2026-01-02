package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/handlers"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCountryHandler(cfg)

	r.POST("/", handler.CreateCountry)
	r.PUT("/:countryId", handler.UpdateCountry)
	r.DELETE("/:countryId", handler.DeleteCountry)
	r.GET("/:countryId", handler.GetCountry)
	r.POST("/get-by-filter", handler.GetCountriesByFilter)
}

func City(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCityHandler(cfg)

	r.POST("/", handler.Create)
	r.PUT("/:cityId", handler.Update)
	r.DELETE("/:cityId", handler.Delete)
	r.GET("/:cityId", handler.GetById)
	r.POST("/get-by-filter", handler.GetByFilter)
}

func File(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewFileHandler(cfg)

	r.POST("/", handler.Create)
	r.PUT("/:fileId", handler.Update)
	r.DELETE("/:fileId", handler.Delete)
	r.GET("/:fileId", handler.GetById)
	r.POST("/get-by-filter", handler.GetByFilter)
}
