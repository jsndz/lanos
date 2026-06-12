package websocket

import (
	"lanos/pkg/id"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:3000" {
			return true
		} else {
			return false
		}
	},
}

func WsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{
			ID:   id.NewID(),
			Conn: conn,
		}
		hub.Register(client, client.ID)
		go client.Read()
		go client.Write()
	}
}
