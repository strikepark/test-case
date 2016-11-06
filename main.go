package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"

	"os"
	"strconv"
	//"database/sql"
	//"log"

	_ "planadotest/routers"
	"fmt"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))

	o := orm.NewOrm()

	res, err := o.Raw("CREATE TABLE IF NOT EXISTS " +
		`orders("id" integer, "code" bigint, ` +
		`"send_address" varchar(255), "recipient_address" varchar(255), ` +
		`"phone_number" bigint, "status" varchar(255))`).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("postgres row affected nums: ", num)
	}

	type Order struct {
		Id int
		Code int64
		SendAddress string
		RecipientAddress string
		PhoneNumber uint64
		Status string
	}

	var order Order

	num, err := o.Raw("(SELECT * FROM orders)").QueryRows(&order)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	} else {
		beego.BConfig.Listen.HTTPPort = 8080
	}

	beego.Run()
}