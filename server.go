package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func sockethandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	var plr *player
	for { // hello world websocket function at the moment
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		} else if bytes.Equal(p, []byte("registerplr")) {
			plr = &player{}
			fmt.Println("Registered player", plr.name)
		}
		conn.WriteMessage(websocket.TextMessage, []byte("dummy message"))
	}
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/socket", sockethandler)
	http.Handle("/game/", http.StripPrefix("/game/", http.FileServer(http.Dir("/home/sam/Documents/go/wargame/files"))))
	http.ListenAndServe(":8080", nil)
}
