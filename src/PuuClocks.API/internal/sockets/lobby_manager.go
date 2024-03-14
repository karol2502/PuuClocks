package sockets

import (
	"puuclocks/internal/service"

	"github.com/google/uuid"
)

type LobbyManager interface {
	CreateLobby(gameplay service.Gameplay) Lobby
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

func (l lobbyManager) CreateLobby(gameplay service.Gameplay) Lobby {
	lobby := NewLobby(gameplay)
	id := lobby.GetID()
	l.Lobbies[id] = lobby

	return lobby
}

func (l lobbyManager) FindLobby(id uuid.UUID) Lobby {
	return l.Lobbies[id]
}
