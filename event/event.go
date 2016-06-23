package event

import (
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

//Event ...
type Event interface {
	SourceId() types.Guid
}

type EventFactory interface {
	NewEvent(id types.Guid) Event
}

type EventFactoryFunc func(id types.Guid) Event

func (veff EventFactoryFunc) NewEvent(id types.Guid) Event {
	return veff(id)
}

//trivival versioned event
type TrivialEvent struct {
	sourceId types.Guid
}

func (te *TrivialEvent) SourceId() types.Guid {
	return te.sourceId
}

func NewEvent(id types.Guid) Event {
	return &TrivialEvent{
		sourceId: id,
	}
}

var (
	eventFactoryMap map[string]EventFactory
)

func init() {
	eventFactoryMap = make(map[string]EventFactory)
	triKey := cqrspkg.TypeName((*TrivialEvent)(nil))
	RegisterEventFactory(triKey, EventFactoryFunc(NewEvent))

}

func RegisterEventFactory(key string, vef EventFactory) {
	eventFactoryMap[key] = vef
}

func GetEventFactory(key string) EventFactory {
	return eventFactoryMap[key]
}
