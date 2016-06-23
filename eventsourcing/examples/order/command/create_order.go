package command

import (
	"github.com/xozrc/cqrs/command"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

type CreateOrder struct {
	id      types.Guid `json:"id"`
	OrderId types.Guid `json:"order_id"`
}

func (co *CreateOrder) Id() types.Guid {
	return co.id
}

func NewCreateOrder(id types.Guid) command.Command {
	return &CreateOrder{
		id: id,
	}
}

func NewCreateOrder1(id types.Guid, orderId types.Guid) command.Command {
	return &CreateOrder{
		id:      id,
		OrderId: orderId,
	}
}

func init() {
	key := cqrspkg.TypeName((*CreateOrder)(nil))
	command.RegisterCommand(key, command.CommandFactoryFunc(NewCreateOrder))
}
