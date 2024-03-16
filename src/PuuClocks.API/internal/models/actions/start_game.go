package actions

type StartGame struct {
	Action
}

func (s StartGame) Validate(data string) *Action {
	if data != string(ActionTypeStartGame) {
		return nil
	}

	return &Action{
		Type: ActionTypeStartGame,
	}
}
