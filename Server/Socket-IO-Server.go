package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

const broadcastKey = "broadcast"
const clientKey = "client message"

func main() {
	server := socketio.NewServer(nil)

	var Clients = make(map[string]socketio.Conn) //連線中的Client清單
	server.OnConnect("/", func(conn socketio.Conn) error {
		conn.SetContext("")
		log.Printf("connected: %v (url:'%v' LocalAddr:'%v' RemoteAddr:'%v' RemoteHeader:'%v')", conn.ID(),conn.URL(),conn.LocalAddr(),conn.RemoteAddr(),conn.RemoteHeader())
		Clients[conn.ID()] = conn
		return nil
	})

	server.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		delete(Clients, conn.ID()) //清除廣播清單
		log.Printf("%v closed : %v\n", conn.ID(), reason)
	})

	server.OnError("/", func(conn socketio.Conn, err error) {
		log.Println("meet error:", err) //錯誤捕捉
	})

	server.OnEvent("/", clientKey, func(conn socketio.Conn, msg string) {
		if msg != "bye" { //當客戶端發送bye則將客戶端關閉
			log.Printf("get from %v : %v \n", conn.ID(), msg)
			conn.Emit(clientKey, fmt.Sprintf("We have received your message : '%v'", msg)) //回傳給客戶端
		} else {
			conn.Emit("bye !")
			log.Printf("force close client %v", conn.ID())
			conn.Close()
		}
	})

	go func(server *socketio.Server) {
		for {
			for index, value := range Clients {
				value.Emit(broadcastKey, "廣播")
				log.Printf("對%v 發出了廣播\n", index)
			}
			time.Sleep(5 * time.Second)
		}
	}(server)

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:51680...")
	log.Fatal(http.ListenAndServe(":51680", nil))
}
