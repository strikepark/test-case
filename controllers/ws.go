package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
)

type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

func (this *WsController) WsHandle() {
	fmt.Println("http://" + this.Ctx.Request.Host)

	if this.Ctx.Request.Header.Get("Origin") != ("http://" + this.Ctx.Request.Host) {
		fmt.Println(this.Ctx.ResponseWriter, "Origin not allowed", 403)
		return
	}

	_, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.Println("err", err, m)
		return
	}
}