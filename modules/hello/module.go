package hello

import (
	"app/modules/hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	initRouter()
}
func initRouter() {
	// Object RESTFUL
	beego.Router("/hello/objects", &controllers.ObjectController{}, "get:Index")
	beego.Router("/hello/objects", &controllers.ObjectController{}, "post:Store")
	beego.Router("/hello/objects/:id", &controllers.ObjectController{}, "get:Show")
	beego.Router("/hello/objects/:id", &controllers.ObjectController{}, "put:Update")
	beego.Router("/hello/objects/:id", &controllers.ObjectController{}, "delete:Destroy")
}
