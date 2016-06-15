package event

import (
	"reflect"

	"github.com/xozrc/cqrs/eventsourcing"
	eventsourcingtypes "github.com/xozrc/cqrs/types"
)

type OrderCancelled struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *OrderCancelled) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *OrderCancelled) Version() int64 {
	return ie.version
}

func NewOrderCancelled(sourceId eventsourcingtypes.Guid, version int64) eventsourcing.VersionedEvent {
	return &OrderCancelled{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	cancelEventKey := reflect.TypeOf((*OrderCancelled)(nil)).Name()
	eventsourcing.RegisterVersionEventFactory(cancelEventKey, eventsourcing.VersionEventFactoryFunc(NewOrderCancelled))
}
