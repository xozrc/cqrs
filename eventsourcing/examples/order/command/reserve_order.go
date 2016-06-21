package command

import (
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/types"
)

type ReserveOrder struct {
	CId     types.Guid `json:"id"`
	OrderId types.Guid `json:"order_id"`
}

func (co *ReserveOrder) Id() types.Guid {
	return co.CId
}

func NewReserveOrder(id types.Guid, orderId types.Guid) command.Command {
	return &ReserveOrder{
		CId:     id,
		OrderId: orderId,
	}
}

func init() {
	command.RegisterCommand((*ReserveOrder)(nil))
}
