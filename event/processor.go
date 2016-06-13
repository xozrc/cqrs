package event

import (
	"encoding/json"

	"golang.org/x/net/context"
)

import (
	"github.com/xozrc/cqrs/messaging"
)

type EventProcessor struct {
	ed EventDispatcher
}

func (ep *EventProcessor) Handle(msg []byte) error {
	m := &messaging.Message{}
	err := json.Unmarshal(msg, m)
	if err != nil {
		return err
	}

	et := m.MessageType

	factory := GetEventFactory(et)
	if factory == nil {
		return EventHandlerNoFound
	}

	e := factory.NewEvent()
	err = json.Unmarshal(m.Payload, e)
	if err != nil {
		return EventHandlerNoFound
	}
	//todo: time out context
	ctx := context.Background()
	err = ep.ed.DispatchEvent(ctx, e)

	if err != nil {
		return err
	}

	return nil
}

func NewEventProcessor(ed EventDispatcher) *EventProcessor {
	ep := &EventProcessor{}
	ep.ed = ed
	return ep
}
