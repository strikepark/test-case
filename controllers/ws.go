package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (this *WsController) WsHandle() {
	conn, _ := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	go func(conn *websocket.Conn) {
		for {
			mType, msg, _ := conn.ReadMessage()

			conn.WriteMessage(mType, msg)
		}
	}(conn)
}