package routes

import (
	"github.com/Michaelpalacce/uptime-bar/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupV1 configures /api/v1
func (r *Router) SetupV1(statusHandler handlers.StatusHandler) {
	gin.SetMode(gin.DebugMode)

	v1 := r.Engine.Group("/api/v1")

	// Uptime Routes
	uptimeRoutes := v1.Group("/uptime")
	{
		uptimeRoutes.GET("/", statusHandler.All)
	}
}
