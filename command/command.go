package command

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/xozrc/cqrs/types"
)

var (
	commandFactoryMap map[string]CommandFactory
)

var (
	CommandFactoryNoFound = errors.New("command factory no found")
)

func init() {
	commandFactoryMap = make(map[string]CommandFactory)
}

//Command
type Command interface {
	Id() types.Guid
}

type CommandFactory interface {
	NewCommand(id types.Guid) Command
}

type CommandFactoryFunc func(id types.Guid) Command

func (veff CommandFactoryFunc) NewCommand(id types.Guid) Command {
	return veff(id)
}

func RegisterCommand(key string, cf CommandFactory) {
	commandFactoryMap[key] = cf
}

func NewCommand(key string, id types.Guid) (cmd Command, err error) {
	factory, ok := commandFactoryMap[key]
	if !ok {
		err = CommandFactoryNoFound
		return
	}
	cmd = factory.NewCommand(id)
	return
}

//CommandHandler
type CommandHandler interface {
	HandleCommand(ctx context.Context, c Command) error
}

//CommandHandlerFunc implements CommandHandler
type CommandHandlerFunc func(ctx context.Context, c Command) error

//HandleCommand call self func
func (chf CommandHandlerFunc) HandleCommand(ctx context.Context, c Command) error {
	return chf(ctx, c)
}
