package service

import ()

type Broadcast interface {
	SetBroadcast(func(message string))

	EndOfReportTimeTurn(reporterNickname string)
}

type broadcast struct {
	broadcast func(message string)
}

func newBroadcast() Broadcast {
	return &broadcast{}
}

func (b broadcast) SetBroadcast(f func(message string)) {
	b.broadcast = f
}

func (b broadcast) EndOfReportTimeTurn(reporterNickname string) {
	b.broadcast(reporterNickname)
}
