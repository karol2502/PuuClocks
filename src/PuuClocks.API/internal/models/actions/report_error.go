package actions

type ReportError struct {
	action
}

func (r ReportError) Validate(data string) *action {
	if data != string(ActionTypeReportTime) {
		return nil
	}

	return &action{
		Type: ActionTypeReportError,
	}
}
