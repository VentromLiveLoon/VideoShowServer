package controllers

import (
	"VideoShow/models"
	"bytes"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type VideoController struct {
	web.Controller
}

func (v *VideoController) ScrollNext() {
	// /video/scroll/next
	// Get user info from cookie and session
	username := v.Ctx.GetCookie("username")
	user := v.GetSession(username).(*models.User)
	if user == nil {
		v.Redirect("/login", 302)
		return
	}
	// Get video info
	vp := new(models.VideoPath)
	vp.Id = user.Viewposition
	isend := vp.SelectNext()
	if isend != "" {
		vp.Id = 0
		vp.SelectNext()
		if vp.Id == 0 {
			logs.Info("No name in database No videos in zlmedia server")
			v.Ctx.WriteString("No name in database No videos in zlmedia server")
			return
		}
	}
	user.Viewposition = vp.Id
	user.UpdateViewposition()
	// params['title']
	// params['m3u8url']
	// params['id']
	data := fmt.Sprint("title=", vp.Name, "&m3u8url=", models.ZlMediaKitDataPath, vp.Name, "/", models.VideoType, "&id=", vp.Id)
	logs.Info(data)
	v.Ctx.WriteString(data)
}

func (vc *VideoController) GetM3u8Data() {
	// http://localhost:8080/record/alsdfjowjfelsajf.mp4/hls_delay.m3u8
	m3u8path := vc.Ctx.Request.URL.Path
	// Rerequest zlmedia server
	resp, err := http.Get(models.DownloadZlMediaKitMediaServerIp + m3u8path)
	if err != nil {
		logs.Error(err.Error())
	}
	var videodata bytes.Buffer
	_, err2 := videodata.ReadFrom(resp.Body)
	if err2 != nil {
		logs.Error(err2.Error())
	}
	vc.Ctx.Output.Body(videodata.Bytes())
}
