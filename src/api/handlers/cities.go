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

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{
		service: services.NewCityService(cfg),
	}
}

// CreateCity godoc
// @Summary Create a City
// @Description Create a City
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body dto.CreateCityRequest true "Create a City"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/ [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to create city"))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(result, true, http.StatusCreated, "city created successfully"))
}

// UpdateCity godoc
// @Summary Update a City
// @Description Update a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCityRequest true "Update a City"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/cities/{id} [put]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	cityId, _ := strconv.Atoi(c.Params.ByName("cityId"))
	req := dto.CreateUpdateCountryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.Update(c, cityId, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to update city"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK, "city updated successfully"))
}

// DeleteCity godoc
// @Summary Delete a City
// @Description Delete a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	cityId, _ := strconv.Atoi(c.Params.ByName("cityId"))

	if cityId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, http.StatusNotFound, "city not found"))
		return
	}

	err := h.service.Delete(c, cityId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to delete city"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, http.StatusOK, "city deleted successfully"))
}

// GetCity godoc
// @Summary Get a City
// @Description Get a City
// @Tags Cities
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	cityId, _ := strconv.Atoi(c.Params.ByName("cityId"))

	if cityId == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, helper.GenerateBaseResponse(nil, false, http.StatusNotFound, "City not found"))
		return
	}

	res, err := h.service.GetByID(c, cityId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to find city"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, http.StatusOK, "successfully"))
}

// GetCities godoc
// @Summary Get Cities
// @Description Get Cities
// @Tags Cities
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "City response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationResultWithFilter{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, http.StatusBadRequest, err, "Invalid request"))
		return
	}

	result, err := h.service.GetByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err), helper.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err, "Failed to get cities by filter"))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(result, true, http.StatusOK, "cities fetched successfully"))
}
