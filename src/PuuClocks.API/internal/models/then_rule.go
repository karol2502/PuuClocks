package models

var ThenRules = map[int]func(*Game){
	1: SynchronizationRule,
	2: ReverseDirection,
}

func SynchronizationRule(g *Game) {
	g.ExpectedSynchronization = true
}

func ReverseDirection(g *Game) {
	g.Direction = !g.Direction
}

