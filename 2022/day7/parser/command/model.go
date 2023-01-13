package command

type Command interface {
	Program() string
	Argument() string
	Stdout() string
}

type commandImpl struct {
	program  string
	argument string
	stdout   string
}

func (c *commandImpl) Program() string {
	return c.program
}

func (c *commandImpl) Argument() string {
	return c.argument
}

func (c *commandImpl) Stdout() string {
	return c.stdout
}

func NewCommand(program, argument, stdout string) Command {
	return &commandImpl{
		program:  program,
		argument: argument,
		stdout:   stdout,
	}
}
