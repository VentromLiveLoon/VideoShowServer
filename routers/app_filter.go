package routers

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	logs.Info("Init filter")
	// Index page filter
	web.InsertFilter("/", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
	web.InsertFilter("/index", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
	web.InsertFilter("/video/*", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
	web.InsertFilter("/video/*", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
	web.InsertFilter("/collect/", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
	web.InsertFilter("/collect/*", web.BeforeExec, func(ctx *context.Context) {
		username := ctx.GetCookie("username")
		s, err := ctx.Session()
		if err != nil {
			panic(err.Error())
		}
		// Get user model from session
		user := s.Get(ctx.Request.Context(), username)
		if user == nil {
			ctx.Redirect(302, "/login")
		}
	})
}
