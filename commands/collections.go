package commands

type CommandCollection map[CommandName]Runner

const (
	SHELL_TYPE            CommandType = "shell"
	EXECUTABLE_FILES_TYPE CommandType = "executableFiles"
)

const (
	CLI_EXIT     CommandName = "exit"
	CLI_ECHO     CommandName = "echo"
	CLI_TYPE     CommandName = "type"
	DEFAULT_PATH string      = "default"
)

var collections *CommandCollection

func SingletonCommandCollection() CommandCollection {

	if collections != nil {
		return *collections
	}

	obj := make(CommandCollection)

	obj[CLI_EXIT] = &ExitCommand{AbstractCommand{CLI_EXIT, SHELL_TYPE, DEFAULT_PATH}}
	obj[CLI_ECHO] = &EchoCommand{AbstractCommand{CLI_ECHO, SHELL_TYPE, DEFAULT_PATH}}
	obj[CLI_TYPE] = &TypeCommand{AbstractCommand{CLI_TYPE, SHELL_TYPE, DEFAULT_PATH}}

	collections = &obj

	return *collections
}
