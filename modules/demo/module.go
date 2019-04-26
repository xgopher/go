package demo

import (
	"app/modules/demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	initRouter()
}
func initRouter() {
	// Object RESTFUL
	beego.Router("/demo/objects", &controllers.ObjectController{}, "get:Index")
	beego.Router("/demo/objects", &controllers.ObjectController{}, "post:Store")
	beego.Router("/demo/objects/:id", &controllers.ObjectController{}, "get:Show")
	beego.Router("/demo/objects/:id", &controllers.ObjectController{}, "put:Update")
	beego.Router("/demo/objects/:id", &controllers.ObjectController{}, "delete:Destroy")
}
