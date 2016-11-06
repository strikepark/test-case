package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

	"fmt"
	"errors"
	"strconv"
	"math/rand"
	"math"
)

type Order struct {
	Id int `valid:"Required"`
	Code int64 `valid:"Required"`
	SendAddress string
	RecipientAddress string
	PhoneNumber uint64
	Status string
}

func init() {
	orm.RegisterModel(new(Order))
}

func (u *Order) TableName() string {
	return "orders"
}

func GetOrders() []*Order {
	o := orm.NewOrm()
	//o.Using("default")

	var orders []*Order

	qs := o.QueryTable("orders")

	num, err := qs.All(&orders)

	if err != nil || num > 0 {
		fmt.Println(err)
	}

	return orders
}

func GetOrder(code string) (order Order, err error) {
	uid, _ := strconv.ParseInt(code, 10, 64)

	order = Order{Id: int(math.Abs(rand.Float64())), Code: uid}

	o := orm.NewOrm()

	err = o.Read(&order)

	if err == orm.ErrNoRows {
		fmt.Println(errors.New("not"))
		return order, errors.New("404")
	} else {
		return order, nil
	}
}

func (order Order) NewOrder() (err error) {
	//order := Order{99999999, "str1", "str2", 9999999999, "status1"}

	o := orm.NewOrm()
	//o.Using("default")

	valid := validation.Validation{}
	valid.Required(order.Code, "code")
	valid.Required(order.Id, "Id")

	if valid.HasErrors() {
		return errors.New("Error: order not valid")
	}

	id, err := o.Insert(&order)

	if err != nil {
		fmt.Println(id)
	}

	return nil
}