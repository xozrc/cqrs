package event

import (
	"reflect"

	"github.com/xozrc/eventsourcing/event"
	eventsourcingtypes "github.com/xozrc/eventsourcing/types"
)

type OrderInit struct {
	sourceId eventsourcingtypes.Guid
	version  int64
}

func (ie *OrderInit) SourceId() eventsourcingtypes.Guid {
	return ie.sourceId
}

func (ie *ConfirmEvent) Version() int64 {
	return ie.version
}

func NewConfirmEvent(sourceId eventsourcingtypes.Guid, version int64) event.VersionedEvent {
	return &CancelEvent{
		sourceId: sourceId,
		version:  version,
	}
}

func init() {
	confirmEventKey := reflect.TypeOf((*ConfirmEvent)(nil)).Name()
	event.RegisterVersionEventFactory(confirmEventKey, event.VersionEventFactoryFunc(NewConfirmEvent))
}
