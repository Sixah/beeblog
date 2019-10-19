package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("tid")
	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")
	err := models.AddReply(tid,nickname,content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid,302)
}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")
	rid := this.Input().Get("rid")
	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid,302)
}
