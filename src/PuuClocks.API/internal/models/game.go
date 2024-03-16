package models

import "github.com/google/uuid"

type Game struct {
	ID             uuid.UUID
	Rules          []Rule
	LastPlayedCard *Card
	DiscardedCards []Card

	AreRulesBroken bool
	Turn           int
	Direction      bool
	
	ExpectedTime   float64
	ExpectedSynchronization *bool

	Players []*Player
	State GameState
	Scoreboard map[*Player]int
}

type GameState int
const (
	GameStateReportTime GameState = iota
	GameStateAction
	GameStateSynchronization
)

type Rule struct {
	WhenID int64
	WhatID int64
}

func NewGame() Game {
	return Game{
		ID: uuid.New(),
	}
}
