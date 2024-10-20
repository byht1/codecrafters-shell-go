package commands

type Runner interface {
	Run(params []string) error
	GetName() string
}

type AbstractCommand struct {
	name string
}

func (c *AbstractCommand) GetName() string {
	return c.name
}

var RunnerCollections []Runner

func init() {
	RunnerCollections = append(RunnerCollections, &ExitCommand{AbstractCommand{CLI_EXIT}})
}
