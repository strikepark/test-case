package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"io"
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
	if this.Ctx.Request.Header.Get("Origin") != ("http://" + this.Ctx.Request.Host) {
		http.Error(this.Ctx.ResponseWriter, "Origin not allowed", 403)
		return
	}
	
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		log.Println("err", err, m)
		return
	}

	for {
		mt, data, err := ws.ReadMessage()
		if err != nil {
			log.Println("Websocket closed!", err)
			break
		}

		switch mt {
			case websocket.TextMessage:
				msg, err := validateMessage(data)
				if err != nil {
					log.Println("Websocket closed!", err)
					break
				}
			rw.publish(data)
			default:
			l.Warning("Unknown Message!")
		}
	}
}