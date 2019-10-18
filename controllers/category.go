package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		fmt.Println("33333333333333333",name)
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category",301)
		return
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.html"
	//var err error
	data,err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = data
	fmt.Println("6666666666",data)
}

func (this *CategoryController) Post() {

}
