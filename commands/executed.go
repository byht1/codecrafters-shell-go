package commands

import (
	"fmt"
	"os"
	"os/exec"
)

type ExecutedCommand struct {
	AbstractCommand
}

func NewExecutedCommand(name CommandName, pathToFile string) ExecutedCommand {
	return ExecutedCommand{AbstractCommand{name, EXECUTABLE_FILES_TYPE, pathToFile}}
}

func (c *ExecutedCommand) Run(params []string) error {
	cmd := exec.Command(c.pathToExecutableFile, params...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", c.name)
	}

	return nil
}
