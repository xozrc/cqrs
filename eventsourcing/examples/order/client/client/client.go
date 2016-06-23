package context

import (
	"github.com/nsqio/go-nsq"
	"github.com/xozrc/cqrs/command"
	ordercommand "github.com/xozrc/cqrs/eventsourcing/examples/order/command"
	messagingnsq "github.com/xozrc/cqrs/messaging/nsq"
	cqrstypes "github.com/xozrc/cqrs/types"
)

type Client interface {
	CreateOrder() error
	CancelOrder(id cqrstypes.Guid) error
	Stop() error
}

type client struct {
	producer *nsq.Producer
	bus      command.CommandBus
}

func (c *client) CreateOrder() error {
	tId := cqrstypes.NewGuid()
	oId := cqrstypes.NewGuid()
	order := ordercommand.NewCreateOrder1(tId, oId)
	err := c.bus.Publish(order)
	return err
}

func (c *client) CancelOrder(id cqrstypes.Guid) error {
	tId := cqrstypes.NewGuid()
	order := ordercommand.NewCancelOrder1(tId, id)
	err := c.bus.Publish(order)
	return err
}

func (c *client) Stop() error {
	c.producer.Stop()
	return nil
}

func NewClient(addr string, topic string) (c Client, err error) {
	tp, err := nsq.NewProducer(addr, nsq.NewConfig())
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tp.Stop()
		}
	}()

	producer, err := messagingnsq.NewProducer(tp, topic)
	if err != nil {
		return
	}

	bus := command.NewCommandBus(producer)
	c = &client{
		bus:      bus,
		producer: tp,
	}
	return
}
