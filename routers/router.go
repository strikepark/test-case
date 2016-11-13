package routers

import (
	"github.com/astaxie/beego"
	c "planadotest/controllers"
	"golang.org/x/net/websocket"
)

func init() {
    beego.ErrorController(&c.MainController{})
	beego.Handler("/ws", websocket.Handler(c.WsHandler))

	ns := beego.NewNamespace("/api",
		beego.NSRouter("/orders", &c.OrderController{}, "put:CreateOrder"),
		beego.NSRouter("/orders", &c.OrderController{}, "get:GetOrders"),
		beego.NSRouter("/orders/:id", &c.OrderController{}, "get:GetOrder"),
		beego.NSRouter("/orders/:id", &c.OrderController{}, "put:UpdateOrder"),

		beego.NSRouter("/orders/costumer/:phoneNumber", &c.OrderController{}, "get:GetCostumerOrders"),
		//beego.NSRouter("/ws", &c.WsController{}, "*:WsHandle"),
	)

	beego.AddNamespace(ns)
}
