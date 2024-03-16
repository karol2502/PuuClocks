package actions

type ActionType string

var (
	ActionTypeStartGame           ActionType = "start-game"

	// Gameplay related
	ActionTypeReportError         ActionType = "report-error"
	ActionTypeReportTime          ActionType = "report-time"
	ActionTypeSynchronizationRule ActionType = "synchronization-rule"
)

type ActionData struct {
	ReportedTime *float64
}

type Action struct {
	Type ActionType
	Data ActionData
}

func (a Action) GetType() ActionType {
	return a.Type
}

func (a Action) GetData() ActionData {
	return a.Data
}

func ValidateIsInstanceAction(data string) *Action {
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
