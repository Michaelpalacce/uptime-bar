package run

import (
	"fmt"

	"github.com/Michaelpalacce/uptime-bar/internal/configuration"
	"github.com/Michaelpalacce/uptime-bar/internal/handlers"
	"github.com/Michaelpalacce/uptime-bar/internal/routes"
	"github.com/Michaelpalacce/uptime-bar/internal/services"
	"github.com/gin-gonic/gin"
)

type RunCommand struct{}

func (c *RunCommand) Name() string {
	return "run"
}

func (c *RunCommand) Run() error {
	configuration, err := configuration.LoadConfiguration()
	if err != nil {
		return fmt.Errorf("error while loading configuration. Err was: %w", err)
	}

	statusService := services.NewStatusService(configuration)

	statusHandler := *handlers.NewStatusHandler(statusService)
	router := routes.Router{
		Args:   c.Args(),
		Engine: gin.Default(),
	}

	router.SetupV1(statusHandler)

	return router.Run()
}
