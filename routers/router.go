package routers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
	c "planadotest/controllers"
	"io"
	"net/http"
)

func init() {
	//beego.Router("/ws", websocket.Handler(WsHandle))

	http.Handle("/echo", websocket.Handler(WsHandle))

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

func WsHandle(ws *websocket.Conn) {
	io.Copy(ws, ws)
}