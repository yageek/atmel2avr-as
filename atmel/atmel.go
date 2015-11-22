package atmel

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const (
	DEVICE_STATEMENT = ".device"
	ARCH_STATEMENT   = ".arch"
	EQU_STATEMENT    = ".equ"
	DEF_STATEMENT    = ".def"
)

type AtmelFile struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewAtmelFile(filename string) (*AtmelFile, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	return &AtmelFile{file: file, scanner: nil}, nil
}

func (a *AtmelFile) Convert(w io.Writer) {
	a.scanner = bufio.NewScanner(a.file)
	a.scanner.Split(bufio.ScanLines)

	for a.scanner.Scan() {
		convert := ConvertLine(a.scanner.Text())
		w.Write([]byte(convert))
	}

	w.Write([]byte("\n"))
}

func (a *AtmelFile) ConvertError() error {
	return a.scanner.Err()
}

func (a *AtmelFile) Close() {
	a.file.Close()
}

func ConvertLine(line string) string {

	returnLine := line
	if strings.Contains(returnLine, DEVICE_STATEMENT) {
		returnLine = strings.Replace(returnLine, DEVICE_STATEMENT, ARCH_STATEMENT, 1)
	} else if strings.Contains(returnLine, EQU_STATEMENT) {
		returnLine = strings.Replace(returnLine, "=", ",", 1)
	} else if strings.Contains(returnLine, DEF_STATEMENT) {
		returnLine = strings.Replace(returnLine, "=", ",", 1)
		returnLine = strings.Replace(returnLine, DEF_STATEMENT, EQU_STATEMENT, 1)
	}

	return returnLine + "\n"
}
