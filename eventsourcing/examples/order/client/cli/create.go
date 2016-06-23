package cli

import (
	"log"

	"github.com/codegangsta/cli"
)

import (
	orderclient "github.com/xozrc/cqrs/eventsourcing/examples/order/client/client"
)

var (
	createCmd = cli.Command{
		Name:        "create",
		Usage:       "create",
		Description: "create an order",
		Action:      create,
	}
)

func init() {
	appendCmd(createCmd)
}

func create(c *cli.Context) {

	client, err := orderclient.NewClient(bus, topic)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()
	err = client.CreateOrder()
	if err != nil {
		log.Fatalln(err)
	}
}
