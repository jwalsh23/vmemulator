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

func New() *codeWriter {
	return &codeWriter{
		fileBuffer: new(bytes.Buffer),
	}
}

func (c *codeWriter) SetFileName(name string) {
	c.fileName = name
}

func (c *codeWriter) WriteArithmetic(cmd string) {
	c.fileBuffer.WriteString("here is a test\n")
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
