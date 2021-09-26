package ctrl

import (
	"ecomb-go-demo/web/di"
	"ecomb-go-demo/web/model/user"
	"ecomb-go-demo/web/svc/constant"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

//todo 这里这么做的作用
type UserController struct {}

func (ctrl *UserController) GetInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		ReturnError(c, constant.MISSING_PARAM_ERR, "user id can not be empty")
		return
	}

	var v user.User
	res := di.GlobalDI().DB("user").First(&v, id)
	if res.Error != nil {
		log.Error().Msg(res.Error.Error())
		ReturnError(c, constant.DB_ERR, res.Error.Error())
		return
	}

	Return(c, v)
}