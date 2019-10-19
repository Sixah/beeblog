package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx) // 检测是否登录
	op := this.Input().Get("op")                  // 分类操作类型
	switch op {
	case "add": // 添加分类
		name := this.Input().Get("name") // 分类名称
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301) // 重定向到分类页面
		return
	case "del": // 删除分类
		id := this.Input().Get("id") // 分类id
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301) // 重定向到分类页面
		return
	}
	this.Data["IsCategory"] = true // 导航栏分类 高亮标志位
	this.TplName = "category.html" // 渲染分类页面模板
	data, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = data // 分类列表数据
}

func (this *CategoryController) Post() {

}
