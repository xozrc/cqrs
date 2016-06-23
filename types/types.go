package types

import (
	"fmt"
	"math/rand"
	"time"
)

type Guid int64

func (guid Guid) String() string {
	return fmt.Sprintf("%d", guid)
}

func NewGuid() Guid {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := rand.Int63()
	return Guid(n)
}
