package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/pkg/helpers"
)

type CdCommand struct {
	AbstractCommand
}

func (c *CdCommand) Run(params []string) error {
	movePath := ""
	if len(params) != 0 {
		movePath = strings.TrimSpace(params[0])
	}

	if movePath == "~" || movePath == "" {
		movePath = helpers.GetEnv("HOME")
	}

	err := os.Chdir(movePath)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v: %s: No such file or directory\n", CLI_CD, movePath)
	}

	return nil
}
