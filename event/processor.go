package event

import (
	"encoding/json"

	"github.com/xozrc/cqrs/messaging"
	"github.com/xozrc/cqrs/types"
	"golang.org/x/net/context"
)

type EventProcessor struct {
	ed EventDispatcher
}

func (ep *EventProcessor) Handle(ctx context.Context, msg *messaging.Message) error {

	et := msg.MessageType

	factory := GetEventFactory(et)
	if factory == nil {
		return EventHandlerNoFound
	}
	id := types.NewGuid()
	e := factory.NewEvent(id)
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
