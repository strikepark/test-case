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
		`orders("Id" integer, "Code" bigint, ` +
		`"SendAddress" varchar(255), "RecipientAddress" varchar(255), ` +
		`"PhoneNumber" bigint, "Status" varchar(255))`).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("postgres row affected nums: ", num)
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