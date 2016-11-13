package controllers

import (
	"golang.org/x/net/websocket"
	"fmt"
)

type JSONCode struct {
	Code int64 `json:"code"`
}

var wsList = make(map[int64] *websocket.Conn)

func WsHandler(ws *websocket.Conn) {
	for {
		var data JSONCode
		websocket.JSON.Receive(ws, &data)

		wsList[data.Code] = ws
	}
}

type WsMessege struct {
	Msg string
	UpdateFlag bool
    History string
}

func WsSend(ws *websocket.Conn, data WsMessege) {
    fmt.Println()

	if err := websocket.JSON.Send(ws, data); err != nil {
		fmt.Printf("%s", err) // "use of closed network connection"
	}
}
