package event

import (
	"errors"
	"reflect"

	"golang.org/x/net/context"
)

var (
	EventHandlerNoFound = errors.New("event handler no found")
)

const (
	initEventHandlerSize = 10
)

type EventDispatcher interface {
	DispatchEvent(ctx context.Context, e Event) error
	Register(et string, ch EventHandler)
}

type eventDispatcher struct {
	handlersMap map[string]EventHandler
}

func (ed *eventDispatcher) DispatchEvent(ctx context.Context, e Event) error {
	et := reflect.TypeOf(e).Name()

	h, ok := ed.handlersMap[et]
	if !ok {
		return EventHandlerNoFound
	}
	return h.HandleEvent(ctx, e)
}

func (cd *eventDispatcher) Register(et string, eh EventHandler) {
	cd.handlersMap[et] = eh
}

func NewEventDispatcher() EventDispatcher {
	cd := &eventDispatcher{}
	cd.handlersMap = make(map[string]EventHandler, initEventHandlerSize)
	return cd
}
