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
	fmt.Println("Upgrade: " + string(this.Ctx.Request.FormValue("Upgrade")))

	_, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	if err != nil {
		fmt.Println(this.Ctx.Request)
		fmt.Println(err)
	} else {
		fmt.Println("Not err")
	}
}