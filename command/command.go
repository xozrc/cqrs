package command

import (
	"errors"

	"github.com/xozrc/cqrs/types"
	"golang.org/x/net/context"
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

func RegisterCommandFactory(key string, vef CommandFactory) {
	commandFactoryMap[key] = vef
}

func GetCommandFactory(key string) CommandFactory {
	return commandFactoryMap[key]
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
