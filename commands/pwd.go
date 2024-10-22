package commands

import (
	"fmt"
	"os"
)

type PwdCommand struct {
	AbstractCommand
}

func (c *PwdCommand) Run(params []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)

	return nil
}
