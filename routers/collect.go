package routers

import (
	"VideoShow/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.CtrlPost("/collect/video", (*controllers.CollectController).Collect)
	web.CtrlPost("/collect/cancel", (*controllers.CollectController).CollectCancel)
	web.CtrlPost("/collect/scroll/next", (*controllers.CollectController).CollectScrollNext)
}
