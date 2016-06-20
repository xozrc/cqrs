package command

import (
	"errors"
	"reflect"

	"golang.org/x/net/context"

	cqrspkg "github.com/xozrc/cqrs/pkg"
	"github.com/xozrc/cqrs/types"
)

var (
	commandTypeMap map[string]reflect.Type
)

var (
	CommandTypeNoFound = errors.New("command type no found")
)

func init() {
	commandTypeMap = make(map[string]reflect.Type)
}

//Command
type Command interface {
	Id() types.Guid
}

func RegisterCommand(cmd Command) {
	key := cqrspkg.TypeName(cmd)
	typ := reflect.TypeOf(cmd).Elem()
	commandTypeMap[key] = typ
}

func NewCommand(key string) (cmd Command, err error) {
	typ, ok := commandTypeMap[key]
	if !ok {
		err = CommandTypeNoFound
		return
	}
	val := reflect.New(typ)

	cmd, ok = val.Interface().(Command)
	if !ok {
		err = errors.New("assert error")
		return
	}
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
