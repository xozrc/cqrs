package command

import (
	"github.com/xozrc/cqrs/command"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

type ConfirmOrder struct {
	id        types.Guid
	OrderId   types.Guid `json:"order_id"`
	PaymentId types.Guid `json:"payment_id"`
}

func (co *ConfirmOrder) Id() types.Guid {
	return co.id
}

func NewConfirmOrder(id types.Guid) command.Command {
	return &ConfirmOrder{
		id: id,
	}
}

func NewConfirmOrder1(id types.Guid, orderId types.Guid, paymentId types.Guid) command.Command {
	return &ConfirmOrder{
		id:        id,
		OrderId:   orderId,
		PaymentId: paymentId,
	}
}

func init() {
	key := cqrspkg.TypeName((*ConfirmOrder)(nil))
	command.RegisterCommand(key, command.CommandFactoryFunc(NewConfirmOrder))
}
