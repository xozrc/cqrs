package eventsourcing

import (
	"github.com/xozrc/cqrs/types"
)

type EventSourced interface {
	Id() types.Guid
	Version() int64
	ApplyEvent(ve VersionedEvent) error
	Events() []VersionedEvent
	Payload() []byte
}
