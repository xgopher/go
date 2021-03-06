package main

import (
	_ "app/core" // 导入核心库
	_ "app/routers" // 加载路由表 (路由表经常会更新，独立出来好处: 可单独编译/更新此包 routes.a 就可以)

	// "app/middlewares/auth"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 权限管理 -- todo...
	// beego.InsertFilter("*", beego.BeforeRouter, auth.AuthMiddleware)

	// 跨域请求中间件
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Authorization", "X-Requested-With", "X_Requested_With", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
