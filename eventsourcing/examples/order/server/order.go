package order

import (
	"github.com/xozrc/cqrs/eventsourcing"
	orderevent "github.com/xozrc/cqrs/eventsourcing/example/event"
	"github.com/xozrc/cqrs/types"
)

type OrderStatus int

const (
	OrderInit OrderStatus = iota
	OrderPending
	OrderConfirm
	OrderVerify
	OrderCancelled
)

//order object
type Order struct {
	id            types.Guid
	version       int64
	pendingEvents []eventsourcing.VersionedEvent
	Status        OrderStatus
}

func (o *Order) Id() types.Guid {
	return o.id
}

func (o *Order) Version() int64 {
	return o.version
}

func (o *Order) ApplyEvent(ve eventsourcing.VersionedEvent) error {
	switch ve.(type) {
	case OrderInit:
		o.Status = OrderInit
		return nil
	}
	return nil
}

func (o *Order) Events() []eventsourcing.VersionedEvent {
	return o.pendingEvents
}

func (o *Order) Payload() []byte {
	return nil
}

func (o *Order) Cancel() error {
	o.Status = OrderCancelled
	return nil
}

func (o *Order) Pending() error {
	o.Status = OrderPending
	return nil
}

func (o *Order) Verify() error {
	o.Status = OrderVerify
	return nil
}

func (o *Order) Init(id types.Guid) error {
	ver := o.version + 1
	ie := orderevent.NewOrderInit(id, ver)
	return o.updateEvent(ie)
}

func (o *Order) updateEvent(ve eventsourcing.VersionedEvent) error {
	err := o.ApplyEvent(ve)
	if err != nil {
		return err
	}
	o.pendingEvents = append(o.pendingEvents, ve)
	return nil
}

func NewOrder() *Order {
	o := &Order{}
	return o
}
