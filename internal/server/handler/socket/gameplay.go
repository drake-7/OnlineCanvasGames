package socket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type GameplayWS struct {
	clientMap ClientMap
}

func NewGameplayWS() *GameplayWS {
	return &GameplayWS{
		clientMap: ClientMap{
			mu: &sync.RWMutex{},
		},
	}
}

func (ws *GameplayWS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection:", err)
		return
	}
	defer conn.Close()

	client := &Client{
		Conn: conn,
		Mu:   &sync.RWMutex{},
	}
	ws.clientMap.mu.Lock()
	ws.clientMap.clients[client] = true
	ws.clientMap.mu.Unlock()

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil || mt == websocket.CloseMessage {
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("error writing connection:", err)
			return
		}

		go ws.handleGameplayMessage(message)
	}
}

func (ws *GameplayWS) handleGameplayMessage(message []byte) {
	for client := range ws.clientMap.clients {
		client.Mu.Lock()
		client.Conn.WriteMessage(websocket.TextMessage, message)
		client.Mu.Unlock()
	}

	fmt.Println(string(message))
}
