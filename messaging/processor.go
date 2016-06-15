package messaging

import (
	"log"

	"golang.org/x/net/context"
)

type MessageHandler interface {
	Handle(ctx context.Context, msg *Message) error
}

type Processor struct {
	ctx context.Context
	mh  MessageHandler
}

func (p *Processor) Handle(msg []byte) error {
	m := &Message{}
	err := m.UnmarshalBinary(msg)
	if err != nil {
		return err
	}
	log.Printf("receive msg %#v\n", m)
	err = p.mh.Handle(p.ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func NewProcessor(ctx context.Context, mh MessageHandler) *Processor {
	p := &Processor{}
	p.ctx = ctx
	p.mh = mh
	return p
}
