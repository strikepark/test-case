package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) HelloWorld() {
	this.Data["Text"] = "Hello, world!"
	this.TplName = "home.tpl"
}
