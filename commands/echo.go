package commands

import (
	"fmt"
	"strings"
)

type EchoCommand struct {
	AbstractCommand
}

func (c *EchoCommand) Run(params []string) error {
	fmt.Print(strings.Join(params, " "))

	return nil
}
