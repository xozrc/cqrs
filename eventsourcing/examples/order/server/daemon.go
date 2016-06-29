package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"net/http"
	_ "net/http/pprof"

	"github.com/jinzhu/gorm"
	"github.com/nsqio/go-nsq"
	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/eventsourcing"
	ordercommand "github.com/xozrc/cqrs/eventsourcing/examples/order/command"
	orderconstant "github.com/xozrc/cqrs/eventsourcing/examples/order/constant"
	orderhandler "github.com/xozrc/cqrs/eventsourcing/examples/order/server/handler"
	"github.com/xozrc/cqrs/eventsourcing/rdb"
	"github.com/xozrc/cqrs/messaging"
	messagingnsq "github.com/xozrc/cqrs/messaging/nsq"
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"golang.org/x/net/context"
)

const (
	dialect  = "mysql"
	user     = "root"
	password = "root"
	host     = "127.0.0.1:3306"
	dbName   = "event_store_test"
	charset  = "utf8"
)

const (
	orderCommandTopic   = "command_order"
	orderCommandChannel = "command_order"
)

var (
	eventStore            eventsourcing.EventStore
	orderRepository       *eventsourcing.EventSourcedRepository
	orderCommandReceiver  messaging.Receiver
	orderCommandProcessor *command.CommandProcessor
	done                  chan struct{}
)

func initCommandReceiver() error {
	nsqConfig := nsq.NewConfig()

	nsqConfig.MaxInFlight = 2500
	c, err := nsq.NewConsumer(orderCommandTopic, orderCommandChannel, nsqConfig)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			c.Stop()
		}
	}()

	orderCommandReceiver, err = messagingnsq.NewConsumer(c, orderconstant.OrderCommandBusAddr)
	if err != nil {
		return err
	}

	return nil
}

func initCommandProcessor() {
	cd := command.NewCommandDispatcher()

	ct := cqrspkg.TypeName((*ordercommand.CreateOrder)(nil))

	cd.Register(ct, command.CommandHandlerFunc(orderhandler.HandleCreateOrder))
	ct2 := cqrspkg.TypeName((*ordercommand.CancelOrder)(nil))
	cd.Register(ct2, command.CommandHandlerFunc(orderhandler.HandleCancelOrder))

	orderCommandProcessor = command.NewCommandProcessor(cd)
}

func initEventStore() error {
	dbArgs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", user, password, host, dbName, charset)
	s, err := gorm.Open(dialect, dbArgs)
	if err != nil {
		return err
	}

	//s.LogMode(true)
	eventStore, err = rdb.NewStore(s)
	if err != nil {
		return err
	}
	return nil
}

func Start() {

	done = make(chan struct{})

	err := initCommandReceiver()

	if err != nil {
		log.Fatal("init command receiver ", err)
	}

	initCommandProcessor()
	err = initEventStore()

	if err != nil {
		log.Fatal("init event store ", err)
	}

	orderRepository = eventsourcing.NewRepository(eventStore)

	ctx := context.Background()
	ctx = eventsourcing.WithRepository(ctx, orderRepository)

	tp := messaging.NewProcessor(ctx, orderCommandProcessor)
	err = orderCommandReceiver.Start(tp)
	if err != nil {
		log.Fatal(err)
	}
}

func Stop() {
	orderCommandReceiver.Stop()
}

func main() {
	log.SetLevel(log.InfoLevel)

	Start()
	http.ListenAndServe(":3000", nil)
	<-done
}

func exit() {
	t := struct{}{}
	done <- t
}
