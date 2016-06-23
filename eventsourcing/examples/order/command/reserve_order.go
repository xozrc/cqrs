package command

import (
	"github.com/xozrc/cqrs/command"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

type ReserveOrder struct {
	id      types.Guid `json:"id"`
	OrderId types.Guid `json:"order_id"`
}

func (co *ReserveOrder) Id() types.Guid {
	return co.id
}

func NewReserveOrder(id types.Guid) command.Command {
	return &ReserveOrder{
		id: id,
	}
}

func NewReserveOrder1(id types.Guid, orderId types.Guid) command.Command {
	return &ReserveOrder{
		id:      id,
		OrderId: orderId,
	}
}

func init() {
	key := cqrspkg.TypeName((*ReserveOrder)(nil))
	command.RegisterCommand(key, command.CommandFactoryFunc(NewReserveOrder))
}
