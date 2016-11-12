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
		`orders("id" SERIAL PRIMARY KEY, "code" bigint UNIQUE, ` +
		`"send_address" varchar(255), "recipient_address" varchar(255), ` +
		`"phone_number" bigint, "status" varchar(255))`).Exec()

	if err != nil {
		num, _ := res.RowsAffected()
		fmt.Println("postgres row affected nums: ", num)
	}

	res, err = o.Raw("CREATE TABLE IF NOT EXISTS " +
		`histories("id" SERIAL PRIMARY KEY, "code" bigint REFERENCES orders (code), ` +
		`"status" varchar(255), "date" date)`).Exec()

	if err != nil {
		num, _ := res.RowsAffected()
		fmt.Println("postgres row affected nums: ", num)
	}
}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}

	for _, path := range []string{"js/bundle.js"} {
		beego.SetStaticPath("/" + path, "static/" + path)
	}

	beego.Run()
}
