package eventsourcing

import (
	"reflect"

	"github.com/xozrc/cqrs/event"
	"github.com/xozrc/cqrs/types"
)

type VersionedEvent interface {
	event.Event
	Version() int64
}

//trivival versioned event
type TrivialVersionedEvent struct {
	sourceId types.Guid
	version  int64
}

func (tve *TrivialVersionedEvent) SourceId() types.Guid {
	return tve.sourceId
}

func (tve *TrivialVersionedEvent) Version() int64 {
	return tve.version
}

type VersionEventFactory interface {
	NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent
}

type VersionEventFactoryFunc func(sourceId types.Guid, version int64) VersionedEvent

func (veff VersionEventFactoryFunc) NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent {
	return veff(sourceId, version)
}

func NewVersionEvent(sourceId types.Guid, version int64) VersionedEvent {
	tve := &TrivialVersionedEvent{
		sourceId: sourceId,
		version:  version,
	}
	return tve
}

var (
	versionEventFactoryMap map[string]VersionEventFactory
)

func init() {
	versionEventFactoryMap = make(map[string]VersionEventFactory)

	triviKey := reflect.TypeOf((*TrivialVersionedEvent)(nil)).Elem().Name()
	RegisterVersionEventFactory(triviKey, VersionEventFactoryFunc(NewVersionEvent))
}

func RegisterVersionEventFactory(key string, vef VersionEventFactory) {
	versionEventFactoryMap[key] = vef
}

func GetVersionEventFactory(key string) VersionEventFactory {
	return versionEventFactoryMap[key]
}
