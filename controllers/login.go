package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true" // 是否退出登录
	if isExit {                                  // 退出登录
		// 删除cookie
		this.Ctx.SetCookie("account", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		this.Redirect("/", 301) // 重定向到首页
		return
	}
	this.TplName = "login.html" // 登录页面模板渲染
}

func (this *LoginController) Post() {
	account := this.Input().Get("account")                                                              // 获取用户帐号
	password := this.Input().Get("password")                                                            // 获取用户密码
	autoLogin := this.Input().Get("autoLogin") == "on"                                                  // 是否勾选自动登录
	if beego.AppConfig.String("account") == account && beego.AppConfig.String("password") == password { // 帐号密码验证
		maxAge := 0
		if autoLogin { // 自动登录 设置cookie时间超长
			maxAge = 1<<31 - 1
		}
		// 设置cookie
		this.Ctx.SetCookie("account", account, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
	}
	this.Redirect("/", 301) // 重定向到首页
	return
}

// 检查是否登录
func checkAccount(ctx *context.Context) bool {
	// 查询cookie
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}
	account := ck.Value
	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	return beego.AppConfig.String("account") == account && beego.AppConfig.String("password") == password // 帐号密码验证
}
