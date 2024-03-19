package service

import (
	"puuclocks/internal/models"
	"puuclocks/internal/models/actions"
	"puuclocks/internal/repository"

	"github.com/google/uuid"
)

type FoulChecker interface {
	CheckForFaul(game *models.Game, socketID uuid.UUID, action actions.Action) error
}

type foulChecker struct {
	redis repository.Redis
}

func newFoulChecker(redis repository.Redis) FoulChecker {
	return &foulChecker{
		redis: redis,
	}
}

func (c foulChecker) CheckForFaul(game *models.Game, socketID uuid.UUID, action actions.Action) error {
	if game.AreRulesBroken {
		return nil
	}

	return nil
}
