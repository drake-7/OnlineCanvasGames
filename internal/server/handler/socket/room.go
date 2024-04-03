package socket

type GameRoom struct {
	clients map[*Client]bool
}

func NewGameRoom() *GameRoom {
	return &GameRoom{
		clients: make(map[*Client]bool),
	}
}
