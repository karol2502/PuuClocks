package models

type Game struct {
	ID int64
	Rules []Rule
	Winner *int64
	//EventHistory
	LastPlayedCard *Card
	DiscardedCards []Card
	AreRulesBroken bool
	Turn int
	Direction bool
	ExpectedTime float64
}

type Rule struct {
	WhenID int64
	WhatID int64
}