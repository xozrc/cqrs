package handler

import (
	"errors"
	"fmt"

	"github.com/xozrc/cqrs/command"
	"github.com/xozrc/cqrs/eventsourcing"
	"github.com/xozrc/cqrs/eventsourcing/examples/order"
	"github.com/xozrc/cqrs/types"
	"golang.org/x/net/context"
)

//HandleCreateOrder is a handler for create order command
func HandleCreateOrder(ctx context.Context, cmd command.Command) error {
	cmd, ok := cmd.(*CreateOrder)
	if !ok {
		return errors.New("")
	}

	//get the object which the command effect in
	repo := eventsourcing.RepositoryInContext(ctx)
	if repo == nil {
		return errors.New("")
	}

	id := types.NewGuid()
	order := order.NewOrder()
	err = order.Init(id)
	if err != nil {
		return err
	}

	cId := fmt.Sprintf("%d", cmd.Id())
	err = repo.Save(es, cId)
	if err != nil {
		return err
	}
	return nil
}
