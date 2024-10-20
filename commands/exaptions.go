package commands

import (
	"fmt"
	"os"
)

func CommandNotFound(input string) string {
	value, isOk := Collections[input]
	if !isOk {
		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
		os.Exit(1)
	}

	return value
}
