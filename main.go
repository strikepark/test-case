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
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
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