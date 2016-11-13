package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	content, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}

	this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	this.Ctx.Output.Body(content)
	this.RenderBytes()
}

// TODO: убрать дублирование
func (this *MainController) Error404() {
    content, err := ioutil.ReadFile("static/index.html")
    if err != nil {
        panic(err)
    }

    this.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
    this.Ctx.Output.Body(content)
    this.RenderBytes()
}
