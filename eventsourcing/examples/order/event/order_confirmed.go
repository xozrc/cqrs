package event

import (
	"github.com/xozrc/cqrs/eventsourcing"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	eventsourcingtypes "github.com/xozrc/cqrs/types"
)

type OrderConfirmed struct {
	sourceId  eventsourcingtypes.Guid
	version   int64
	PaymentId eventsourcingtypes.Guid `json:"payment_id"`
}

func (ie *OrderConfirmed) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *OrderConfirmed) Version() int64 {
	return ie.version
}

func NewOrderConfirmed(sourceId eventsourcingtypes.Guid, version int64) eventsourcing.VersionedEvent {
	return &OrderConfirmed{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	orderConfirmedEventKey := cqrspkg.TypeName((*OrderConfirmed)(nil))
	eventsourcing.RegisterVersionEventFactory(orderConfirmedEventKey, eventsourcing.VersionEventFactoryFunc(NewOrderConfirmed))
}
