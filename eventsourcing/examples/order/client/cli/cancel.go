package cli

import (
	"log"
	"strconv"

	"github.com/codegangsta/cli"
)

import (
	orderclient "github.com/xozrc/cqrs/eventsourcing/examples/order/client/client"
	cqrstypes "github.com/xozrc/cqrs/types"
)

var (
	cancelCmd = cli.Command{
		Name:        "cancel",
		Usage:       "cancel order_id",
		Description: "cancel an order",
		Action:      cancel,
	}
)

func init() {
	appendCmd(cancelCmd)
}

func cancel(c *cli.Context) {
	//args no enough
	if c.NArg() < 1 {
		log.Fatalln("need order id")
	}
	id, err := strconv.ParseInt(c.Args()[0], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := orderclient.NewClient(bus, topic)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Stop()
	err = client.CancelOrder(cqrstypes.Guid(id))
	if err != nil {
		log.Fatalln(err)
	}
}
