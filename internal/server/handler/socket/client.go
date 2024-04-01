package socket

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait  = 60 * time.Second
	// must be shorter that pong wait, because ping should be sent before, than read deadline happens
	pingPeriod = pongWait * 9 / 10
)

// Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client
type Client struct {
	conn *websocket.Conn
	hub  *GameplayHub

	write chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.hub.disconnect <- c
		c.conn.Close()
	}()

	// On Pong
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		// Get message from client and send it to hub
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("error reading ws message:", err)
			}
			break
		}
		c.hub.messageChan <- message
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// Ping
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case message, ok := <-c.write:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			// if hub closed c.write chan
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte("closed!"))
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			log.Println("sent to client:", string(message))

			n := len(c.write)
			for i := 0; i < n; i++ {
				w.Write(<-c.write)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
