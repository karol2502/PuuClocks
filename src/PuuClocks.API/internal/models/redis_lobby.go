package models

type Lobby struct {
	ID int64
	Players []int64
	CurrentGame int64
	Settings Settings
}

type Settings struct {}