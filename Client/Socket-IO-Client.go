package main

import (
	"fmt"
	"log"
	"time"

	socketIoClient "github.com/mofadeyunduo/go-socket.io-client"
)

const broadcastKey = "broadcast"
const clientKey = "client message"

func main() {

	// options 為選用功能
	// options := &socketIoClient.Options{
	// 	Transport: "websocket",
	// 	Query:     make(map[string]string),
	// }
	// options.Query["user"] = "user"
	// options.Query["pwd"] = "pass"

	var socketClient *socketIoClient.Client //socket server 連線控制器
	socketHost := "http://localhost:51680"
	socketOptions := &socketIoClient.Options{}
	go func(client *socketIoClient.Client, uri string, opts *socketIoClient.Options) { //由於socket斷線後不會有任何提示，所以需要自行檢測與重連
		for {
			if client == nil || client.Emit("ping", "") != nil {
				var err error
				client, err = socketIoClient.NewClient(uri, socketOptions) //與socket server 連線
				if err != nil {
					client = nil //清空連線
					log.Printf("Reconnect error:%v\n", err)
				} else {
					socketClient = client
					log.Printf("suecess to connect socket.io > '%v'\n", socketHost)

					client.On(broadcastKey, func(msg string) { //開啟監聽Server端廣播
						log.Printf("got broadcast: '%v'\n", msg)
					})

					client.On(clientKey, func(msg string) { //開啟監聽Server端傳訊
						log.Printf("got message: '%v'\n", msg)
					})
				}
			}

			time.Sleep(1 * time.Second) //每秒重新連線，並重新監聽
		}
	}(socketClient, socketHost, socketOptions)

	var sendMsg string
	for {
		fmt.Scanln(&sendMsg)
		if socketClient.Emit(clientKey, sendMsg) == nil {
			log.Printf("sent : '%v'", sendMsg)
		} else {
			log.Println("sent fail !")
		}
		time.Sleep(1 * time.Second) //每秒重新連線，並重新監聽
	}
}
