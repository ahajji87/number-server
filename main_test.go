package main

import (
	"flag"
	"log"
	"math/rand"
	"net/url"
	"sync"
	"testing"

	"github.com/gorilla/websocket"
)

var result int
var letterRunes = []rune("0123456789")
var addr = flag.String("addr", "localhost:4000", "ws service address")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkNumberServer(b *testing.B) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	var waitGroup sync.WaitGroup

	for i := 0; i < b.N; i++ {
		waitGroup.Add(1)
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Fatal("dial:", err)
		}
		go func() {

			if err := c.WriteMessage(websocket.TextMessage, []byte(RandStringRunes(9))); err != nil {
				log.Println("write:", err)
			}
			_, _, err := c.ReadMessage()
			if ce, ok := err.(*websocket.CloseError); ok {
				switch ce.Code {
				case websocket.CloseNormalClosure,
					websocket.CloseGoingAway,
					websocket.CloseNoStatusReceived:
					return
				}
			}
			defer waitGroup.Done()
		}()
	}
}
