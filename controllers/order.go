package controllers

import (
	"github.com/astaxie/beego"

	m "planadotest/models"
	"encoding/json"
	"strconv"
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
		this.Data["json"] = result
	}

	this.ServeJSON()
}

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