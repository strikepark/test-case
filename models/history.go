package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"fmt"
	"time"
	"errors"
)

type History struct {
	Id int
	Code int64 `valid:"Required";orm:"unique"`
	Status string `valid:"Required"`
	Date time.Time `valid:"Required"`
}

func init() {
	orm.RegisterModel(new(History))
}

func (u *History) TableName() string {
	return "histories"
}

func GetHistory(code int64) (histories []*History) {
	o := orm.NewOrm()

	_, err := o.QueryTable("histories").Filter("code", code).All(&histories)

	if err != nil {
		fmt.Println(err)
	}

	return histories
}

func (history History) AddToHistory() (History, error) {
	o := orm.NewOrm()

	valid := validation.Validation{}

	valid.Required(history.Code, "Code is required")
	valid.Required(history.Status, "SendAddress is required")
	valid.Required(history.Date, "RecipientAddress is required")

	if valid.HasErrors() {
		return history, errors.New("TODO: order not vaild errors")
	}

	_, err := o.Insert(&history)

	if err != nil {
		return history, err
	}

	return history, nil
}