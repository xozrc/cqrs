package version

import (
	"github.com/xozrc/cqrs/event"
)

type VersionedEvent interface {
	event.Event
	Version() int64
}
