package routers

import (
	"VideoShow/controllers"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	logs.Info("Init page router")
	// Index page
	web.CtrlGet("/", (*controllers.MainController).Index)
	logs.Info("Init router")
	web.CtrlGet("/index", (*controllers.MainController).Index)
	// Login page
	web.CtrlGet("/login", (*controllers.MainController).Login)
	// Register page
	web.CtrlGet("/register", (*controllers.MainController).Register)
	web.CtrlGet("/collect", (*controllers.MainController).Collect)
}
