package old_asm

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func AssembleFile(in, out string) error {
	program := &Program{}
	file, err := os.Open(in)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := AssembleReader(file, program)
	if err != nil {
		return err
	}
	err = os.WriteFile(out, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote %d bytes to %s", len(bytes), out)
	return nil
}

// Assemble assembles the provided asm code
// Returns []byte of binary assembled data
func AssembleReader(reader io.Reader, program *Program) ([]byte, error) {
	// Parse ASM
	err := ParseReader(reader, program)
	if err != nil {
		return nil, err
	}
	return program.Compile(), nil
}

// File assembles the provided asm file
// Returns []byte of binary assembled data
func ParseFile(path string, program *Program) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return ParseReader(file, program)
}

// Parse parses a reader
func ParseReader(reader io.Reader, program *Program) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// Trim line
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines and comments
		if len(line) == 0 || (len(line) >= 2 && line[0:2] == "//") {
			continue
		}
		err := ParseLine(line, program)
		if err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// ParseLine parses a line of code
func ParseLine(inp string, program *Program) error {
	// Parse any import statements
	if matched, err := ParseImport(inp, program); matched {
		return err
	}

	// Parse any labels
	if matched, err := ParseLabel(inp, program); matched {
		return err
	}

	// Parse any instructions
	if matched := ParseInstruction(inp, program); matched {
		return nil
	}

	return fmt.Errorf("invalid line '%s'", inp)
}

func ParseImport(inp string, program *Program) (bool, error) {
	if match := Submatch(inp, `(?i)^\import\s+"?([^\n\r"]*)"?$`); match != nil {
		err := ParseFile(match[1], program)
		if err != nil {
			return true, err
		}
		return true, err
	}
	return false, nil
}

func ParseLabel(inp string, program *Program) (bool, error) {
	if match := Submatch(inp, `^([a-zA-Z0-9_]*):(?:\s+//.*)*$`); match != nil {
		name := strings.ToLower(match[1])
		if program.HasLabel(name) {
			return true, fmt.Errorf("duplicate label '%s'", name)
		}
		program.Labels[name] = program.Size()
		return true, nil
	}
	return false, nil
}

func ParseInstruction(inp string, program *Program) bool {
	if instruction := EitherInstruction(inp, Instructions...); instruction != nil {
		program.Instructions = append(program.Instructions, *instruction)
		return true
	}
	return false
}
