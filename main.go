package main

import (
	"golang.org/x/net/websocket"
	"net/http"
	"fmt"
	//"strconv"
	"os"
	"log"
)

func echoHandler(ws *websocket.Conn) {

	for {
		receivedtext := make([]byte, 100)

		n, err := ws.Read(receivedtext)

		if err != nil {
			fmt.Printf("Received: %d bytes\n",n)
		}

		s := string(receivedtext[:n])
		fmt.Printf("Received: %d bytes: %s\n",n,s)
	}
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}

//import (
//	"github.com/astaxie/beego"
//	_ "github.com/lib/pq"
//	"github.com/astaxie/beego/orm"
//
//	"os"
//	"strconv"
//	//"database/sql"
//	//"log"
//
//	_ "planadotest/routers"
//	"fmt"
//)
//
//func init() {
//	orm.RegisterDriver("postgres", orm.DRPostgres)
//	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))
//
//	o := orm.NewOrm()
//
//	res, err := o.Raw("CREATE TABLE IF NOT EXISTS " +
//		`orders("id" SERIAL PRIMARY KEY, "code" bigint UNIQUE, ` +
//		`"send_address" varchar(255), "recipient_address" varchar(255), ` +
//		`"phone_number" bigint, "status" varchar(255))`).Exec()
//
//	if err != nil {
//		num, _ := res.RowsAffected()
//		fmt.Println("postgres row affected nums: ", num)
//	}
//
//	res, err = o.Raw("CREATE TABLE IF NOT EXISTS " +
//		`histories("id" SERIAL PRIMARY KEY, "code" bigint REFERENCES orders (code), ` +
//		`"status" varchar(255), "date" date)`).Exec()
//
//	if err != nil {
//		num, _ := res.RowsAffected()
//		fmt.Println("postgres row affected nums: ", num)
//	}
//}

//func main() {
	//port, err := strconv.Atoi(os.Getenv("PORT"))
	//
	//if err == nil {
	//	beego.BConfig.Listen.HTTPPort = port
	//}
	//
	//beego.SetStaticPath("/", "static_html")
	//
	//beego.Run()
//}