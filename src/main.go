package main

import (
  "net/http"
  "log"
  "fmt"
  "io"
  "sync"
  "time"
)

var ballPos [2]int = [2]int{3, 5}
var ballForce [2]int = [2]int{-1, 1}
var ballock sync.RWMutex
var p1pad int = 20
var p2pad int = 0
var paddlelock sync.RWMutex

func main()  {
  go tick(1)
  http.HandleFunc("/", handleRoot)
  http.HandleFunc("/api/v1/getboard", handleGetBoard)
  http.HandleFunc("/api/ws/getboard", wsGetBoard)
  http.HandleFunc("/api/v1/moveleft/1", move1Left)
  http.HandleFunc("/api/v1/moveright/1", move1Right)
  http.HandleFunc("/api/v1/moveleft/2", move2Left)
  http.HandleFunc("/api/v1/moveright/2", move2Right)

  fmt.Println("Starting server on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "public/" + r.URL.Path)
}

func tick(count int){
  ballock.Lock()
  paddlelock.Lock()
  ballPos[0] += ballForce[0]
  ballPos[1] += ballForce[1]
  if ballPos[1] == 2 {
    if p1pad == ballPos[0] {
      ballForce[1] = 1
    }
    if p1pad + 4 == ballPos[0] {
      ballForce[1] = 1
      ballForce[0] = 2
    } else if p1pad >= ballPos[0] - 4 && p1pad < ballPos[0] {
      ballForce[1] = 1
      ballForce[0] = 1
    }
    if p1pad - 4 == ballPos[0] {
      ballForce[1] = 1
      ballForce[0] = -2
    } else if p1pad <= ballPos[0] + 4 && p1pad > ballPos[0] {
      ballForce[1] = 1
      ballForce[0] = -1
    }

  }
  if ballPos[1] == 22 {
    if p2pad == ballPos[0] {
      ballForce[1] = -1
    }
    if p2pad + 4 == ballPos[0] {
      ballForce[1] = -1
      ballForce[0] = 2
    } else if p2pad >= ballPos[0] - 4 && p2pad < ballPos[0] {
      ballForce[1] = -1
      ballForce[0] = 1
    }
    if p2pad - 4 == ballPos[0] {
      ballForce[1] = -1
      ballForce[0] = -2
    } else if p2pad <= ballPos[0] + 4 && p2pad > ballPos[0] {
      ballForce[1] = -1
      ballForce[0] = -1
    }

  }
  if ballPos[0] >= 40 || 0 >= ballPos[0] {
    ballForce[0] = -ballForce[0]
  }

  if ballPos[1] > 25 || ballPos[1] < 0 {
    ballPos[0] = 10
    ballPos[1] = 13
    ballForce = [2]int{0,1}
    count = 1
  }

  paddlelock.Unlock()
  ballock.Unlock()
  genTmap()
  multi := 1000 / ((count / 20) + 1)
  time.Sleep(time.Duration(multi * 1000 * 50) + 50 * time.Millisecond)
  fmt.Println(multi)
  tick(count + 1)
}

func handleGetBoard(w http.ResponseWriter, r *http.Request) {
  genTmap()
  tmaplock.RLock()
  io.WriteString(w, tmap)
  tmaplock.RUnlock()
}

var tmap string
var tmaplock sync.RWMutex

func genTmap()  {
  tmaplock.Lock()
  paddlelock.RLock()
  tmap = xZeros(40) + "\n"
  tmap += xZeros(p1pad - 3) + xOnes(5) + xZeros((40 - p1pad) - 2) + "\n"
  ballock.RLock()
  for i:=0; i<ballPos[1] - 1 && i < 20; i++ {
    tmap += xZeros(40) + "\n"
  }
  tmap += xZeros(ballPos[0] - 1)
  tmap += "1"
  tmap += xZeros((40 - ballPos[0])) + "\n"
  ballock.RUnlock()
  for i:= 0; i<21 - ballPos[1]; i++ {
    tmap += xZeros(40) + "\n"
  }
  tmap += xZeros(p2pad - 3) + xOnes(5) + xZeros((50 - p2pad) - 2) + "\n"
  tmap += xZeros(40)
  paddlelock.RUnlock()
  tmaplock.Unlock()
}



func move1Left(w http.ResponseWriter, r *http.Request) {
  paddlelock.Lock()
  p1pad -= 2
  for ;p1pad < -2; {
    p1pad = 39
  }
  paddlelock.Unlock()
}

func move1Right(w http.ResponseWriter, r *http.Request) {
  paddlelock.Lock()
  p1pad += 2
  for ;p1pad > 42; {
    p1pad = 1
  }
  paddlelock.Unlock()
}

func move2Left(w http.ResponseWriter, r *http.Request) {
  paddlelock.Lock()
  p2pad -= 2
  for ;p2pad < -3; {
    p2pad = 40
  }
  paddlelock.Unlock()
}

func move2Right(w http.ResponseWriter, r *http.Request) {
  paddlelock.Lock()
  p2pad += 2
  for ;p2pad > 43; {
    p2pad = 0
  }
  paddlelock.Unlock()
}
