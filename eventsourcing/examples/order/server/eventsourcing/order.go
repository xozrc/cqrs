package eventsourcing

import (
	"errors"

	"github.com/xozrc/cqrs/eventsourcing"
	orderevent "github.com/xozrc/cqrs/eventsourcing/examples/order/event"
	"github.com/xozrc/cqrs/types"
)

type OrderStatus int

const (
	OrderInit OrderStatus = iota
	OrderReserved
	OrderConfirmed
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
	case *orderevent.OrderInit:
		{
			e, ok := ve.(*orderevent.OrderInit)
			if !ok {
				return errors.New("null order init")
			}
			o.init(e.SourceId())
			goto Version
		}
	case *orderevent.OrderCancelled:
		{
			_, ok := ve.(*orderevent.OrderCancelled)
			if !ok {
				return errors.New("null order cancel")
			}
			o.cancel()
			goto Version
		}
	}
Version:
	o.version = ve.Version()
	return nil
}

func (o *Order) init(id types.Guid) {
	o.id = id
	o.Status = OrderInit
}

func (o *Order) cancel() {
	o.Status = OrderCancelled
}

func (o *Order) Events() []eventsourcing.VersionedEvent {
	return o.pendingEvents
}

func (o *Order) Payload() []byte {
	return nil
}

func (o *Order) Cancel() error {
	if o.Status == OrderCancelled {
		return nil
	}
	ver := o.version + 1
	ie := orderevent.NewOrderCancelled(o.Id(), ver)
	return o.updateEvent(ie)
}

func (o *Order) Init(id types.Guid) error {
	ver := o.version + 1
	ie := orderevent.NewOrderInit(id, ver)
	return o.updateEvent(ie)
}

func (o *Order) Reserve() error {
	ver := o.version + 1
	ie := orderevent.NewOrderReserved(o.Id(), ver)
	return o.updateEvent(ie)
}

func (o *Order) Confirm(paymentId types.Guid) error {
	ver := o.version + 1
	ie := orderevent.NewOrderConfirmed(o.Id(), ver)
	tIe, _ := ie.(*orderevent.OrderConfirmed)
	tIe.PaymentId = paymentId
	return o.updateEvent(tIe)
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
