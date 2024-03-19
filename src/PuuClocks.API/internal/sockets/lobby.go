package sockets

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/service"

	"github.com/google/uuid"
)

type Lobby interface {
	GetID() uuid.UUID

	JoinLobby(Client)
	LeaveLobby(Client)

	ForwardMessage(Message)
	Broadcast(message string)
}

type lobby struct {
	ID uuid.UUID

	Owner Client

	Join    chan Client
	Leave   chan Client
	Forward chan Message

	Clients map[Client]bool

	Game     *models.Game
	Gameplay service.Gameplay

	Settings Settings
}

type Settings struct{}

type Message struct {
	SocketID uuid.UUID
	Data     string
}

func NewLobby(gameplay service.Gameplay) Lobby {
	gameplay
	id := uuid.New()

	l := lobby{
		ID: id,

		Forward: make(chan Message),
		Join:    make(chan Client),
		Leave:   make(chan Client),
		Clients: make(map[Client]bool),

		Gameplay: gameplay,
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
			fmt.Println("Action From: ", msg.SocketID, " Data: ", msg.Data)
			action := actions.ValidateIfUserProvidedActionInstance(msg.Data)
			if action == nil {
				fmt.Println("There was a error during valdiation")
				break
			}
			l.Gameplay.ProcessAction(l.Game, msg.SocketID, *action)
		}
	}
}

func (l *lobby) GetID() uuid.UUID {
	return l.ID
}

func (l *lobby) ForwardMessage(msg Message) {
	l.Forward <- msg
}

func (l *lobby) JoinLobby(c Client) {
	if l.Owner == nil {
		l.Owner = c
	}

	l.Join <- c
}

func (l *lobby) LeaveLobby(c Client) {
	l.Leave <- c
}

func (l *lobby) Broadcast(message string) {
	for c, _ := range l.Clients {
		c.SendMessage([]byte(message))
	}
}
