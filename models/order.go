package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Order struct {
	Code uint32
	SendAddress string
	RecipientAddress string
	PhoneNumber uint64
	Status string
}

func init() {
	orm.RegisterModel(new(Order))
}

func (order Order) NewOrder() {
	o := orm.NewOrm()
	o.Using("default")

	//order := Order{99999999, "str1", "str2", 9999999999, "status1"}

	fmt.Println(o.Insert(order))
}