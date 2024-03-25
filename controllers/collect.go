package controllers

import (
	"VideoShow/models"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type CollectController struct {
	web.Controller
}

func (c *CollectController) CollectScrollNext() {
	username := c.Ctx.GetCookie("username")
	user := c.GetSession(username).(*models.User)
	if user == nil {
		c.Redirect("/login", 302)
		return
	}

	id, _ := c.GetUint64("collectid")
	logs.Info(id)
	collect := new(models.Collect)
	collect.Userid = user.Id
	collect.Id = id

	isend := collect.SelectNextByUserIdAndId()
	// if not found
	if isend != "" {
		logs.Info(isend)
		collect.Id = 0
		collect.SelectNextByUserIdAndId()
	}

	vp := new(models.VideoPath)
	vp.Id = collect.Videoid
	s := vp.Select()
	if s != "" {
		c.Ctx.WriteString(s)
		return
	}

	data := fmt.Sprint("title=", vp.Name, "&m3u8url=", models.ZlMediaKitDataPath, vp.Name, "/", models.VideoType, "&collectid=", collect.Id, "&id=", vp.Id)
	logs.Info(data)
	c.Ctx.WriteString(data)

}
func (c *CollectController) CollectCancel() {
	username := c.Ctx.GetCookie("username")
	user := c.GetSession(username).(*models.User)
	if user == nil {
		c.Redirect("/login", 302)
		return
	}
	id, _ := c.GetUint64("collectid")
	collect := new(models.Collect)
	collect.Userid = user.Id
	collect.Videoid = id

	// check is already collect
	check := collect.DeleteByUserIdAndVideoid()
	if check != "" {
		c.Ctx.WriteString("Already canceled!!")
		return
	}
	c.Ctx.WriteString("ok")
}
func (c *CollectController) Collect() {
	username := c.Ctx.GetCookie("username")
	user := c.GetSession(username).(*models.User)
	if user == nil {
		c.Redirect("/login", 302)
		return
	}
	id, _ := c.GetUint64("collectid")
	collect := new(models.Collect)
	collect.Userid = user.Id
	collect.Videoid = id

	// check is already collect
	check := collect.SelectByUseridAndVideoid()
	if check == "" {
		c.Ctx.WriteString("Already collected!!")
		return
	}

	collecterr := collect.Insert()
	if collecterr != "" {
		logs.Error(collecterr)
	}
	c.Ctx.WriteString("ok")
}
