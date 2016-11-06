package routers

import (
	"github.com/astaxie/beego"
	"planadotest/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/hello-world", &controllers.MainController{}, "get:HelloWorld")

	ns := beego.NewNamespace("/api",
		beego.NSRouter("/order/:code", &controllers.OrderController{}, "get:NewOrder"),
		//beego.NSRouter("/orders", &controllers.OrderController{}, "post:NewOrder"),
		beego.NSRouter("/orders", &controllers.OrderController{}, "get:GetOrders"),
		beego.NSRouter("/orders/:code", &controllers.OrderController{}, "get:GetOrder"),
	)

	beego.AddNamespace(ns)
}