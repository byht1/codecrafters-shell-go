package commands

import (
	"fmt"
	"os"
	"strings"
)

type CdCommand struct {
	AbstractCommand
}

func (c *CdCommand) Run(params []string) error {
	err := os.Chdir(params[0])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v: %s: No such file or directory\n", CLI_CD, strings.Join(params, " "))
	}

	return nil
}
