package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	//"net/http"
	//"fmt"
	//"time"
	"fmt"
)

type WsController struct {
	beego.Controller
}

func (this *WsController) WsHandle() {
	_, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Not err")
	}
}