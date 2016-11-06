package routers

import (
	"github.com/astaxie/beego"
	c "planadotest/controllers"
)

func init() {
	//ns := beego.NewNamespace("/api",
		//beego.NSRouter("/orders", &controllers.OrderController{}, "post:CreateOrder"),
		//beego.NSRouter("/orders", &controllers.OrderController{}, "get:GetOrders"),
		//beego.NSRouter("/orders/:id", &controllers.OrderController{}, "get:GetOrder"),
		//beego.NSRouter("/orders/:id", &controllers.OrderController{}, "put:UpdateOrder"),
		//
		//beego.NSRouter("/orders/costumer/:code", &controllers.OrderController{}, "get:GetCostumerOrders"),
	//)

	//beego.AddNamespace(ns)

	beego.Router("/api/orders", &c.OrderController{}, "post:CreateOrder")
	beego.Router("/api/orders", &c.OrderController{}, "get:GetOrders")
}