package routes

import (
	"github.com/Michaelpalacce/uptime-bar/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the application routes.
func SetupRouter(
	statusHandler handlers.StatusHandler,
) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")

	// Uptime Routes
	uptimeRoutes := v1.Group("/uptime")
	{
		uptimeRoutes.GET("/", statusHandler.All)
	}

	return r
}
