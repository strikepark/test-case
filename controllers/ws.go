package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
	
	"io"
)

type WsController struct {
	beego.Controller
}

func WsHandle(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
