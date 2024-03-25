package routers

import (
	"VideoShow/controllers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	logs.Info("Init user router")
	web.CtrlPost("/user/login", (*controllers.UserController).Login)
	web.CtrlPost("/user/register", (*controllers.UserController).Register)
}
