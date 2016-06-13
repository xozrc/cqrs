package event

import "golang.org/x/net/context"

type EventHandler interface {
	HandleEvent(ctx context.Context, e Event) error
}
