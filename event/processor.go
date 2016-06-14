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

func (ep *EventProcessor) Handle(ctx context.Context, msg messaging.Message) error {

	et := msg.MessageType

	factory := GetEventFactory(et)
	if factory == nil {
		return EventHandlerNoFound
	}

	e := factory.NewEvent()
	err := json.Unmarshal(msg.Payload, e)
	if err != nil {
		return EventHandlerNoFound
	}

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
