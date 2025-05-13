package handlers

import (
	"github.com/gin-gonic/gin"
)

// StatusHandler is the handler for uptime statuses
type StatusHandler struct{}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

// GetAll will retrieve the status of all items
func (h *StatusHandler) All(c *gin.Context) {
}
