package main

import (
	"log"

	"github.com/nsqio/go-nsq"
	"github.com/xozrc/cqrs/command"
	ordercommand "github.com/xozrc/cqrs/eventsourcing/examples/order/command"
	orderconstant "github.com/xozrc/cqrs/eventsourcing/examples/order/constant"
	"github.com/xozrc/cqrs/messaging"
	messagingnsq "github.com/xozrc/cqrs/messaging/nsq"
	"github.com/xozrc/cqrs/types"
)

var (
	bus      command.CommandBus
	producer messaging.Sender
)

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
	}

	err = CreateOrder()
	if err != nil {
		log.Fatal(err)
	}

	err = destroy()
	log.Println("destroy error:", err)
}

func setup() error {
	tp, err := nsq.NewProducer(orderconstant.OrderCommandBusAddr, nsq.NewConfig())
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tp.Stop()
		}
	}()

	producer, err = messagingnsq.NewProducer(tp, orderconstant.OrderCommandTopic)
	if err != nil {
		return err
	}
	log.Println("setup producer success")
	bus = command.NewCommandBus(producer)
	return nil
}

func destroy() error {
	err := producer.Close()
	if err != nil {
		return err
	}
	return nil
}

func CreateOrder() error {
	tId := types.NewGuid()
	order := ordercommand.NewCreateOrder(tId)
	err := bus.Publish(order)
	if err != nil {
		return err
	}
	return nil
}
