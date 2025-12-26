package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/handlers"
	"github.com/mhmojtaba/golang-car-web-api/config"
)

func Country(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCountryHandler(cfg)

	r.POST("/create/", handler.CreateCountry)
	r.PUT("/update/:countryId", handler.UpdateCountry)
	r.DELETE("/delete/:countryId", handler.DeleteCountry)
	r.GET("/get/:countryId", handler.GetCountry)
}
