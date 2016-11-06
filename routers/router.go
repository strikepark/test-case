package routers

import (
	"github.com/astaxie/beego"
	c "planadotest/controllers"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSRouter("/orders", &c.OrderController{}, "put:CreateOrder"),
		beego.NSRouter("/orders", &c.OrderController{}, "get:GetOrders"),
		//beego.NSRouter("/orders/:id", &с.OrderController{}, "get:GetOrder"),
		//beego.NSRouter("/orders/:id", &с.OrderController{}, "put:UpdateOrder"),

		//beego.NSRouter("/orders/costumer/:code", &с.OrderController{}, "get:GetCostumerOrders"),
	)

	beego.AddNamespace(ns)
}