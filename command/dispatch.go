package command

import (
	"errors"
	"reflect"

	"golang.org/x/net/context"
)

var (
	CommandHandlerNoFound = errors.New("event handler no found")
)

const (
	initCommandHandlerSize = 10
)

type CommandDispatcher interface {
	DispatchCommand(ctx context.Context, c Command) error
	Register(et string, ch CommandHandler)
}

type commandDispatcher struct {
	handlersMap map[string]CommandHandler
}

func (cd *commandDispatcher) DispatchCommand(ctx context.Context, c Command) error {
	et := reflect.TypeOf(c).Name()

	h, ok := cd.handlersMap[et]
	if !ok {
		return CommandHandlerNoFound
	}
	return h.HandleCommand(ctx, c)
}

func (cd *commandDispatcher) Register(et string, ch CommandHandler) {
	cd.handlersMap[et] = ch
}

func NewCommandDispatcher() CommandDispatcher {
	cd := &commandDispatcher{}
	cd.handlersMap = make(map[string]CommandHandler, initCommandHandlerSize)
	return cd
}
