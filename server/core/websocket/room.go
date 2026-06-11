package websocket

type Room struct {
	RoomId  string
	Clients map[string]*Client
}
