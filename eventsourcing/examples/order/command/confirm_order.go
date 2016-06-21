package command

import (
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/types"
)

type ConfirmOrder struct {
	CId       types.Guid `json:"id"`
	OrderId   types.Guid `json:"order_id"`
	PaymentId types.Guid `json:"payment_id"`
}

func (co *ConfirmOrder) Id() types.Guid {
	return co.CId
}

func NewConfirmOrder(id types.Guid, orderId types.Guid, paymentId types.Guid) command.Command {
	return &ConfirmOrder{
		CId:       id,
		OrderId:   orderId,
		PaymentId: paymentId,
	}
}

func init() {
	command.RegisterCommand((*ConfirmOrder)(nil))
}
