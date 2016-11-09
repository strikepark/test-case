package controllers

import (
	"github.com/astaxie/beego"


	//"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/pkg/errors"
	//"io"
	//"io"
	"log"
)

type WsController struct {
	beego.Controller
}

type updateMessege struct {
	Update bool
	History string
}

var upgrader = websocket.Upgrader{}

func validateMessage(data []byte) (msg updateMessege, err error) {
	if err = json.Unmarshal(data, &msg); err != nil {
		return msg, errors.Wrap(err, "Unmarshaling message")
	}

	if msg.Update == false && msg.History == "" {
		return msg, errors.New("Message has no Handle or Text")
	}

	return msg, nil
}

var ws *websocket.Conn
var err error

func (this *WsController) WsHandle() {
	ws, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
//
//func WsSend(ws *websocket.Conn) error {
//	defer ws.Close()
//
//	for {
//		if err := ws.WriteMessage(websocket.TextMessage, []byte("Update")); err != nil {
//			fmt.Println("Websocket close(write messege error): ", err)
//			ws.Close()
//			break
//		} else {
//			fmt.Println("Websocket messege send")
//		}
//
//		ws.Close()
//	}
//
//	return nil
//}