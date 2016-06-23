package command

import (
	"github.com/xozrc/cqrs/command"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

type CancelOrder struct {
	id      types.Guid
	OrderId types.Guid `json:"order_id"`
}

func (co *CancelOrder) Id() types.Guid {
	return co.id
}

func NewCancelOrder(id types.Guid) command.Command {
	return &CancelOrder{
		id: id,
	}
}
func NewCancelOrder1(id types.Guid, orderId types.Guid) command.Command {
	return &CancelOrder{
		id:      id,
		OrderId: orderId,
	}
}

func init() {
	key := cqrspkg.TypeName((*CancelOrder)(nil))
	command.RegisterCommand(key, command.CommandFactoryFunc(NewCancelOrder))
}
