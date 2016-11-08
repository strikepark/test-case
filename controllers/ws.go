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

	fmt.Println("Upgrade: " + string(this.Ctx.Request.FormValue("Upgrade")))

	if err != nil {
		fmt.Println(this.Ctx.Request)
		fmt.Println(err)
	} else {
		fmt.Println("Not err")
	}
}