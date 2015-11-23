package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/yageek/atmel2avr-as/atmel"
	"os"
)

var input = flag.String("input", "", "The file to convert")
var output = flag.String("output", "", "The output file")
var force = flag.Bool("F", false, "Overwrite output")

func main() {

	flag.Parse()

	if *input == "" || *output == "" {
		fmt.Printf("Invalid provided input and output\n")
		os.Exit(-1)
	}

	if fileExistsAtPath(*output) && !(*force) {
		fmt.Printf("File already exists at :%s\nAdd -F to override\n", *output)
		os.Exit(0)
	}

	output, err := os.Create(*output)

	if err != nil {
		panic(err)
	}

	file, err := atmel.NewAtmelFile(*input)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	defer output.Close()

	w := bufio.NewWriter(output)
	file.Convert(w)
}

func fileExistsAtPath(path string) bool {

	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
