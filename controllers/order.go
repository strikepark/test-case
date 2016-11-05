package controllers

import (
	"github.com/astaxie/beego"

	"encoding/json"
	"strconv"
)

// Order controller
type OrderController struct {
	beego.Controller
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