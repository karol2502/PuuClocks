package models

var WhenRules = map[int]func(*Game, *Card) bool{
	1: SameClockRule,
	2: WehicleCard,
}

func SameClockRule(g *Game, c *Card) bool {
	return g.LastCalledTime == c.Hour
}

func WehicleCard(g *Game, c *Card) bool {
	return c.ID == 1
}
