package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	//"strconv"

	m "planadotest/models"
)

// Order controller
type OrderController struct {
	beego.Controller
}

func (this *OrderController) NewOrder() {
	order := m.Order{10, 99999999, "str1", "str2", 9999999999, "status1"}
	err := order.NewOrder()

	if err == nil {
		str, _ := json.Marshal(order)
		this.Data["json"] = string(str)
	} else {
		str, _ := json.Marshal(err)
		this.Data["json"] = string(str)
	}

	this.ServeJSON()
}

func (this *OrderController) GetOrders() {
	str, _ := json.Marshal(m.GetOrders())

	//if str == "" {
	//	this.Abort("403")
	//}
	//
	//order, err := m.GetOrder(str)
	//
	//if err != nil {
	//	this.Abort("500")
	//} else {
	//	this.Data["json"] = order
	//}

	this.Data["json"] = string(str)
	this.ServeJSON()
}

func (this *OrderController) GetOrder() {
	code := this.Ctx.Input.Param(":code")
	//jsonCode, _ := json.Marshal(code)

	str, _ := m.GetOrder(code)
	strJson, _ := json.Marshal(str)
	this.Data["json"] = string(strJson)
	this.ServeJSON()
}