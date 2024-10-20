package commands

import (
	"os"
	"strconv"
)

type ExitCommand struct {
	AbstractCommand
}

func (c *ExitCommand) Run(params []string) error {
	num, err := strconv.Atoi(params[0])
	if err != nil {
		os.Exit(0)
	}
	os.Exit(num)

	return nil
}
