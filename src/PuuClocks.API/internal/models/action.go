package models

import "puuclocks/internal/models/actions"

type Action interface {
	GetType() actions.ActionType
	GetData() actions.ActionData
}