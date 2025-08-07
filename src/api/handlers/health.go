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

// methods
func (h *HealthHandlers) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
