package command

import (
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/types"
)

type CreateOrder struct {
	id types.Guid
}

func (co *CreateOrder) Id() types.Guid {
	return co.id
}

func NewCreateOrder(id types.Guid) command.Command {
	return &CreateOrder{
		id: id,
	}
}

func init() {
	command.RegisterCommand((*CreateOrder)(nil))
}
