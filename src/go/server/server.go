package server

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
// Game is the game that the server is running.
var Game = core.MakeGame()

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
			plr = core.MakePlayer(string(p[12:]), Game)
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
					if Game.UnitMap[core.Vec{X: x, Y: y}] != nil {
						strToSend += Game.UnitMap[core.Vec{X: x, Y: y}].Stats.Name
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

// RunServer sets up a server at a given port.
func RunServer(port string) {
	fmt.Println("Starting server")
	http.HandleFunc("/socket", sockethandler)
	http.Handle("/game/", http.StripPrefix("/game/", http.FileServer(http.Dir("src/webpage"))))
	http.ListenAndServe(port, nil)
}
