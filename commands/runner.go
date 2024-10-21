package commands

type CommandName string
type CommandType string

type Runner interface {
	Run(params []string) error
	GetName() CommandName
	GetType() CommandType
	GetPathToExecutableFile() string
}

type AbstractCommand struct {
	name                 CommandName
	typeCommand          CommandType
	pathToExecutableFile string
}

func (c *AbstractCommand) GetName() CommandName {
	return c.name
}

func (c *AbstractCommand) GetType() CommandType {
	return c.typeCommand
}

func (c *AbstractCommand) GetPathToExecutableFile() string {
	return c.pathToExecutableFile
}
