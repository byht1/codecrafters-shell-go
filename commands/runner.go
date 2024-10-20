package commands

type CommandName string
type CommandType string

type Runner interface {
	Run(params []string) error
	GetName() CommandName
	GetType() CommandType
}

type AbstractCommand struct {
	name        CommandName
	typeCommand CommandType
}

func (c *AbstractCommand) GetName() CommandName {
	return c.name
}

func (c *AbstractCommand) GetType() CommandType {
	return c.typeCommand
}
