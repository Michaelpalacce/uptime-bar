package routes

import (
	"fmt"

	"github.com/Michaelpalacce/uptime-bar/internal/options"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Args   *options.RunOptions
	Engine *gin.Engine
}

func (r *Router) Run() error {
	return r.Engine.Run(fmt.Sprintf("%s:%s", r.Args.RouterOptions.Address, r.Args.RouterOptions.Port))
}
