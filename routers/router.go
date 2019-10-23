package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply",&controllers.ReplyController{})
	beego.Router("/reply/add",&controllers.ReplyController{},"post:Add")
	beego.Router("/reply/del",&controllers.ReplyController{},"get:Delete")
	// 附件目录作为静态文件
	//beego.SetStaticPath("/attachment","attachment")
	// 附件目录作为一个单独的控制器来处理
	beego.Router("/attachment/:all",&controllers.AttachmentController{})
}
