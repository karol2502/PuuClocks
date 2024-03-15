package actions

type ActionType string

var (
	ActionTypeStartGame           ActionType = "start-game"
	ActionTypeReportError         ActionType = "report-error"
	ActionTypeReportTime          ActionType = "report-time"
	ActionTypeDrawCard            ActionType = "draw-card"
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

	return nil
}
