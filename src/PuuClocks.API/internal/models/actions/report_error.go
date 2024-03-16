package actions

type ReportError struct {
	Action
}

func (r ReportError) Validate(data string) *Action {
	if data != string(ActionTypeReportTime) {
		return nil
	}

	return &Action{
		Type: ActionTypeReportError,
	}
}
