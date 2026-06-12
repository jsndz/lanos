package websocket

import (
	"lanos/pkg/id"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn  *websocket.Conn
	ID    string
	send  chan []byte
	Rooms map[string]struct{}
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Conn:  conn,
		ID:    id.NewID(),
		Rooms: make(map[string]struct{}),
	}
}

func (c *Client) Read() {
	for {
		//blocking
		_, _, _ = c.Conn.ReadMessage()

	}
}

func (c *Client) Write() {
	for data := range c.send {
		c.Conn.WriteMessage(websocket.BinaryMessage, data)
	}
}
