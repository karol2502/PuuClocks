package service

import (
	"puuclocks/internal/repository"
)

type Service interface {
	Gameplay() Gameplay
	Validate() Validator
	Action() ActionExecutor
	Conclude() FoulChecker
	OutcomeEvaluator() OutcomeEvaluator
}

type service struct {
	gameplay         Gameplay
	validate         Validator
	action           ActionExecutor
	conclude         FoulChecker
	outcomeEvaluator OutcomeEvaluator
}

func NewService(databases repository.Databases) Service {
	validate := newValidate()
	outcomeEvaluator := newOutcomeEvaluator()
	action := newAction(databases.RedisDB())
	conclude := newFoulChecker(databases.RedisDB())

	return &service{
		action:           action,
		conclude:         conclude,
		validate:         validate,
		outcomeEvaluator: outcomeEvaluator,
		gameplay: newGameplay(gamePlayServices{
			validator:        validate,
			actionExecutor:   action,
			faultChecker:     conclude,
			outcomeEvaluator: outcomeEvaluator,
		}),
	}
}

func (s *service) Gameplay() Gameplay {
	return s.gameplay
}

func (s *service) Validate() Validator {
	return s.validate
}

func (s *service) Action() ActionExecutor {
	return s.action
}

func (s *service) Conclude() FoulChecker {
	return s.conclude
}

func (s *service) OutcomeEvaluator() OutcomeEvaluator {
	return s.outcomeEvaluator
}
