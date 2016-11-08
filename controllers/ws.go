package controllers

import (
	"github.com/astaxie/beego"

	//"net/http"
	//"fmt"
	//"time"
	"fmt"
	"github.com/gorilla/websocket"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"io"
)

type WsController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

// message sent to us by the javascript client
type message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

func validateMessage(data []byte) (message, error) {
	var msg message

	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, errors.Wrap(err, "Unmarshaling message")
	}

	if msg.Handle == "" && msg.Text == "" {
		return msg, errors.New("Message has no Handle or Text")
	}

	return msg, nil
}

func (this *WsController) WsHandle() {
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		fmt.Println("err", err, m)
		http.Error(this.Ctx.ResponseWriter, m, http.StatusBadRequest)
		return
	}

	fmt.Println("РАБОТАЕТ")

	if err != nil {
		m := "Unable to upgrade to websockets"
		fmt.Println("err", err, m)
		http.Error(this.Ctx.ResponseWriter, m, http.StatusBadRequest)
		return
	}

	fmt.Println("РАБОТАЕТ")

	for {
		mt, data, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Websocket closed!")
			} else {
				fmt.Println("Error reading websocket message")
			}

			break
		}

		switch mt {
		case websocket.TextMessage:
			msg, err := validateMessage(data)
			if err != nil {
				fmt.Println("msg:", msg, "err:", err)
				break
			}
			fmt.Println(string(data[:]))
		default:
			fmt.Println("Unknown Message!")
		}
	}

	ws.WriteMessage(websocket.CloseMessage, []byte{})
}