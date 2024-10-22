package commands

import (
	"fmt"
	"os"
	"strings"
)

type EchoCommand struct {
	AbstractCommand
}

func (c *EchoCommand) Run(params []string) error {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(params, " "))
	return nil
}
