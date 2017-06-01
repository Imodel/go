package routers

import (
	"github.com/astaxie/beego"
	"webproject/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/:id:int", &controllers.UserController{}, "get:GetUser")
	beego.Router("/user/", &controllers.UserController{}, "post:AddUser;get:ListUser")
}
