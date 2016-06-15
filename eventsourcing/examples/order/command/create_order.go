package command

import "github.com/xozrc/cqrs/types"

type CreateOrder struct {
	id types.Guid
}

func (co *CreateOrder) Id() types.Guid {
	return co.id
}

func NewCreateOrder(id types.Guid) *CreateOrder {
	return &CreateOrder{
		id: id,
	}
}
