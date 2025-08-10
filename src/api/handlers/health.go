package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler
type HealthHandlers struct{}

// create new handler
func NewHealthHandler() *HealthHandlers {
	return &HealthHandlers{}
}

// health godoc
// @Summary Health handler
// @Description Health handler function
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "failed"
// @Router /v1/health/ [get]
func (h *HealthHandlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
