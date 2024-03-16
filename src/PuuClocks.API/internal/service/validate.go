package service

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"slices"

	"github.com/google/uuid"
)

type Validate interface {
	ValidateAction(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error)
}

type validate struct{}

func newValidate() Validate {
	return &validate{}
}

func (v validate) ValidateAction(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error) {
	if action.GetType() == actions.ActionTypeEndOfTurn {
		return true, nil
	}
	
	switch game.State {
	case models.GameStateReportTime:
		if action.GetType() != actions.ActionTypeReportTime {
			return false, nil
		}

	case models.GameStateAction:
	case models.GameStateSynchronization:
		allowedActions := []actions.ActionType{actions.ActionTypeReportError, actions.ActionTypeSynchronizationRule}
		if !slices.Contains(allowedActions, action.GetType()) {
			return false, nil
		}
	default:
		return false, fmt.Errorf("unknown game state %d", game.State)
	}

	return true, nil
}
