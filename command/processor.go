package command

import (
	"encoding/json"
	"log"

	"github.com/xozrc/cqrs/messaging"
	"golang.org/x/net/context"
)

//CommandProcessor implements messaging.MessageHandler
type CommandProcessor struct {
	cd CommandDispatcher //dispatcher
}

//Handle process message to command, then dispatch command to CommandHandler
func (cp *CommandProcessor) Handle(ctx context.Context, msg *messaging.Message) error {

	et := msg.MessageType

	c, err := NewCommand(et)
	if err != nil {
		return err
	}

	err = json.Unmarshal(msg.Payload, c)
	if err != nil {
		return err
	}

	log.Printf("receive command %#v\n", c)
	err = cp.cd.DispatchCommand(ctx, c)

	if err != nil {
		return err
	}

	return nil
}

//NewCommandProcessor constructor CommandProject object
func NewCommandProcessor(cd CommandDispatcher) *CommandProcessor {
	cp := &CommandProcessor{}
	cp.cd = cd
	return cp
}
