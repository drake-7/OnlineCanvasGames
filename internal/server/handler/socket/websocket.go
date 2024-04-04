package socket

import (
	"log"
	"net/http"
	"strconv"

	gamemodel "github.com/1001bit/OnlineCanvasGames/internal/model/game"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GamesWS struct {
	hubs    map[int]*GameHub
	clients map[*Client]bool

	connect     chan *Client
	disconnect  chan *Client
	messageChan chan []byte
}

func NewGamesWS() (*GamesWS, error) {
	games, err := gamemodel.GetAll()
	if err != nil {
		return nil, err
	}

	hubs := make(map[int]*GameHub)

	for _, game := range games {
		hubs[game.ID] = NewGameHub()
		go hubs[game.ID].Run()
	}

	return &GamesWS{
		hubs:    hubs,
		clients: make(map[*Client]bool),

		connect:     make(chan *Client),
		disconnect:  make(chan *Client),
		messageChan: make(chan []byte),
	}, nil
}

func (ws *GamesWS) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading connection:", err)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage, []byte("wrong id!"))
		conn.Close()
		return
	}
	hub, ok := ws.hubs[id]
	if !ok {
		conn.WriteMessage(websocket.CloseMessage, []byte("wrong id!"))
		conn.Close()
		return
	}

	client := NewClient(conn, ws, hub)
	ws.connect <- client

	go client.readPump()
	go client.writePump()
}

func (ws *GamesWS) Run() {
	for {
		select {
		case client := <-ws.connect:
			client.hub.connect <- client
			ws.clients[client] = true
			log.Println("<GameWS Connect>")

		case client := <-ws.disconnect:
			client.hub.disconnect <- client
			delete(ws.clients, client)
			log.Println("<GameWS Disconnect>")

		case message := <-ws.messageChan:
			ws.handleMessage(message)
		}
	}
}

func (ws *GamesWS) handleMessage(message []byte) {
	for client := range ws.clients {
		select {
		case client.write <- append([]byte("someone said "), message...):
			// if there is client.write waiting, unblock the select
			log.Println("<GameWS Read>:", string(message))
		default:
			// if there is no clint.write waiting, close the connection
			close(client.write)
			delete(ws.clients, client)
		}
	}
}
