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
