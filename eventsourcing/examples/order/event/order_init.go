package event

import (
	"reflect"

	"github.com/xozrc/cqrs/eventsourcing"
	eventsourcingtypes "github.com/xozrc/cqrs/types"
)

type OrderInit struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *OrderInit) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *OrderInit) Version() int64 {
	return ie.version
}

func NewOrderInit(sourceId eventsourcingtypes.Guid, version int64) eventsourcing.VersionedEvent {
	return &OrderInit{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	initEventKey := reflect.TypeOf((*OrderInit)(nil)).Name()
	eventsourcing.RegisterVersionEventFactory(initEventKey, eventsourcing.VersionEventFactoryFunc(NewOrderInit))
}