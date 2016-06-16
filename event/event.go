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
	NewEvent() Event
}

type EventFactoryFunc func() Event

func (veff EventFactoryFunc) NewEvent() Event {
	return veff()
}

//trivival versioned event
type TrivialEvent struct {
	sourceId types.Guid `json:"source_id"`
}

func (te *TrivialEvent) SourceId() types.Guid {
	return te.sourceId
}

func NewEvent() Event {
	return &TrivialEvent{}
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
