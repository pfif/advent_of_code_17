package command

import (
	"fmt"

	"florent/adventofcode/2022/day7/utils"
)

func Parse(input []byte) ([]Command, error) {
	state := NewParserState()
	err := utils.Reduce(input, parseReducer, state)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing input: %w", err)
	}
	state.finalizeCurrentCommand()

	return state.completedCommands, nil
}

func parseReducer(acc *parserState, char byte) error {
	switch char {
	case '$':
		switch acc.status {
		case notParsing, parsingStdout:
			acc.finalizeCurrentCommand()
			err := acc.switchStatus(parsingProgram)
			if err != nil {
				return err
			}
		default:
			return parseError(char, acc.status)
		}
	case ' ':
		switch acc.status {
		case notParsing, parsingArgument:
		case parsingProgram:
			if len(acc.currentText()) > 0 {
				err := acc.switchStatus(parsingArgument)
				if err != nil {
					return err
				}
			}
			return nil
		case parsingStdout:
			acc.addChar(char)
		}
	case '\n':
		switch acc.status {
		case notParsing:
		case parsingArgument, parsingProgram:
			err := acc.switchStatus(parsingStdout)
			if err != nil {
				return err
			}
		case parsingStdout:
			acc.addChar(char)
		}
	default:
		switch acc.status {
		case notParsing:
		case parsingArgument, parsingStdout, parsingProgram:
			acc.addChar(char)
		}
	}

	return nil
}

func parseError(char byte, status parserStatus) error {
	return fmt.Errorf("Encountered %v in status %s", char, status)
}

type parserStatus string

const (
	notParsing      = "NOT_PARSING"
	parsingProgram  = "PARSING_PROGRAM"
	parsingArgument = "PARSING_ARGUMENT"
	parsingStdout   = "PARSING_STDOUT"
)

type parserState struct {
	status            parserStatus
	program           []byte
	argument          []byte
	stdout            []byte
	current           *[]byte
	completedCommands []Command
}

func (p *parserState) switchStatus(status parserStatus) error {
	set := func(currentText *[]byte) {
		p.status = status
		p.current = currentText
	}

	switch status {
	case notParsing:
		set(nil)
	case parsingProgram:
		set(&p.program)
	case parsingArgument:
		set(&p.argument)
	case parsingStdout:
		set(&p.stdout)
	default:
		return fmt.Errorf("Attempting to switch to unknown status %s", status)
	}

	return nil
}

func (p *parserState) currentText() []byte {
	return *p.current
}

func (p *parserState) addChar(char byte) {
	*(p.current) = append(*p.current, char)
}

func (p *parserState) finalizeCurrentCommand() {
	if p.status != notParsing {
		p.completedCommands = append(
			p.completedCommands,
			NewCommand(string(p.program), string(p.argument), string(p.stdout)),
		)
	}

	p.resetParsing()
}

func (p *parserState) resetParsing() {
	p.status = notParsing
	p.program = []byte{}
	p.stdout = []byte{}
	p.argument = []byte{}
	p.current = nil
}

func NewParserState() *parserState {
	result := parserState{
		completedCommands: []Command{},
	}
	result.resetParsing()
	return &result
}
