package socket

import (
	"log"
	"net/http"
)

type GameplayHub struct {
	clients map[*Client]bool

	connect     chan *Client
	disconnect  chan *Client
	messageChan chan []byte
}

func ServeWS(hub *GameplayHub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection:", err)
		return
	}

	client := &Client{
		conn: conn,
		hub:  hub,
	}
	client.hub.connect <- client

	go client.readPump()
	go client.writePump()
}

func NewGameplayHub() *GameplayHub {
	return &GameplayHub{
		clients: make(map[*Client]bool),

		connect:     make(chan *Client),
		disconnect:  make(chan *Client),
		messageChan: make(chan []byte),
	}
}

func (hub *GameplayHub) Run() {
	for {
		select {
		case client := <-hub.connect:
			hub.clients[client] = true
		case client := <-hub.disconnect:
			delete(hub.clients, client)
		case message := <-hub.messageChan:
			handleMessage(message)
		}
	}
}

func handleMessage(message []byte) {
	log.Println(string(message))
}
