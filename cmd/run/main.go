package run

import (
	"github.com/Michaelpalacce/uptime-bar/internal/handlers"
	"github.com/Michaelpalacce/uptime-bar/internal/routes"
	"github.com/gin-gonic/gin"
)

type RunCommand struct{}

func (c *RunCommand) Name() string {
	return "run"
}

func (c *RunCommand) Run() error {
	statusHandler := *handlers.NewStatusHandler()
	router := routes.Router{
		Args:   c.Args(),
		Engine: gin.Default(),
	}

	router.SetupV1(statusHandler)

	return router.Run()
}
