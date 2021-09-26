package svc

import (
	"ecomb-go-demo/web/svc/ctrl"
	"github.com/gin-gonic/gin"
)

func setupRoutes(app *gin.Engine) {
	app.GET("/", ctrl.Health)
	//app.GET("/")

	userGroup := app.Group("/user")
	{
		userCtrl := new(ctrl.UserController)
		userGroup.GET("/info", userCtrl.GetInfo)
	}
}
