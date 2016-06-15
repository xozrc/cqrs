package main

import (
	"log"
    "github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/messaging"
	messagingnsq "github.com/xozrc/cqrs/messaging/nsq"
    "github.com/nsqio/go-nsq"
    "github.com/xozrc/cqrs/eventsourcing"
    orderconstant "github.com/xozrc/cqrs/eventsourcing/examples/order/constant"
    "golang.org/x/net/context"
)

const(
    orderCommandTopic = "command_order"
    orderCommandChannel = "command_order"
)

var (
    orderRepository *eventsourcing.EventSourcedRepository
	orderCommandReceiver messaging.Receiver
    orderCommandProcessor *command.CommandProcessor
    done chan struct{}
    
)


func initCommandReceiver() error{
    nsqConfig := nsq.NewConfig()
    c,err:=nsq.NewConsumer(orderCommandTopic,orderCommandChannel,nsqConfig)
    if err!=nil{
        return err
    }
    
    defer func(){
       if err!=nil{
           c.Stop()
       } 
    }()

    orderCommandReceiver, err = messagingnsq.NewConsumer(c,orderconstant.OrderCommandBusAddr)
    if err != nil {
        return err
    }
    
   
    return nil
}

func initCommandProcessor() {
    cd := command.NewCommandDispatcher()
    orderCommandProcessor = command.NewCommandProcessor(cd)
}

func Start() {
    done = make(chan struct{})
     
    err:=initCommandReceiver()
    
    if err!=nil{
        log.Fatal("init command receiver ",err)
    }
    
    initCommandProcessor()
    
    ctx := context.Background()
    ctx = eventsourcing.WithRepository(ctx,orderRepository)
    
    tp := messaging.NewProcessor(ctx,orderCommandProcessor)
	err = orderCommandReceiver.Start(tp)
    if err!=nil{
          log.Fatal(err)
    }
}

func Stop() {
	orderCommandReceiver.Stop()
}


func main(){
    Start()
    <-done
}

func exit(){
    t := struct{}{}
    done<-t
}