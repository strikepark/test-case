package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"strconv"

	m "planadotest/models"

	"encoding/json"
)

// Order controller
type OrderController struct {
	beego.Controller
}

func (this *OrderController) NewOrder() {
	order := m.Order{10, 99999999, "str1", "str2", 9999999999, "status1"}
	err := order.NewOrder()

	if err != nil {
		this.Data["json"] = "Ok"
	} else {
		str, _ := json.Marshal(err)
		this.Data["json"] = string(str)
	}

	this.ServeJSON()
}

func (this *OrderController) GetOrders() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	this.Data["json"] = str
	this.ServeJSON()
}

func (this *OrderController) GetOrder() {
	code, _ := strconv.Atoi(this.Ctx.Input.Param(":code"))
	jsonCode, _ := json.Marshal(code)

	this.Data["json"] = string(jsonCode)
	this.ServeJSON()
}