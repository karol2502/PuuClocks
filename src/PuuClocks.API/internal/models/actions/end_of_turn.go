package actions

import "puuclocks/internal/models"

type EndOfTurn struct {
	action
}

func NewActionEndOfTurn(state models.GameState) action {
	return action{
		Type: ActionTypeEndOfTurn,
		Data: ActionData{
			State: &state,
		},
	}
}