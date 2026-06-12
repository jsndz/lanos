package websocket

import (
	"errors"
	"sync"
)

type Hub struct {
	Clients map[string]*Client
	Rooms   map[string]*Room
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[string]*Client),
		Rooms:   make(map[string]*Room),
	}
}

func (h *Hub) Register(client *Client, id string) {
	h.Clients[id] = client
}

func (h *Hub) CreateRoom() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	room := NewRoom()
	h.Rooms[room.Id] = room
	return room.Id
}

func (h *Hub) GetRoom(id string) (*Room, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	room, ok := h.Rooms[id]
	if !ok {
		return nil, errors.New("Room doesn't exist")
	}
	return room, nil
}

func (h *Hub) DeleteRoom(roomId string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.Rooms, roomId)
}
