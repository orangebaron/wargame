get-dependencies:
	go get github.com/gorilla/websocket
build: get-dependencies
	cd server; go build
	mv server/server wargame
run: build
	./wargame
