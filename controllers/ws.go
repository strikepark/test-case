package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	//"net/http"
	//"fmt"
	"time"
)

type WsController struct {
	beego.Controller
}

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var upgrader = websocket.Upgrader{}

func (this *WsController) WsHandle() {
	var conn, _ = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	go func(conn *websocket.Conn) {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				conn.Close()
			}
		}
	}(conn)

	go func(conn *websocket.Conn) {
		ch := time.Tick(5 * time.Second)

		for range ch {
			conn.WriteJSON(myStruct{
				Username:  "mvansickle",
				FirstName: "Michael",
				LastName:  "Van Sickle",
			})
		}
	}(conn)
}