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
	"github.com/gorilla/websocket"
	"net/http"
	"time"
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

type myStruct struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var upgrader = websocket.Upgrader{}

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err == nil {
		beego.BConfig.Listen.HTTPPort = port
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static_html/index.html")
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					conn.Close()
				}
			}
		}(conn)

		go func(conn *websocket.Conn) {
			ch := time.Tick(5 * time.Second)

			for range ch {
				conn.WriteJSON(myStruct{
					Username:  "mvansickle",
					FirstName: "Michael",
					LastName:  "Van Sickle",
				})
			}
		}(conn)
	})

	http.ListenAndServe(":" + string(port), nil)

	//beego.SetStaticPath("/", "static_html")

	//beego.Run()
}