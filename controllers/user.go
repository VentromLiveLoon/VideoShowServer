package controllers

import (
	"VideoShow/models"

	"github.com/beego/beego/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	// Query the database to check the user is exists
	logs.Info(username)
	logs.Info(password)
	user := new(models.User)
	user.Name = username
	s := user.Select()
	if s != "" {
		u.Ctx.WriteString("<h1>Login failed ,check you username is correct!!</h1>")
		return
	}
	if password == user.Password {
		// SetCookie
		u.Ctx.SetCookie("username", user.Name, 2592000)
		// SetSession
		u.SetSession(user.Name, user)
		// Login success
		u.Redirect("/index", 302)
	} else {
		u.Ctx.WriteString("<h1>Login failed ,check you username or password is correct!!</h1>")
	}
}

func (u *UserController) Register() {
	username := u.GetString("username")
	password := u.GetString("password")
	comfirm_password := u.GetString("comfirm_password")
	// Check
	if password != comfirm_password {
		u.Ctx.WriteString("<h1> Password is not equal comfirm_password!! </h1>")
	}
	// Save the message to database;
	logs.Info(username)
	logs.Info(password)
	logs.Info(comfirm_password)
	user := models.User{}
	user.Name = username
	user.Password = password
	s := user.Insert()

	if s == "" {
		u.TplName = "login.html"
	} else {
		u.Ctx.WriteString("<h1>Register failed </h1>" + "<p>" + s + "</p>")
	}
}
