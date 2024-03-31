package socket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Mu   *sync.RWMutex
}

type ClientMap struct {
	clients map[*Client]bool
	mu      *sync.RWMutex
}

type WebsocketServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleMessage(message []byte)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
