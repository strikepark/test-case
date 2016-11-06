package controllers

import (
	"github.com/astaxie/beego"

	m "planadotest/models"
	"fmt"
	"encoding/json"
	"bytes"
)

type OrderController struct {
	beego.Controller
}

func (this *OrderController) CreateOrder() {
	order := m.Order{}

	n := bytes.IndexByte(this.Ctx.Input.RequestBody, 0)
	s := string(this.Ctx.Input.RequestBody[:n-1])
	fmt.Println("Order: ", s)

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &order); err != nil {
		this.Abort("500")
	}

	result, err := order.CreateOrder()
	if err != nil {
		this.Abort("500")
	} else {
		this.Data["json"] = result
	}

	this.ServeJSON()
}

func (this *OrderController) GetOrders() {

	orders := m.GetOrders()

	this.Data["json"] = orders
	this.ServeJSON()
}

//func (this *OrderController) GetOrder() {
//	id := this.Ctx.Input.Param(":id")
//	if id == "" {
//		this.Abort("403")
//	}
//
//	order, err := m.GetOrder(id)
//	if err != nil {
//		this.Abort("500")
//	} else {
//		this.Data["json"] = order
//	}
//
//	this.ServeJSON()
//}
//
//func (this *OrderController) UpdateOrder() {
//	id := this.Ctx.Input.Param(":id")
//	if id == "" {
//		this.Abort("403")
//	}
//
//	order := m.Order{Id: id}
//
//	json.Unmarshal(this.Ctx.Input.RequestBody, &order)
//
//	result, err := m.Update(order)
//	if err != nil {
//		this.Abort("500")
//	} else {
//		this.Data["json"] = result
//	}
//
//	this.ServeJSON()
//}
//
//func (this *OrderController) GetCostumerOrders() {
//	code := this.Ctx.Input.Param(":code")
//	if code == "" {
//		this.Abort("403")
//	}
//
//	orders, err := m.GetOrder(code)
//	if err != nil {
//		this.Abort("500")
//	} else {
//		this.Data["json"] = orders
//	}
//
//	this.ServeJSON()
//}


//import (
//	"github.com/astaxie/beego"
//
//	"encoding/json"
//	//"strconv"
//
//	m "planadotest/models"
//	"strconv"
//)
//
//// Order controller
//type OrderController struct {
//	beego.Controller
//}
//
//func (this *OrderController) NewOrder() {
//	code := this.Ctx.Input.Param(":code")
//
//	var intCode int
//
//	intCode, _ = strconv.Atoi(code)
//	order := m.Order{intCode, 99999999, "str1", "str2", 9999999999, "status1"}
//	err := order.NewOrder()
//
//	if err == nil {
//		str, _ := json.Marshal(order)
//		this.Data["json"] = string(str)
//	} else {
//		str, _ := json.Marshal(err)
//		this.Data["json"] = string(str)
//	}
//
//	this.ServeJSON()
//}
//
//func (this *OrderController) GetOrders() {
//	str, _ := json.Marshal(m.GetOrders())
//
//	//if str == "" {
//	//	this.Abort("403")
//	//}
//	//
//	//order, err := m.GetOrder(str)
//	//
//	//if err != nil {
//	//	this.Abort("500")
//	//} else {
//	//	this.Data["json"] = order
//	//}
//
//	this.Data["json"] = string(str)
//	this.ServeJSON()
//}
//
//func (this *OrderController) GetOrder() {
//	code := this.Ctx.Input.Param(":code")
//	//jsonCode, _ := json.Marshal(code)
//
//	str, _ := m.GetOrder(code)
//	strJson, _ := json.Marshal(str)
//	this.Data["json"] = string(strJson)
//	this.ServeJSON()
//}