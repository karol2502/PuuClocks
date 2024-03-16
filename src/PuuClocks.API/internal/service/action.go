package service

import (
	"fmt"
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/repository"

	"github.com/google/uuid"
)

type Action interface {
	PerformAction(game *models.Game, socketID uuid.UUID, action actions.Action) error
}

type action struct {
	redis repository.Redis
}

func newAction(redis repository.Redis) Action {
	return &action{
		redis: redis,
	}
}

func (a action) PerformAction(game *models.Game, socketID uuid.UUID, action actions.Action) error {
	fmt.Println(action.GetType())

	return nil
}
