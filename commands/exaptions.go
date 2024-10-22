package commands

import (
	"fmt"
	"os"
	"strings"
)

func ProcessInputData(input string) (Runner, []string, bool) {
	splitInput := strings.Split(strings.TrimSpace(input), " ")
	command := CommandName(splitInput[0])
	commandCollection := SingletonCommandCollection()

	value, isOk := commandCollection[command]
	if !isOk {
		fmt.Fprint(os.Stdout, strings.TrimSpace(splitInput[0])+": command not found\n")
	}

	return value, splitInput[1:], isOk
}
