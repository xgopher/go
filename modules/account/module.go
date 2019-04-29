package account

import (
	"app/modules/account/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 路由初始化
	beego.Router("/account/users", &controllers.UserController{}, "get:Index")
	beego.Router("/account/users", &controllers.UserController{}, "post:Store")
	beego.Router("/account/users/:id([0-9]+)", &controllers.UserController{}, "get:Show")
	beego.Router("/account/users/:id([0-9]+)", &controllers.UserController{}, "put:Update")
	beego.Router("/account/users/:id([0-9]+)", &controllers.UserController{}, "delete:Destroy")
}
