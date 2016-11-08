package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
)

type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (this *WsController) WsHandle() {
	var conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	//go func(conn *websocket.Conn) {
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			if err = conn.WriteMessage(mType, msg); err != nil {
				log.Println(err)
				return
			}

			conn.WriteMessage(mType, msg)
		}
	//}(conn)
}