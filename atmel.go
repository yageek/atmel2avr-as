package main

import (
	"bufio"
	"io"
	"os"
)

type AtmelFile struct {
	file         *os.File
	bufio        *bufio.Reader
	convertError error
}

func NewAtmelFile(filename string) (*AtmelFile, error) {

	file, err := os.Open(filename)

	if err != nil {
		return err
	}
}

func (a *AtmelFile) Convert(w io.Writer) bool {
	a.bufio = bufio.NewReader(a.file)

	line, _, err := a.bufio.ReadLine()
}

func (a *AtmelFile) ConvertError() error {
	return a.convertError
}

func (a *AtmelFile) Close() {
	a.file.Close()
}
