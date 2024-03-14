package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/repository"

	"github.com/google/uuid"
)

type Conclude interface {
	ConcludeAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) error
}

type conclude struct {
	redis repository.Redis
}

func newConclude(redis repository.Redis) Conclude {
	return &conclude{
		redis: redis,
	}
}

func (c conclude) ConcludeAction(game *models.Game, socketID uuid.UUID, action models.Action, data any) error {
	return nil
}
