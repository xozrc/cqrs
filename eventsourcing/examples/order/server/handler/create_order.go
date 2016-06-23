package handler

import (
	"errors"
	"fmt"

	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/eventsourcing"
	ordercommand "github.com/xozrc/cqrs/eventsourcing/examples/order/command"
	ordereventsourcing "github.com/xozrc/cqrs/eventsourcing/examples/order/server/eventsourcing"
	"golang.org/x/net/context"
)

//HandleCreateOrder is a handler for create order command
func HandleCreateOrder(ctx context.Context, cmd command.Command) error {
	cmd1, ok := cmd.(*ordercommand.CreateOrder)
	if !ok {
		return errors.New("command error")
	}

	//get the object which the command effect in
	repo := eventsourcing.RepositoryInContext(ctx)
	if repo == nil {
		return errors.New("repo no exist")
	}
	order := ordereventsourcing.NewOrder()
	//check if order already exist
	err := repo.Find(cmd1.OrderId, order)
	if err != nil && err != eventsourcing.EventSourcedNoFound {
		return err
	}

	err = order.Init(cmd1.OrderId)
	if err != nil {
		return err
	}

	cId := fmt.Sprintf("%d", cmd.Id())
	err = repo.Save(order, cId)
	if err != nil {
		return err
	}
	return nil
}
