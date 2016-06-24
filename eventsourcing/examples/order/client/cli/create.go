package cli

import (
	"log"

	"github.com/codegangsta/cli"
)

import (
	orderclient "github.com/xozrc/cqrs/eventsourcing/examples/order/client/client"
)

var (
	repeat = 1
)

var (
	createCmd = cli.Command{
		Name:        "create",
		Usage:       "create",
		Description: "create an order",
		Action:      create,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "repeat,r",
				Value:       repeat,
				Usage:       "create repeat",
				Destination: &repeat,
			},
		},
	}
)

func init() {
	appendCmd(createCmd)
}

func create(c *cli.Context) {

	if repeat < 0 {
		repeat = 0
	}

	client, err := orderclient.NewClient(bus, topic)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()

	for i := 0; i < repeat || repeat == 0; i++ {
		err = client.CreateOrder()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
