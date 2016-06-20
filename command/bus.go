package command

import (
	"encoding/json"
	"fmt"
	"log"
)

import(
	"github.com/xozrc/cqrs/messaging"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

//Command Bus
type CommandBus interface {
	Publish(c Command) error
}

//Command Publisher
type CommandPublisher interface {
	Commands() []Command
}

//Command Bus implementation
type commandBus struct {
	sender messaging.Sender
}

//Publish
func (cb *commandBus) Publish(c Command) error {
	m, err := buildMessage(c)
	if err != nil {
		return err
	}

	log.Printf("build message %#v\n", m)

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

//buildMessage
func buildMessage(cmd Command) (m *messaging.Message, err error) {
	m = &messaging.Message{}
	m.CorrelationId = fmt.Sprintf("%s", cmd.Id())
	m.Id = fmt.Sprintf("%s", types.NewGuid())
	m.MessageType = cqrspkg.TypeName(cmd)
	bs, err := json.Marshal(cmd)
	if err != nil {
		return
	}
	m.Payload = bs
	return
}
