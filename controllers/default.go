package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}

func (c *MainController) Login() {
	c.TplName = "login.html"
}
func (c *MainController) Register() {
	c.TplName = "register.html"
}
func (c *MainController) Collect() {
	c.TplName = "mycollect.html"
}
