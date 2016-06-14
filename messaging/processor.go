package messaging

import (
	"golang.org/x/net/context"
)

type MessageHandler interface {
	Handle(ctx context.Context, msg *Message) error
}

type Processor struct {
	mh MessageHandler
}

func (p *Processor) Handle(msg []byte) error {
	m := &Message{}
	err := m.UnmarshalBinary(msg)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = p.mh.Handle(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func NewProcessor(mh MessageHandler) *Processor {
	p := &Processor{}
	p.mh = mh
	return p
}
