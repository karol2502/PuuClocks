package models

import "github.com/google/uuid"

type Game struct {
	ID             uuid.UUID
	Rules          []Rule
	LastPlayedCard *Card
	DiscardedCards []Card

	AreRulesBroken bool
	Turn           int
	Direction      GameDirection

	LastCalledTime *float64

	ExpectedTime            float64
	ExpectedSynchronization bool
	Synchronization         map[uuid.UUID]bool

	Players    []*Player
	State      GameState
	Scoreboard map[*Player]int
}

type GameState int

const (
	GameStateReportTime GameState = iota
	GameStateAction
	GameStateSynchronization
)

type GameDirection bool

const (
	GameDirectionClockWise        GameDirection = true
	GameDirectionCounterClockWise GameDirection = false
)

func NewGame() Game {
	return Game{
		ID:    uuid.New(),
		Rules: DefaultRules(),
	}
}
