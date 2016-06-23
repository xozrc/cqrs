package event

import (
	"github.com/xozrc/cqrs/eventsourcing"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	eventsourcingtypes "github.com/xozrc/cqrs/types"
)

type OrderReserved struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *OrderReserved) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *OrderReserved) Version() int64 {
	return ie.version
}

func NewOrderReserved(sourceId eventsourcingtypes.Guid, version int64) eventsourcing.VersionedEvent {
	return &OrderInit{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	orderReservedEventKey := cqrspkg.TypeName((*OrderReserved)(nil))
	eventsourcing.RegisterVersionEventFactory(orderReservedEventKey, eventsourcing.VersionEventFactoryFunc(NewOrderReserved))
}
