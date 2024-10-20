package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/commands"
)

// var _ = fmt.Fprint

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

		for _, item := range commands.RunnerCollections {
			if item.GetName() == command {
				item.Run(params)
			}
		}

	}
}
