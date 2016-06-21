package command

import (
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/types"
)

type CancelOrder struct {
	CId     types.Guid `json:"id"`
	OrderId types.Guid `json:"order_id"`
}

func (co *CancelOrder) Id() types.Guid {
	return co.CId
}

func NewCancelOrder(id types.Guid, orderId types.Guid) command.Command {
	return &CancelOrder{
		CId:     id,
		OrderId: orderId,
	}
}

func init() {
	command.RegisterCommand((*CancelOrder)(nil))
}
