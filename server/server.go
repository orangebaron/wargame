package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
		} else if bytes.Equal(p[:12], []byte("registerplr:")) {
			plr = core.MakePlayer(string(p[12:]), game)
			fmt.Println("Registered player", plr.Name)
		} else if bytes.Equal(p[:9], []byte("gettiles:")) {
			s := string(p[9:])
			commaLoc := strings.Index(s, ",")
			xstartstring, s := s[:commaLoc], s[commaLoc+1:]
			commaLoc = strings.Index(s, ",")
			xendstring, s := s[:commaLoc], s[commaLoc+1:]
			commaLoc = strings.Index(s, ",")
			ystartstring, yendstring := s[:commaLoc], s[commaLoc+1:]
			xstart, _ := strconv.Atoi(xstartstring)
			xend, _ := strconv.Atoi(xendstring)
			ystart, _ := strconv.Atoi(ystartstring)
			yend, _ := strconv.Atoi(yendstring)
			for x := xstart; x <= xend; x++ {
				for y := ystart; y <= yend; y++ {
					strToSend := "block:" + strconv.Itoa(x) + "," + strconv.Itoa(y) + ","
					if game.UnitMap[core.Vec{X: x, Y: y}] != nil {
						strToSend += game.UnitMap[core.Vec{X: x, Y: y}].Stats.Name
					} else {
						strToSend += "water"
					}
					conn.WriteMessage(websocket.TextMessage, []byte(strToSend))
					fmt.Println(strToSend)
				}
			}
		}
	}
}

func main() {
	fmt.Println("Server started")
	http.HandleFunc("/socket", sockethandler)
	http.Handle("/game/", http.StripPrefix("/game/", http.FileServer(http.Dir("/home/sam/Documents/go/wargame/files"))))
	http.ListenAndServe(":8080", nil)
}
