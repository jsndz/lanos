package websocket

type Hub struct {
	Clients map[string]*Client
	Rooms   map[string]*Room
}

func (h *Hub) Register(client *Client, id string) {
	h.Clients[id] = client
}
