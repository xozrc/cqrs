package aggregate

import "github.com/xozrc/cqrs/types"

//AggregateRoot
type AggregateRoot interface {
	Id() types.Guid
}
