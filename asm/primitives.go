package asm

import (
	"strconv"

	"github.com/SnareChops/6502-tools/conv"
)

// Uint16 parses a uint16 literal
func Uint16(inp string) []byte {
	i, err := strconv.ParseUint(inp, 0, 16)
	if err == nil {
		return conv.Uint16(uint16(i))
	}
	return nil
}

// Int8 parses an int8 literal
func Int8(inp string) []byte {
	i, err := strconv.ParseInt(inp, 0, 8)
	if err == nil {
		return []byte{byte(int8(i))}
	}
	return nil
}

// Uint8 parses a uint8 literal
func Uint8(inp string) []byte {
	i, err := strconv.ParseUint(inp, 0, 8)
	if err == nil {
		return []byte{uint8(i)}
	}
	return nil
}

// Char parses a char literal
func Char(inp string) []byte {
	if match := Submatch(inp, `^'(.)'$`); match != nil {
		ascii := []rune(match[1])
		if ascii[0] < 256 {
			return []byte{uint8(ascii[0])}
		}
	}
	return nil
}
