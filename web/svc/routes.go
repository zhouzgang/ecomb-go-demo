package svc

import (
	"ecomb-go-demo/web/svc/ctrl"
	"github.com/gin-gonic/gin"
)

func setupRoutes(app *gin.Engine) {
	app.GET("/", ctrl.Health)
	//app.GET("/")


}
