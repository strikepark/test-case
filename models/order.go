package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"errors"
	"fmt"
)

type Order struct {
	Id int
	Code int64 `valid:"Required"`
	SendAddress string `valid:"Required"`
	RecipientAddress string `valid:"Required"`
	PhoneNumber uint64 `valid:"Required"`
	Status string `valid:"Required"`
}

func init() {
	orm.RegisterModel(new(Order))
}

func (u *Order) TableName() string {
	return "orders"
}

func (order Order) CreateOrder() (result Order, err error) {
	o := orm.NewOrm()

	valid := validation.Validation{}

	valid.Required(order.Code, "Code is required")
	valid.Required(order.SendAddress, "Code is required")
	valid.Required(order.RecipientAddress, "Code is required")
	valid.Required(order.PhoneNumber, "Code is required")
	valid.Required(order.Status, "Code is required")

	if valid.HasErrors() {
		return result, errors.New("TODO: order not vaild errors")
	}

	_, err = o.Insert(&order)

	if err != nil {
		return result, nil
	}

	return order, nil
}

func GetOrders() []*Order {
	o := orm.NewOrm()

	var orders []*Order

	num, err := o.QueryTable("orders").All(&orders)

	if num > 0 {
		fmt.Println(err)
	}

	return orders
}

//func GetOrder(code string) (order Order, err error) {
//	uid, _ := strconv.ParseInt(code, 10, 64)
//
//	order = Order{Code: uid}
//	fmt.Println("Code ", uid)
//
//	o := orm.NewOrm()
//
//	err = o.Read(&order)
//
//	if err == orm.ErrNoRows {
//		fmt.Println(errors.New("not"))
//		return order, errors.New("404")
//	} else {
//		return order, nil
//	}
//}
