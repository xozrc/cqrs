package nsq

import (
	log "github.com/Sirupsen/logrus"

	"runtime"

	"github.com/nsqio/go-nsq"
	"github.com/xozrc/cqrs/messaging"
)

type NSQConsumer struct {
	consumer    *nsq.Consumer
	addr        string
	receiveChan chan []byte
	handler     messaging.Handler
}

func (nsqc *NSQConsumer) Start(h messaging.Handler) error {
	err := nsqc.consumer.ConnectToNSQD(nsqc.addr)
	if err != nil {
		return err
	}
	nsqc.handler = h
	return nil
}

func (nsqc *NSQConsumer) Stop() error {
	nsqc.consumer.Stop()
	return nil
}

func (nsqc *NSQConsumer) HandleMessage(msg *nsq.Message) error {
	return nil
	log.Debugf("receive msg %v\n", msg)
	err := nsqc.handler.Handle(msg.Body)
	if err != nil {
		return err
	}
	return nil
}

func NewConsumer(c *nsq.Consumer, addr string) (nsqc *NSQConsumer, err error) {
	nsqc = &NSQConsumer{}
	nsqc.consumer = c
	nsqc.addr = addr
	n := runtime.GOMAXPROCS(0)
	nsqc.consumer.AddConcurrentHandlers(nsqc, n)
	return
}
