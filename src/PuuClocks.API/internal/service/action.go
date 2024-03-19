package service

import (
	"fmt"
	"puuclocks/internal/log"
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
	switch action.GetType() {
	case actions.ActionTypeReportTime:
		var function func(*models.Game)
		p := a.findPlayerBySocketID(game, socketID)
		if p == nil {
			return fmt.Errorf("couldn't player who reported time with %d connection ID in %d game", socketID, game.ID)
		}

		drawedCard, err := a.drawCard(game, p)
		if err != nil {
			return err
		}

		function, err = a.customRules(game, drawedCard)
		if err != nil {
			log.Log.Warn(err)
		}
		if err != nil && function != nil {
			function(game)
		} else {
			a.defaultRule(game)
		}

		game.DiscardedCards = append(game.DiscardedCards, drawedCard)
		game.State = models.GameStateAction

		

	case actions.ActionTypeSynchronizationRule:

	}

	return nil
}

func (a action) drawCard(game *models.Game, reporter *models.Player) (models.Card, error) {
	var card *models.Card

	for _, p := range game.Players {
		if p == reporter {
			card, p.PlayingHand = &p.PlayingHand[0], p.PlayingHand[1:]
		}
	}

	if card == nil {
		return models.Card{}, fmt.Errorf("couldn't draw card from player with %d connection", reporter.ConnectionID)
	}

	return *card, nil
}

func (a action) findPlayerBySocketID(game *models.Game, socketID uuid.UUID) *models.Player {
	for _, p := range game.Players {
		if p != nil && p.ConnectionID == socketID {
			return p
		}
	}

	return nil
}

func (a action) customRules(game *models.Game, card models.Card) (func(*models.Game), error) {
	var f func(*models.Game)
	occured := 0
	for _, rule := range game.Rules {
		doesOccure, err := rule.Occure(game, &card)
		if err != nil {
			return nil, err
		}

		if doesOccure {
			occured++
		}

		if occured == 1 {
			f, _ = rule.RetrieveThen()
		}

		if occured > 1 {
			return nil, nil
		}
	}

	return f, nil
}

func (a action) defaultRule(game *models.Game) {
	var exp float64
	if game.Direction == models.GameDirectionClockWise {
		exp = game.ExpectedTime + 1
		if exp > 12 {
			exp -= 12
		}
	} else {
		exp = game.ExpectedTime - 1
		if exp < 0 {
			exp += 12
		}
	}
	game.ExpectedTime = exp
}
