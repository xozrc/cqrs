package command

import (
	"github.com/xozrc/cqrs/command"
	cqrspkg "github.com/xozrc/cqrs/pkg"
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
	createOrderTyp := cqrspkg.TypeName((*CreateOrder)(nil))
	command.RegisterCommandFactory(createOrderTyp, command.CommandFactoryFunc(NewCreateOrder))
}
