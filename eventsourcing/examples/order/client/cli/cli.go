package cli

import (
	"os"

	"github.com/codegangsta/cli"
	orderconstant "github.com/xozrc/cqrs/eventsourcing/examples/order/constant"
)

var (
	bus   string = orderconstant.OrderCommandBusAddr
	topic string = orderconstant.OrderCommandTopic
)

var (
	commands []cli.Command
)

func appendCmd(cmd cli.Command) {
	commands = append(commands, cmd)
}

func Start() {
	app := cli.NewApp()
	app.Name = "order client"
	app.Usage = "orderclient [global options] command [command options] [arguments...]."

	app.Author = ""
	app.Email = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "bus,b",
			Value:       bus,
			Usage:       "bus address for order client to publish",
			Destination: &bus,
		},
		cli.StringFlag{
			Name:        "topic,t",
			Value:       topic,
			Usage:       "topic for order client  to publish",
			Destination: &topic,
		},
	}

	app.Commands = commands
	app.Run(os.Args)
}
