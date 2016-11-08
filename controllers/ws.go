package controllers

import (
	"github.com/astaxie/beego"

	//"net/http"
	//"fmt"
	//"time"
	"fmt"
	"github.com/gorilla/websocket"
)

type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (this *WsController) WsHandle() {
	_, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Not err")
	}
}