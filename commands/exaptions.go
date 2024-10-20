package commands

import (
	"fmt"
	"os"
	"strings"
)

func ProcessInputData(input string) (string, []string, bool) {
	splitInput := strings.Split(input, " ")
	command := splitInput[0]

	value, isOk := Collections[command]
	if !isOk {
		fmt.Fprint(os.Stdout, strings.TrimSpace(command)+": command not found\n")
	}

	return value, splitInput[1:], isOk
}
