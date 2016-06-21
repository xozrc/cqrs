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
func HandleConfirmOrder(ctx context.Context, cmd command.Command) error {
	tcmd, ok := cmd.(*ordercommand.ConfirmOrder)
	if !ok {
		return errors.New("command error")
	}

	//get the object which the command effect in
	repo := eventsourcing.RepositoryInContext(ctx)
	if repo == nil {
		return errors.New("repo no exist")
	}

	order := ordereventsourcing.NewOrder()
	err := repo.Find(order)
	if err != nil {
		return err
	}
	err = order.Confirm(tcmd.PaymentId)
	if err != nil {
		return err
	}
	cId := fmt.Sprintf("%d", tcmd.Id())
	err = repo.Save(order, cId)
	if err != nil {
		return err
	}
	return nil
}
