package main

import (
	"bufio"
	"log"
	"os"

	"github.com/vmemulator/codewriter"
	"github.com/vmemulator/parser"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Must suppply input file and output file as parameters")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file with err %s", err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	p := parser.Parser{
		Reader: reader,
	}
	codeWriter := codewriter.New()
	for p.HasMoreCommands() {
		p.Advance()
		if p.CommandType() == parser.C_ARITHMETIC {
			codeWriter.WriteArithmetic(p.Arg1())
			continue
		}
		codeWriter.WritePushPop(p.CommandType(), p.Arg1(), p.Arg2())
	}
	// codeWriter.SetFileName(os.Args[2])
	codeWriter.Close()
}
