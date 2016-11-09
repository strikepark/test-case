package controllers

import (
	"github.com/astaxie/beego"


	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/pkg/errors"
	//"io"
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

func (this *WsController) WsHandle() {
	ws, _ = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	//if err != nil {
	//	fmt.Println("Websocket error(connection): ", err)
	//	return
	//}

	//for {
	//	messageType, data, err := ws.ReadMessage()
	//	if err != nil {
	//		if err == io.EOF {
	//			fmt.Println("Websocket error(closed): ", err)
	//		} else {
	//			fmt.Println("Websocket error(reading messege): ", err)
	//		}
	//
	//		// break
	//	}
	//
	//	if err = ws.WriteMessage(messageType, data); err != nil {
	//		fmt.Println("Websocket error(write messege): ", err)
	//		return
	//	}
	//
	//	switch messageType {
	//		case websocket.TextMessage:
	//			msg, err := validateMessage(data)
	//			if err != nil {
	//				fmt.Println("Websocket error(validation messege): ", err, msg)
	//				break
	//			}
	//			fmt.Println(msg)
	//		default:
	//			fmt.Println("Websocket error: unknown messageType!")
	//	}
	//}
	//
	//ws.WriteMessage(websocket.CloseMessage, []byte{})
}

func WsSend(ws *websocket.Conn) error {
	fmt.Println("Websocket start sending")
	defer ws.Close()

	for {
		if err := ws.WriteMessage(websocket.TextMessage, []byte("Update")); err != nil {
			fmt.Println("Websocket close(write messege error): ", err)
			ws.Close()
			break
		} else {
			fmt.Println("Websocket messege send")
		}

		ws.Close()
		fmt.Println("Websocket close")
	}

	return nil
}