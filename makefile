get-dependencies:
	go get github.com/gorilla/websocket
build: get-dependencies
	go build
run: build
	./wargame
 
