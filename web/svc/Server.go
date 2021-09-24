package svc

import (
	"ecomb-go-demo/web/cfg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(cfg.E.Mode)

	app := gin.Default()

	setupRoutes(app)
	app.Run(fmt.Sprintf(":%d", cfg.E.Port))
}
