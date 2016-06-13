package types

import (
	"fmt"
	"math/rand"
)

type Guid int64

func (guid Guid) String() string {
	return fmt.Sprintf("%d", guid)
}

func NewGuid() Guid {
	n := rand.Int63()
	return Guid(n)
}
