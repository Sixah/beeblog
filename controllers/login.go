package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("account","",-1,"/")
		this.Ctx.SetCookie("password","",-1,"/")
		//this.Data["IsLogin"] = false
		this.Redirect("/",301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	account := this.Input().Get("account")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("account") == account && beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31-1
		}
		this.Ctx.SetCookie("account",account,maxAge,"/")
		this.Ctx.SetCookie("password",password,maxAge,"/")
	}
	this.Redirect("/",301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck,err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}
	account := ck.Value
	ck,err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value
	return beego.AppConfig.String("account") == account && beego.AppConfig.String("password") == password
}
