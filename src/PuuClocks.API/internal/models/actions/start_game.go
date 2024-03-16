package actions

type StartGame struct {
	action
}

func (s StartGame) Validate(data string) *action {
	if data != string(ActionTypeStartGame) {
		return nil
	}

	return &action{
		Type: ActionTypeStartGame,
	}
}
