package routers

import (
	"github.com/astaxie/beego"
	"planadotest/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world", &controllers.MainController{}, "get:HelloWorld")
}