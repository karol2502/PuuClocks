package service

import (
	"fmt"
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
	if action.GetType() == actions.ActionTypeReportError || action.GetType() == actions.ActionTypeEndOfTurn {
		return nil
	}

	var player *models.Player
	if game.AreRulesBroken {
		for _, p := range game.Players {
			if p != nil && p.ConnectionID == socketID {
				player = p
			}
		}

		if player == nil {
			return fmt.Errorf("couldn't obtain player with %d connection to determine who did action", socketID)
		}

		game.LastActionCaller = player

		return nil
	}

	game.LastActionCaller = player

	switch action.GetType() {
	case actions.ActionTypeSynchronizationRule:
		if game.ExpectedSynchronization {
			return nil
		}
	case actions.ActionTypeReportTime:
		if game.Players[game.Turn] == player {
			return nil
		}
	}

	game.AreRulesBroken = true

	return nil
}
