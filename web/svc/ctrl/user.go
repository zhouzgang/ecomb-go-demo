package ctrl

import (
	"ecomb-go-demo/web/model/user"
	"ecomb-go-demo/web/svc/constant"
	"github.com/gin-gonic/gin"
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
	res := di.
}