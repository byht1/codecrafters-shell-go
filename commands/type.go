package commands

import (
	"fmt"
	"strings"
)

type TypeCommand struct {
	AbstractCommand
}

func (c *TypeCommand) Run(params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("pass at least one parameter")
	}

	searchCommand := strings.TrimSpace(params[0])
	commandCollection := SingletonCommandCollection()
	command, isOk := commandCollection[CommandName(searchCommand)]
	if !isOk {
		c.commandNotFound(searchCommand)
		return nil
	}

	commandType := command.GetType()
	commandName := command.GetName()

	switch commandType {
	case SHELL_TYPE:
		fmt.Printf("%v is a shell builtin\n", commandName)
	case EXECUTABLE_FILES_TYPE:
		fmt.Printf("%v is %v\n", commandName, "path")
	default:
		c.commandNotFound(searchCommand)
	}

	return nil
}

func (c *TypeCommand) commandNotFound(value string) {
	fmt.Printf("%v: not found\n", value)
}
