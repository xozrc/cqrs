package aggregate

import "github.com/xozrc/cqrs/types"

type AggregateRoot interface {
	Id() types.Guid
}
