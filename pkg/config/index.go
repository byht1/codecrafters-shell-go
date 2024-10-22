package config

import (
	"os"
	"path"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/commands"
	"github.com/codecrafters-io/shell-starter-go/pkg/helpers"
)

func ProcessEnvPath() {
	pathValue := helpers.GetEnv("PATH")

	commandCollection := commands.SingletonCommandCollection()
	paths := strings.Split(pathValue, ":")
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			if os.IsNotExist(err) {
				continue
			}
		}

		files, err := os.ReadDir(p)
		if err != nil {
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}

			commandName := commands.CommandName(f.Name())
			if _, isOk := commandCollection[commandName]; isOk {
				continue
			}

			pathToExecutableFile := path.Join(p, f.Name())
			executedCommand := commands.NewExecutedCommand(commandName, pathToExecutableFile)
			commandCollection[commandName] = &executedCommand
		}
	}
}
