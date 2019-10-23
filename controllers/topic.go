package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"path"
	"strings"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx) // 检测帐号是否登录
	this.TplName = "topic.html"                   // 文章页面模板渲染
	this.Data["IsTopic"] = true // 导航栏文章 高亮标志位
	cate := this.Input().Get("cate")
	label := this.Input().Get("label")
	topics, err := models.GetAllTopic(cate,label,false)      // 查询文章列表
	if err != nil {
		beego.Error(err)

	} else {
		this.Data["Topics"] = topics //文章列表数据
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) { // 如果没有登录
		this.Redirect("/login", 302) // 重定向到登录页面
		return
	}
	// 解析表单
	tType := this.Input().Get("type")      // 文章类型
	title := this.Input().Get("title")     // 文章标题
	content := this.Input().Get("content") // 文章内容
	tid := this.Input().Get("tid")         // 文章id
	labels := this.Input().Get("labels") // 文章标签

	// 获取附件
	_,fh,err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}
	var attachment string
	if fh != nil {
		// 保存附件
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment",path.Join("attachment",attachment))
		// TODO:http保存附件的另一种写法
		if err != nil {
			beego.Error(err)
		}
	}

	// 通过判断文章id是否存在 辨别文章添加和修改操作
	if len(tid) == 0 { // 文章id不存在 添加文章
		err = models.AddTopic(tType, labels,title, content,attachment)
	} else { // 文章id存在 修改文章
		err = models.ModifyTopic(tType,labels, tid, title, content,attachment)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302) // 重定向到文章页面
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) { // 如果没有登录 重定向到登录页面
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "topic_add.html" // 添加文章页面模板渲染
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"                         // 查看文章页面模板渲染
	topic, err := models.GetTopic(this.Ctx.Input.Param("0")) // /topic/view/xxx 注册自动路由，url /topic/view/之后的会解析为参数，存放到map中，key为顺序(0,1,2,3)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic // 文章数据
	this.Data["Labels"] = strings.Split(topic.Labels," ")
	tid := this.Ctx.Input.Param("0")
	this.Data["Tid"] = tid // 文章id
	replies,err := models.GetAllReplies(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html" // 修改文章页面模板渲染
	tid := this.Input().Get("tid")     // 文章id
	topic, err := models.GetTopic(tid) // 查询文章信息
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic // 文章数据
	this.Data["Tid"] = tid     // 文章id
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) { // 检测帐号是否登录
		this.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(this.Ctx.Input.Param("0")) // 删除文章
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
