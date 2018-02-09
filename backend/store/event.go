package store

import "github.com/zumak/zumo/datatypes"

type EventEmitter struct {
	// Write Only
	PutMessage    chan<- PutMessageEvent
	PutChannel    chan<- datatypes.Channel
	DeleteChannel chan<- string
}
type EventReciever struct {
	// Read only
	PutMessage    <-chan PutMessageEvent
	PutChannel    <-chan datatypes.Channel
	DeleteChannel <-chan string
}

type PutMessageEvent struct {
	ChannelID string
	Message   datatypes.Message
}

func newEventSystem() (*EventEmitter, *EventReciever) {
	pm := make(chan PutMessageEvent)
	pc := make(chan datatypes.Channel)
	dc := make(chan string)

	return &EventEmitter{pm, pc, dc}, &EventReciever{pm, pc, dc}
}
