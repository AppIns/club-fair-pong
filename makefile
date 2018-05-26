all:
	go install github.com/gorilla/websocket
	go run src/*.go
