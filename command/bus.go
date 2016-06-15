package command

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/xozrc/cqrs/messaging"
	"github.com/xozrc/cqrs/types"
)

type CommandBus interface {
	Publish(c Command) error
}

type CommandPublisher interface {
	Commands() []Command
}

type commandBus struct {
	sender messaging.Sender
}

func (cb *commandBus) Publish(c Command) error {
	m, err := buildMessage(c)
	if err != nil {
		return err
	}
	bs, err := m.MarshalBinary()
	if err != nil {
		return err
	}
	err = cb.sender.Send(bs)
	return err
}

func NewCommandBus(sender messaging.Sender) CommandBus {
	cb := &commandBus{
		sender: sender,
	}
	return cb
}

func buildMessage(cmd Command) (m *messaging.Message, err error) {
	m = &messaging.Message{}
	m.CorrelationId = fmt.Sprintf("%s", cmd.Id())
	m.Id = fmt.Sprintf("%s", types.NewGuid())
	m.MessageType = reflect.TypeOf(cmd).Name()
	bs, err := json.Marshal(cmd)
	if err != nil {
		return
	}
	m.Payload = bs
	return
}
