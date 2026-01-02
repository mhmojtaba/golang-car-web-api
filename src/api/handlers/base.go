package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhmojtaba/golang-car-web-api/api/helper"
	"github.com/mhmojtaba/golang-car-web-api/config"
	"github.com/mhmojtaba/golang-car-web-api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err, "Invalid request"))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err, "Failed to create resource"))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success, "Resource created successfully"))
}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err, "Invalid request"))
		return
	}

	res, err := caller(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err, "Failed to update resource"))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success, "Resource updated successfully"))
}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError, "Resource not found"))
		return
	}

	err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err, "Failed to delete resource"))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(nil, true, helper.Success, "Resource deleted successfully"))
}

func GetById[To any](c *gin.Context, caller func(c context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.ValidationError, "Resource not found"))
		return
	}

	res, err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err, "Failed to get resource"))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success, "Resource retrieved successfully"))
}

func GetByFilter[Ti any, To any](c *gin.Context, caller func(c context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err, " Invalid request"))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err, "Failed to get resources"))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success, "Resources retrieved successfully"))
}
