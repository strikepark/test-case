package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"

	"io"
)

type WsController struct {
	beego.Controller
}

func (this *WsController) WsHandle(ws *websocket.Conn) {
	io.Copy(ws, ws)
}