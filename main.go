package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"

	"os"
	"strconv"

	_ "planadotest/routers"
	"database/sql"
	"log"
)

func main() {
	_, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	} else {
		beego.BConfig.Listen.HTTPPort = 8080
	}

	beego.Run()
}