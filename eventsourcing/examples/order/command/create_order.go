package command

import (
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/types"
)

type CreateOrder struct {
	CId types.Guid `json:"id"`
}

func (co *CreateOrder) Id() types.Guid {
	return co.CId
}

func NewCreateOrder(id types.Guid) command.Command {
	return &CreateOrder{
		CId: id,
	}
}

func init() {
	command.RegisterCommand((*CreateOrder)(nil))
}
