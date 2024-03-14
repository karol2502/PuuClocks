package models

type Action string

var (
	ActionStartGame Action = "start-game"
	ActionReportError Action = "report-error"
	ActionReportTime Action = "report-time"
	ActionDrawCard Action = "draw-card"
	ActionSynchronizationRule Action = "synchronization-rule"
)
