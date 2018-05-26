package main

import (
  "net/http"

  "time"
  "fmt"
  "strconv"

  "github.com/gorilla/websocket"
)

// var writes []*websocket.Conn
// var writelocks []*sync.Mutex
// var ticklock sync.RWLock

func wsGetBoard(w http.ResponseWriter, r *http.Request) {

  ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	defer ws.Close()
	defer ws.WriteMessage(websocket.CloseMessage, []byte{})

	if err != nil {
		fmt.Println("Error upgrading to websocket (ws.go)")
		return
	}

  // _, headtail, err := ws.ReadMessage()
  //
  // if err != nil {
  //   return
  // }

  for {

    ws.WriteMessage(websocket.TextMessage, []byte(format()))
    time.Sleep(time.Millisecond * 10)
  }

}

func format() string {
  var str string

  paddlelock.RLock()
  str += strconv.Itoa(p1pad) + ";"
  str += strconv.Itoa(p2pad) + ";"
  paddlelock.RUnlock()
  ballock.RLock()
  str += strconv.Itoa(ballPos[0]) + ";"
  str += strconv.Itoa(ballPos[1]) + ";"
  ballock.RUnlock()

  return str
}
