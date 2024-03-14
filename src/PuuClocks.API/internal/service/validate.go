package service

import (
	"puuclocks/internal/models"

	"github.com/google/uuid"
)

type Validate interface {
	ValidateAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) (bool, error)
}

type validate struct{}

func newValidate() Validate {
	return &validate{}
}

func (v validate) ValidateAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) (bool,error) {
	return true, nil
}


