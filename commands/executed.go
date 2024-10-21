package commands

import "fmt"

type ExecutedCommand struct {
	AbstractCommand
}

func NewExecutedCommand(name CommandName, pathToFile string) ExecutedCommand {
	return ExecutedCommand{AbstractCommand{name, EXECUTABLE_FILES_TYPE, pathToFile}}
}

func (c *ExecutedCommand) Run(params []string) error {
	fmt.Println(c.pathToExecutableFile)

	return nil
}
