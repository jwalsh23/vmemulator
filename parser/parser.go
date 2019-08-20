package parser

import (
	"bufio"
	"fmt"
	"log"
)

// Parser struct
type Parser struct {
	Reader         *bufio.Reader
	currentCommand string
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
	p.currentCommand = string(line)
}
func (p *Parser) PrintCurrent() {
	fmt.Println(p.currentCommand)
}
