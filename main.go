package main

import (
	"bufio"
	"github.com/yageek/atmel2avr-as/atmel"
	"os"
)

func main() {

	file, err := atmel.NewAtmelFile("tn2313.h")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	output, err := os.Create("output")
	if err != nil {
		panic(err)
	}

	defer output.Close()

	w := bufio.NewWriter(output)
	file.Convert(w)
}
