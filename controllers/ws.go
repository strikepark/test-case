package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	//"net/http"
	"fmt"
)

type WsController struct {
	beego.Controller
}

func (this *WsController) WsHandle() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		fmt.Println(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		//http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(string(p))
	}
}