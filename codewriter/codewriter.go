package codewriter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

type codeWriter struct {
	fileName   string
	fileBuffer *bytes.Buffer
}

var arithmeticCommandMap = map[string]string{
	"add": addCommand,
	"sub": subCommand,
	"neg": negCommand,
	"eq":  eqCommand,
	"lt":  ltCommand,
}

func New() *codeWriter {
	return &codeWriter{
		fileBuffer: new(bytes.Buffer),
	}
}

func (c *codeWriter) SetFileName(name string) {
	c.fileName = name
}

func (c *codeWriter) WriteArithmetic(cmd string) {
	c.fileBuffer.WriteString(arithmeticCommandMap[cmd])
	c.fileBuffer.WriteString("\n")
}
func (c *codeWriter) WritePushPop(cmd, segment string, index int) {

}
func (c *codeWriter) Close() {
	if c.fileName == "" {
		fmt.Println(string(c.fileBuffer.Bytes()))
		return
	}
	err := ioutil.WriteFile(c.fileName, c.fileBuffer.Bytes(), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	addCommand = `@SP
AM=M-1
D=M
A=A-1
M=M+D`
	subCommand = `@SP
AM=M-1
D=M
A=A-1
M=M-D`
	negCommand = `@SP
A=A-1
M=-M`
	eqCommand = `@SP
AM=M-1
D=-M
A=A-1
M=M+D
@TRUE
M;JE
@FALSE
0;JMP`
	ltCommand = `@SP
AM=M-1
D=-M
A=A-1
M=M+D
@FALSE
M;JGE
@TRUE
0;JMP`
	gtCommand = `@SP
AM=M-1
D=-M
A=A-1
M=M+D
@FALSE
M;LTE
@TRUE
0;JMP`
	andCommand = `AM=M-1
D=M 
A=A-1
M=M&D`
	orCommand = `AM=M-1
D=M 
A=A-1
M=M&D`
)
