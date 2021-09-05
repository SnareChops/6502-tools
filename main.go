package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/SnareChops/6502-tools/asm"
	"github.com/SnareChops/6502-tools/gui"
)

const (
	SCREEN_HEIGHT = 1280
	SCREEN_WIDTH  = 720
)

var (
	command func([]string) error
)

func init() {
	if len(os.Args) < 2 {
		command = simulate
	} else {
		switch os.Args[1] {
		case "asm":
			command = assemble
		case "sim":
			command = simulate
		default:
			command = simulate
		}
	}
}

func main() {
	args := []string{}
	if len(os.Args) >= 2 {
		args = os.Args[2:]
	}
	err := command(args)
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
	err := asm.AssembleFile(filepath, outputname)
	if err != nil {
		return err
	}

	return nil
}

func simulate([]string) error {
	// game := &Game{}
	// ebiten.SetWindowSize(SCREEN_HEIGHT, SCREEN_WIDTH)
	// ebiten.SetWindowTitle("6502 Simulator")

	// if err := ebiten.RunGame(game); err != nil {
	// 	log.Fatal(err)
	// }
	// return nil
	fmt.Println("Starting GUI")
	gui.Start()
	return nil
}
