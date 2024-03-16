package actions

type SynchronizationRule struct {
	Action
}

func (s SynchronizationRule) Validate(data string) *Action {
	if data != string(ActionTypeSynchronizationRule) {
		return nil
	}

	return &Action{
		Type: ActionTypeSynchronizationRule,
	}
}
