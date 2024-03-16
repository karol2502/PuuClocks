package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/repository"

	"github.com/google/uuid"
)

type Conclude interface {
	ConcludeAction(game *models.Game, socketID uuid.UUID, action actions.Action) error
}

type conclude struct {
	redis repository.Redis
}

func newConclude(redis repository.Redis) Conclude {
	return &conclude{
		redis: redis,
	}
}

func (c conclude) ConcludeAction(game *models.Game, socketID uuid.UUID, action actions.Action) error {
	if game.AreRulesBroken {
		return nil
	}

	return nil
}
