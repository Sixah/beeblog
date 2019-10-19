package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true                 // 导航栏首页 高亮标志位
	c.TplName = "home.html"                 // 首页页面模板渲染
	c.Data["IsLogin"] = checkAccount(c.Ctx) // 检查是否登录
	cate := c.Input().Get("cate")
	label := c.Input().Get("label")
	topics, err := models.GetAllTopic(cate,label,true) // 查询文章列表
	if err != nil {
		beego.Error(err)

	} else {
		c.Data["Topics"] = topics // 文章列表数据
	}

	categories,err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
}
