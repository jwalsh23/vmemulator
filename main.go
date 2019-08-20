package main

import (
	"bufio"
	"log"
	"os"

	"github.com/vmemulator/parser"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file with err %s", err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	p := parser.Parser{
		Reader: reader,
	}
	for p.HasMoreCommands() {
		p.Advance()
		p.PrintCurrent()
	}
}
