package controllers

import (
	"github.com/astaxie/beego"

	m "planadotest/models"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) GetHistory() {
	histories := m.GetHistory()

	this.Data["json"] = histories
	this.ServeJSON()
}

//func (this *HistoryController) AddToHistory() {
//	history := m.History{}
//
//	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &history); err != nil {
//		this.Abort("500")
//	}
//
//	result, err := history.AddToHistory()
//	if err != nil {
//		this.Abort("500")
//	} else {
//		this.Data["json"] = result
//	}
//
//	this.ServeJSON()
//}