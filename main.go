package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/SnareChops/6502-tools/asm"
	"github.com/hajimehoshi/ebiten"
)

const (
	SCREEN_HEIGHT = 1280
	SCREEN_WIDTH  = 720
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
		command = simulate
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

func simulate([]string) error {
	game := &Game{}
	ebiten.SetWindowSize(SCREEN_HEIGHT, SCREEN_WIDTH)
	ebiten.SetWindowTitle("6502 Simulator")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	return nil
}

func invalid([]string) error {
	return errors.New("invalid subcommand")
}
