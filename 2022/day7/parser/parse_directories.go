package parser

import (
	"florent/adventofcode/2022/day7/parser/command"
	"florent/adventofcode/2022/day7/parser/directory"
	"florent/adventofcode/2022/day7/parser/input"
	"fmt"
)

// ParseDirectories parse sequence of commands from the raw input, and convert it to a directory tree. Return the root Directory
func ParseDirectories() (directory.Directory, error) {
	raw, errGet := input.GetInput()
	if errGet != nil {
		return nil, fmt.Errorf("Could not read input: %w", errGet)
	}
	commands, errParse := command.Parse(raw)
	if errParse != nil {
		return nil, errParse
	}

	adaptedCommands := []directory.Command{}
	for _, command := range commands {
		adaptedCommands = append(adaptedCommands, &commandAdapter{command: command})
	}

	return directory.ConvertCommands(adaptedCommands)
}

type Directory directory.Directory

type commandAdapter struct {
	command command.Command
}

func (c *commandAdapter) Name() string {
	return c.command.Program()
}

func (c *commandAdapter) Rawlisting() string {
	return c.command.Stdout()
}

func (c *commandAdapter) TargetDirectory() string {
	return c.command.Argument()
}
