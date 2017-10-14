package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"../core"
)

var upgrader = websocket.Upgrader{} // don't know if I need this
var game = core.MakeGame()

func sockethandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var plr *core.Player
	for { // hello world websocket function at the moment
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		} else if bytes.Equal(p[:12], []byte("registerplr-")) {
			plr = core.MakePlayer(string(p[12:]), game)
			fmt.Println("Registered player", plr.Name)
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
