package service

type OutcomeEvaluator interface {
	ShouldPunishOrAward()
}

type outcomeEvaluator struct {
}

func newOutcomeEvaluator() OutcomeEvaluator {
	return &outcomeEvaluator{}
}


func (o outcomeEvaluator) ShouldPunishOrAward() {
	
}