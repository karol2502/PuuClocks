package sockets

import (
	"github.com/google/uuid"
)

type LobbyManager interface {
	CreateLobby() Lobby
	FindLobby(uuid.UUID) Lobby
}

type lobbyManager struct {
	Lobbies map[uuid.UUID]Lobby
}

func NewLobbyManager() LobbyManager {
	return &lobbyManager{
		Lobbies: make(map[uuid.UUID]Lobby),
	}
}

func (l lobbyManager) CreateLobby() Lobby {
	lobby := NewLobby()
	id := lobby.GetID()
	l.Lobbies[id] = lobby

	return lobby
}

func (l lobbyManager) FindLobby(id uuid.UUID) Lobby {
	return l.Lobbies[id]
}
