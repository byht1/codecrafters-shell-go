package commands

const (
	CLI_EXIT = "exit"
	CLI_ECHO = "echo"
)

var Collections map[string]string

func init() {
	Collections = make(map[string]string)

	Collections[CLI_EXIT] = CLI_EXIT
	Collections[CLI_ECHO] = CLI_ECHO
}
