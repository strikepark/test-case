//package main
//
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
//
//func main() {
//	port, err := strconv.Atoi(os.Getenv("PORT"))
//
//	if err == nil {
//		beego.BConfig.Listen.HTTPPort = port
//	}
//
//	beego.SetStaticPath("/", "static_html")
//
//	beego.Run()
//}

package main

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"io"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// message sent to us by the javascript client
type message struct {
	Handle string `json:"handle"`
	Text   string `json:"text"`
}

// validateMessage so that we know it's valid JSON and contains a Handle and
// Text
func validateMessage(data []byte) (message, error) {
	var msg message

	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, errors.Wrap(err, "Unmarshaling message")
	}

	if msg.Handle == "" && msg.Text == "" {
		return msg, errors.New("Message has no Handle or Text")
	}

	return msg, nil
}

// handleWebsocket connection.
func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m := "Unable to upgrade to websockets"
		fmt.Println("err", err, m)
		http.Error(w, m, http.StatusBadRequest)
		return
	}

	fmt.Println("РАБОТАЕТ")

	for {
		mt, data, err := ws.ReadMessage()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Websocket closed!")
			} else {
				fmt.Println("Error reading websocket message")
			}

			break
		}

		switch mt {
		case websocket.TextMessage:
			msg, err := validateMessage(data)
			if err != nil {
				fmt.Println("msg:", msg, "err:", err)
				break
			}
			fmt.Println(string(data[:]))
		default:
			fmt.Println("Unknown Message!")
		}
	}

	ws.WriteMessage(websocket.CloseMessage, []byte{})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT ERROR")
	}

	http.Handle("/", http.FileServer(http.Dir("./static_html")))
	http.HandleFunc("/ws", handleWebsocket)
	fmt.Println(http.ListenAndServe(":" + port, nil))
}
