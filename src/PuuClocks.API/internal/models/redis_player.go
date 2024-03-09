package models

type Player struct {
	AccountID    int64
	ConnectionID int64
	Nickname     string
	PlayingHand  []Card
}
