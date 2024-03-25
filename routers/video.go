package routers

import (
	"VideoShow/controllers"
	"VideoShow/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	logs.Info("Init video router")
	web.CtrlPost("/video/scroll/next", (*controllers.VideoController).ScrollNext)
	// /record/alsdfjowjfelsajf.mp4/hls_delay.m3u8
	logs.Info(models.ZlMediaKitDataPath)
	web.CtrlGet(models.ZlMediaKitDataPath+"*", (*controllers.VideoController).GetM3u8Data)
}
