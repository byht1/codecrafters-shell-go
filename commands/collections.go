package commands

const (
	CLI_EXIT = "exit"
)

var Collections map[string]string

func init() {
	Collections = make(map[string]string)

	Collections[CLI_EXIT] = CLI_EXIT
}
