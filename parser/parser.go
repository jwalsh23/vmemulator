package parser

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Parser struct
type Parser struct {
	Reader         *bufio.Reader
	currentCommand command
}
type command struct {
	commandType string
	arg1        string
	arg2        int
}

const (
	C_ARITHMETIC = "C_ARITHMETIC"
	C_PUSH       = "C_PUSH"
	C_POP        = "C_POP"
	C_LABEL      = "C_LABEL"
	C_GOTO       = "C_GOTO"
	C_IF         = "C_IF"
	C_FUNCTION   = "C_FUNCTION"
	C_RETURN     = "C_RETURN"
	C_CALL       = "C_CALL"
)
const (
	Push = "push"
	Pop  = "pop"
	Add  = "add"
	Sub  = "sub"
	Neg  = "neg"
	Eq   = "eq"
	Gt   = "gt"
	Lt   = "lt"
	And  = "and"
	Or   = "or"
	Not  = "not"
)

var commandMap = map[string]string{
	Push: C_PUSH,
	Pop:  C_POP,
	Add:  C_ARITHMETIC,
	Sub:  C_ARITHMETIC,
	Neg:  C_ARITHMETIC,
	Eq:   C_ARITHMETIC,
	Gt:   C_ARITHMETIC,
	Lt:   C_ARITHMETIC,
	And:  C_ARITHMETIC,
	Or:   C_ARITHMETIC,
	Not:  C_ARITHMETIC,
}

func (p *Parser) HasMoreCommands() bool {
	_, err := p.Reader.Peek(1)
	if err != nil {
		return false
	}
	return true
}
func (p *Parser) Advance() {
	line, _, err := p.Reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	commandString := removeComment(string(line))
	if commandString == "" {
		if p.HasMoreCommands() {
			p.Advance()
		}
	}
	p.currentCommand = parseCommand(commandString)
}
func (p *Parser) PrintCurrent() {
	fmt.Printf("%+v\n", p.currentCommand)
}
func (p *Parser) Arg1() string {
	return p.currentCommand.arg1
}
func (p *Parser) Arg2() int {
	return p.currentCommand.arg2
}

func removeComment(line string) string {
	lineSlc := strings.Split(line, "//")
	return strings.TrimRight(lineSlc[0], " ")
}

func parseCommand(commandString string) command {
	commandSlc := strings.Split(commandString, " ")
	if len(commandSlc) == 1 {
		return command{
			commandType: commandMap[commandSlc[0]],
			arg1:        commandSlc[0],
		}
	}
	if len(commandSlc) == 3 {
		arg2, err := strconv.Atoi(commandSlc[2])
		if err != nil {
			log.Fatal(err)
		}
		return command{
			commandType: commandMap[commandSlc[0]],
			arg1:        commandSlc[1],
			arg2:        arg2,
		}
	}
	return command{}
}
