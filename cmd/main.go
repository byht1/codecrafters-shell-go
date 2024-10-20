package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/commands"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")

		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		command, params, isOk := commands.ProcessInputData(text)
		if !isOk {
			continue
		}

		command.Run(params)
	}
}
