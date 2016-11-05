package main

import (
	"github.com/astaxie/beego"

	"os"
	"strconv"

	_ "planadotest/routers"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	} else {
		beego.BConfig.Listen.HTTPPort = 8080
	}

	beego.Run()
}