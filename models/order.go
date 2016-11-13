package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"errors"
	"fmt"
	"strconv"
	"encoding/json"
)

type Order struct {
	Id int `json:"id"`
	Code int64 `json:"code";valid:"Required",orm:"unique";`
	SendAddress string `json:"sendAddress",valid:"Required";`
	RecipientAddress string `json:"recipientAddress",valid:"Required";`
	PhoneNumber int64 `json:"phoneNumber",valid:"Required";`
	Status string `json:"status",valid:"Required";`

	ChangeHistories string `orm:"-",json:"changeHistories"`
}

func init() {
	orm.RegisterModel(new(Order))
}

func (u *Order) TableName() string {
	return "orders"
}

func (order Order) CreateOrder() (Order, error) {
	o := orm.NewOrm()


	valid := validation.Validation{}

	valid.Required(order.Code, "Code is required")
	valid.Required(order.SendAddress, "SendAddress is required")
	valid.Required(order.RecipientAddress, "RecipientAddress is required")
	valid.Required(order.PhoneNumber, "PhoneNumber is required")
	valid.Required(order.Status, "Status is required")

	if valid.HasErrors() {
		return order, errors.New("TODO: order not vaild errors")
	}

	_, err := o.Insert(&order)

	if err != nil {
		return order, err
	}

	return order, nil
}

func GetOrders() []*Order {
	o := orm.NewOrm()

	var orders []*Order

	_, err := o.QueryTable("orders").All(&orders)

	if err != nil {
		fmt.Println(err)
	}

	return orders
}

func GetOrder(uid string) (order Order, err error) {
	id, _ := strconv.Atoi(uid)

	order = Order{Id: id}

	o := orm.NewOrm()

	err = o.Read(&order)

	history, _ := json.Marshal(GetHistory(order.Code))

	order.ChangeHistories = string(history)

	if err == orm.ErrNoRows {
		return order, errors.New("404")
	} else {
		return order, nil
	}
}

func UpdateOrder(order Order) (Order, error) {
	o := orm.NewOrm()

	valid := validation.Validation{}

	valid.Required(order.Code, "Code is required")
	valid.Required(order.SendAddress, "SendAddress is required")
	valid.Required(order.RecipientAddress, "RecipientAddress is required")
	valid.Required(order.PhoneNumber, "PhoneNumber is required")
	valid.Required(order.Status, "Status is required")

	if valid.HasErrors() {
		return order, errors.New("TODO: order not vaild errors")
	}

	_, err := o.Update(&order)

	if err != nil {
		return order, errors.New("404")
	} else {
		return order, nil
	}
}

func GetCostumerOrders(phoneNumber string) (orders []*Order, err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable("orders").Filter("phone_number", phoneNumber).All(&orders)

	if err == orm.ErrNoRows {
		return orders, errors.New("404")
	} else {
		for _, order := range orders {
			history, _ := json.Marshal(GetHistory(order.Code))

			order.ChangeHistories = string(history)
		}

		return orders, nil
	}
}
