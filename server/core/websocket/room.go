package websocket

import (
	"lanos/pkg/id"
	"sync"
)

type Room struct {
	Id      string
	Clients map[string]*Client
	mu      sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		Id:      id.NewID(),
		Clients: make(map[string]*Client),
	}
}

func (r *Room) AddClient(c *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Clients[c.ID] = c
}

func (r *Room) RemoveClient(client *Client) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.Clients, client.ID)
}

func (r *Room) GetAllClient() map[string]*Client {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.Clients

}

func (r *Room) IsEmpty() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.Clients) == 0
}
