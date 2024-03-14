package service

import (
	"puuclocks/internal/models"

	"github.com/google/uuid"
)

type Gameplay interface {
	ProcessAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) (bool, error)
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

func (g gameplay) ProcessAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) (bool, error) {
	canBePerformed, err := g.validate.ValidateAction(game, socketID, action, data)
	if err != nil {
		return true, err
	}

	if !canBePerformed {
		return false, nil
	} 

	err = g.conclude.ConcludeAction(game, socketID, action, data)
	if err != nil {
		return true, err
	}

	err = g.action.PerformAction(game, socketID, action, data)
	if err != nil {
		return true, err
	}


	return g.shouldCloseGame(game, socketID, action, data)
}

func (g gameplay) shouldCloseGame(game *models.Game, socketID uuid.UUID, action models.Action, data any) (bool, error) {
	return false, nil
}