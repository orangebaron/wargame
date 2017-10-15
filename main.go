// dummy file to run a single server; this'll have more options/run more servers later

package main

import "./src/go/server"

func main() {
	server.RunServer(":8080")
	for true {
	}
}
