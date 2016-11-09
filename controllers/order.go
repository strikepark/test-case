package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"

	m "planadotest/models"
	"encoding/json"
	"strconv"
	"time"
	"fmt"
	"io"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) CreateOrder() {
	order := m.Order{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order); err != nil {
		this.Abort("500")
	}

	result, err := order.CreateOrder()
	if err != nil {
		fmt.Println(err)
		this.Abort("500")
	} else {
		history := m.History{Code: result.Code, Status: result.Status, Date: time.Now()}

		_, err = history.AddToHistory()

		if err != nil {
			fmt.Println("Error add history")
			fmt.Println(err)
		}

		this.Data["json"] = result
	}

	this.ServeJSON()
}

func (this *OrderController) GetOrders() {
	orders := m.GetOrders()

	this.Data["json"] = orders

	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.ServeJSON()
}

func (this *OrderController) GetOrder() {
	id := this.Ctx.Input.Param(":id")
	if id == "" {
		this.Abort("403")
	}

	order, err := m.GetOrder(id)
	if err != nil {
		this.Abort("500")
	} else {
		this.Data["json"] = order
	}

	this.ServeJSON()
}

//var wsList = make(map[string] *websocket.Conn)

func WsHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	_, err := ws.Read(msg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ws.Config())

	//wsList["id:" + ] = ws
	io.Copy(ws, ws)
}

//type WsMessege struct {
//	Msg string
//	UpdateFlag bool
//}
//
//func WsSend(ws []*websocket.Conn, data WsMessege) {
//	for _, conn := range ws {
//		if err := websocket.JSON.Send(conn, data); err != nil {
//			fmt.Printf("%s", err)   //"use of closed network connection"
//		}
//	}
//}



func (this *OrderController) UpdateOrder() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	idStr := strconv.Itoa(id)
	if idStr == "" {
		this.Abort("403")
	}

	order := m.Order{Id: id}

	json.Unmarshal(this.Ctx.Input.RequestBody, &order)

	result, err := m.UpdateOrder(order)
	if err != nil {
		this.Abort("500")
	} else {
		history := m.History{Code: result.Code, Status: result.Status, Date: time.Now()}

		_, err = history.AddToHistory()

		if err != nil {
			fmt.Println("Error add history")
			fmt.Println(err)
		}

		fmt.Println("Exec webscocket func")

		//wsMsg := WsMessege{Msg: "hello", UpdateFlag: true}
		//
		//go WsSend(wsList, wsMsg)

		this.Data["json"] = result
	}

	this.ServeJSON()
}

func (this *OrderController) GetCostumerOrders() {
	phoneNumber := this.Ctx.Input.Param(":phoneNumber")
	if phoneNumber == "" {
		this.Abort("403")
	}

	result, err := m.GetCostumerOrders(phoneNumber)
	if err != nil {
		this.Abort("500")
	} else {
		this.Data["json"] = result
	}

	this.ServeJSON()
}