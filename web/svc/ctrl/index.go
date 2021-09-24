package ctrl

import (
	"ecomb-go-demo/web/cfg"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	m := make(gin.H)

	m["env"] = cfg.E.Env
	m["status"] = "OK"

	Return(c, m)
}