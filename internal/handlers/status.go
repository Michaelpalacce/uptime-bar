package handlers

import (
	"github.com/Michaelpalacce/uptime-bar/internal/services"
	"github.com/gin-gonic/gin"
)

// StatusHandler is the handler for uptime statuses
type StatusHandler struct {
	service *services.StatusService

	// preloadedConfig string
}

func NewStatusHandler(service *services.StatusService) *StatusHandler {
	return &StatusHandler{
		service: service,
	}
}

// GetAll will retrieve the status of all items
func (h *StatusHandler) All(c *gin.Context) {
}
