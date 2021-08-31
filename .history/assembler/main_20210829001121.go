package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/snarechops/6502-assembler/asm"
)

var (
	command func([]string) error
	input   string
)

func init() {
	switch os.Args[1] {
	case "asm":
		command = asm
		break
	case "sim":
		command = sim
		break
	default:
		command = invalid
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := command(os.Args[2:])
	// Assemble file
	data, err := asm.File(args.Input)
	check(err)

	// Create output file name
	filename := path.Base(args.Input)
	basename := strings.Replace(filename, path.Ext(args.Input), "", -1)
	outputname := basename + ".bin"

	// Create binary file
	output, err := os.Create(outputname)
	check(err)
	defer output.Close()

	// Output to file
	size, err := output.Write(data)
	check(err)

	fmt.Printf("Wrote %d bytes to %s", size, outputname)
}

type options struct {
	Input string
}

func processArgs(args []string) *options {

	return &options{
		Input: path.Clean(args[1]),
	}
}
