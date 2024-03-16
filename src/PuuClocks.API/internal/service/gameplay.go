package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"

	"github.com/google/uuid"
)

type Gameplay interface {
	ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error)
}

type gameplay struct {
	validate Validate
	action   Action
	conclude Conclude
}

type gamePlayServices struct {
	Validate Validate
	Action   Action
	Conclude Conclude
}

func newGameplay(services gamePlayServices) Gameplay {
	return &gameplay{
		validate: services.Validate,
		action:   services.Action,
		conclude: services.Conclude,
	}
}

func (g gameplay) ProcessAction(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error) {
	canBePerformed, err := g.validate.ValidateAction(game, socketID, action)
	if err != nil {
		return true, err
	}

	if !canBePerformed {
		return false, nil
	}

	err = g.conclude.ConcludeAction(game, socketID, action)
	if err != nil {
		return true, err
	}

	err = g.action.PerformAction(game, socketID, action)
	if err != nil {
		return true, err
	}

	return g.shouldCloseGame(game, socketID, action)
}

func (g gameplay) shouldCloseGame(game *models.Game, socketID uuid.UUID, action actions.Action) (bool, error) {
	return false, nil
}
