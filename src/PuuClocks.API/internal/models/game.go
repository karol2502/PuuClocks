package models

import "github.com/google/uuid"

type Game struct {
	ID     uuid.UUID
	Rules  []Rule
	LastPlayedCard *Card
	DiscardedCards []Card

	AreRulesBroken bool
	Turn           int
	Direction      bool
	ExpectedTime   float64

	State []*Player
	Scoreboard map[*Player]int
}

type Rule struct {
	WhenID int64
	WhatID int64
}

func NewGame() Game {
	return Game{
		ID: uuid.New(),
	}
}
