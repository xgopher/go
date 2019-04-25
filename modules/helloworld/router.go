package helloworld

import (
	"app/modules/helloworld/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// Object RESTFUL
	beego.Router("/helloworld/objects", &controllers.ObjectController{}, "get:Index")
	beego.Router("/helloworld/objects", &controllers.ObjectController{}, "post:Store")
	beego.Router("/helloworld/objects/:id", &controllers.ObjectController{}, "get:Show")
	beego.Router("/helloworld/objects/:id", &controllers.ObjectController{}, "put:Update")
	beego.Router("/helloworld/objects/:id", &controllers.ObjectController{}, "delete:Destroy")
}
