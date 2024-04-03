package socket

import "log"

type GameHub struct {
	rooms   map[*GameRoom]bool
	clients map[*Client]bool

	connect    chan *Client
	disconnect chan *Client
}

func NewGameHub() *GameHub {
	return &GameHub{
		rooms:   make(map[*GameRoom]bool),
		clients: make(map[*Client]bool),

		connect:    make(chan *Client),
		disconnect: make(chan *Client),
	}
}

func (hub *GameHub) Run() {
	for {
		select {
		case client := <-hub.connect:
			log.Println("<GameHub Connect>")
			hub.clients[client] = true
		case client := <-hub.disconnect:
			delete(hub.clients, client)
		}
	}
}
