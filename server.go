package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func sockethandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	for { // hello world websocket function at the moment
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		} else if bytes.Equal(p, []byte("ready")) {
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte("dummy message"))
	}
}

func main() {
	http.HandleFunc("/", sockethandler)
	http.Handle("/", http.FileServer(http.Dir("/home/sam/Documents/go/wargame/files")))
	http.ListenAndServe(":8080", nil)
}
