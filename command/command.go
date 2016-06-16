package command

import (
	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
	"golang.org/x/net/context"
)

type Command interface {
	Id() types.Guid
}

type CommandHandler interface {
	HandleCommand(ctx context.Context, c Command) error
}

type CommandHandlerFunc func(ctx context.Context, c Command) error

func (chf CommandHandlerFunc) HandleCommand(ctx context.Context, c Command) error {
	return chf(ctx, c)
}

type CommandFactory interface {
	NewCommand(id types.Guid) Command
}

type CommandFactoryFunc func(id types.Guid) Command

func (veff CommandFactoryFunc) NewCommand(id types.Guid) Command {
	return veff(id)
}

//trivival command
type TrivialCommand struct {
	id types.Guid `json:"id"`
}

func (tc *TrivialCommand) Id() types.Guid {
	return tc.id
}

func NewCommand(id types.Guid) Command {
	return &TrivialCommand{
		id: id,
	}
}

var (
	commandFactoryMap map[string]CommandFactory
)

func init() {
	commandFactoryMap = make(map[string]CommandFactory)
	triKey := cqrspkg.TypeName((*TrivialCommand)(nil))
	RegisterCommandFactory(triKey, CommandFactoryFunc(NewCommand))
}

func RegisterCommandFactory(key string, vef CommandFactory) {
	commandFactoryMap[key] = vef
}

func GetCommandFactory(key string) CommandFactory {
	return commandFactoryMap[key]
}
