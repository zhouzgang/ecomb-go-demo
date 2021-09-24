package ctrl

import (
	"ecomb-go-demo/web/cfg"
	"ecomb-go-demo/web/svc/constant"
	"ecomb-go-demo/web/svc/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ReturnResult(c *gin.Context, r model.Result) {
	c.JSON(http.StatusOK, r)
}

func Return(c *gin.Context, data interface{}) {
	locale := c.DefaultQuery("locale", cfg.DefaultLocale)

	ReturnResult(c, model.Result{
		Code: constant.SUCCESS,
		Msg:  cfg.I.Get(locale, "errors", constant.SUCCESS),
		Data: data,
	})
}

func ReturnError(c *gin.Context, code string, messages ...string) {
	var msg string
	if len(messages) == 0 {
		locale := c.DefaultQuery("locale", cfg.DefaultLocale)
		msg = cfg.I.Get(locale, "errors", code)
	} else {
		msg = strings.Join(messages, " | ")
	}

	ReturnResult(c, model.Result{
		Code: code,
		Msg: msg,
	})
}
