package main

import (
	"errors"
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
		command = assemble
		break
	case "sim":
		command = sim
		break
	default:
		command = invalid
	}
}

func main() {
	err := command(os.Args[2:])
	if err != nil {
		panic(err)
	}
}
func assemble(args []string) error {
	if len(args) < 1 {
		return errors.New("missing args")
	}

	filepath := path.Clean(args[0])
	filename := path.Base(filepath)
	basename := strings.Replace(filename, path.Ext(filepath), "", -1)
	outputname := basename + ".bin"

	// Assemble file
	data, err := asm.File(filepath)
	if err != nil {
		return err
	}

	// Create binary file
	output, err := os.Create(outputname)
	if err != nil {
		return err
	}
	defer output.Close()

	// Output to file
	size, err := output.Write(data)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes to %s", size, outputname)
	return nil
}
