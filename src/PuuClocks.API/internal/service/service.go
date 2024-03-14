package service

import (
	"puuclocks/internal/repository"
)

type Service interface{
	Gameplay() Gameplay
	Validate() Validate
	Action() Action
	Conclude() Conclude 
}

type service struct{
	gameplay Gameplay
	validate Validate
	action Action
	conclude Conclude
}

func NewService(databases repository.Databases) Service {
	validate := newValidate()
	action := newAction(databases.RedisDB())
	conclude := newConclude(databases.RedisDB())
	
	return &service{
		action: action,
		conclude: conclude,
		validate: validate,
		gameplay: newGameplay(gamePlayServices{
			Validate: validate,
			Action: action,
			Conclude: conclude,
		}),
	}
}

func (s *service) Gameplay() Gameplay {
	return s.gameplay
}

func (s *service) Validate() Validate {
	return s.validate
}

func (s *service) Action() Action {
	return s.action
}

func (s *service) Conclude() Conclude {
	return s.conclude
}