package actions

import "puuclocks/internal/models"

type ActionType string

var (
	ActionTypeStartGame ActionType = "start-game"

	// ServerActions
	ActionTypeEndOfTurn ActionType = "validate-turn-result"

	// Gameplay related
	ActionTypeReportError         ActionType = "report-error"
	ActionTypeReportTime          ActionType = "report-time"
	ActionTypeSynchronizationRule ActionType = "synchronization-rule"
)

type ActionData struct {
	ReportedTime *float64
	State *models.GameState
}

type Action interface {
	GetType() ActionType
	GetData() ActionData
}

type action struct {
	Type ActionType
	Data ActionData
}

func (a action) GetType() ActionType {
	return a.Type
}

func (a action) GetData() ActionData {
	return a.Data
}

func ValidateIsInstanceAction(data string) *action {
	if action := (ReportTime{}).Validate(data); action != nil {
		return action
	}

	if action := (ReportError{}).Validate(data); action != nil {
		return action
	}

	if action := (StartGame{}).Validate(data); action != nil {
		return action
	}

	if action := (SynchronizationRule{}).Validate(data); action != nil {
		return action
	}

	return nil
}
