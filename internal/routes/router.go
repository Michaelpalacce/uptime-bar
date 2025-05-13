package routes

import (
	"os"

	"github.com/Michaelpalacce/uptime-bar/internal/options"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Args   *options.RunOptions
	Engine *gin.Engine
}

func (r *Router) Run() error {
	if err := os.Setenv("PORT", r.Args.RouterOptions.Port); err != nil {
		return err
	}

	return r.Engine.Run(r.Args.RouterOptions.Address)
}
