package sockets

import (
	"github.com/google/uuid"
)

type Lobby interface {
	GetID() uuid.UUID

	JoinLobby(Client)
	LeaveLobby(Client)

	ForwardMessage([]byte)
}

type lobby struct {
	ID uuid.UUID

	Owner Client

	Join    chan Client
	Leave   chan Client
	Forward chan []byte

	Clients map[Client]bool
}

func NewLobby() Lobby {
	id := uuid.New()

	l := lobby{
		ID: id,

		Forward: make(chan []byte),
		Join:    make(chan Client),
		Leave:   make(chan Client),
		Clients: make(map[Client]bool),
	}

	go l.run()

	return &l
}

func (l *lobby) run() {
	for {
		select {
		case client := <-l.Join:
			l.Clients[client] = true
		case client := <-l.Leave:
			delete(l.Clients, client)
			client.Close()
		case msg := <-l.Forward:
			for c := range l.Clients {
				c.ReceiveMessage(msg)
			}
		}
	}
}

func (l lobby) GetID() uuid.UUID {
	return l.ID
}

func (l lobby) ForwardMessage(msg []byte) {
	l.Forward <- msg
}

func (l *lobby) JoinLobby(c Client) {
	if l.Owner == nil {
		l.Owner = c
	}

	l.Join <- c
}

func (l lobby) LeaveLobby(c Client) {
	l.Leave <- c
}
