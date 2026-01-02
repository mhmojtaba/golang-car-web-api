package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/dto"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/services"
)

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	service := services.NewCountryService(cfg)
	return &CountryHandler{
		service: service,
	}
}

// createCountry godoc
// @Summary Create Country
// @Description Create a new country
// @Tags countries
// @Accept json
// @Produce json
// @Param request body dto.CreateUpdateCountryRequest true "Create Country Request"
// @Success 201 {object} helper.BaseHttpResponse{result: dto.CountryResponse} "country response"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.CreateCountry(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to create country"))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(result, true, http.StatusCreated, "Country created successfully"))
}

// updateCountry godoc
// @Summary Update Country
// @Description Update an existing country
// @Tags countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result: dto.CountryResponse} "country response"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/countries/ [put]
// @Security AuthBearer
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	countryId, _ := strconv.Atoi(c.Params.ByName("countryId"))
	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.UpdateCountry(c, uint(countryId), &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to update country"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK, "Country updated successfully"))
}

// deleteCountry godoc
// @Summary Delete Country
// @Description Delete an existing country
// @Tags countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	countryId, _ := strconv.Atoi(c.Params.ByName("countryId"))

	if countryId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, http.StatusNotFound, "Country not found"))
		return
	}

	err := h.service.DeleteCountry(c, uint(countryId))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to delete country"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, http.StatusOK, "Country deleted successfully"))
}

// getCountry godoc
// @Summary Get Country
// @Description Get a country
// @Tags countries
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseHttpResponse{result: dto.CountryResponse} "country response"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetCountry(c *gin.Context) {
	countryId, _ := strconv.Atoi(c.Params.ByName("countryId"))

	if countryId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, http.StatusNotFound, "Country not found"))
		return
	}

	res, err := h.service.GetCountryById(c, uint(countryId))
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to find country"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, http.StatusOK, "Country deleted successfully"))
}

// getCountriesByFilter godoc
// @Summary Get Countries By Filter
// @Description Get countries by filter
// @Tags countries
// @Accept json
// @Produce json
// @Param request body dto.PaginationResultWithFilter true "request"
// @Success 201 {object} helper.BaseHttpResponse{result: dto.Pagination[dto.CountryResponse]} "country response"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Failure 409 {object} helper.BaseHttpResponse "conflict"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetCountriesByFilter(c *gin.Context) {
	req := dto.PaginationResultWithFilter{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.GetCountriesByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to get countries by filter"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK, "Countries fetched successfully"))
}
