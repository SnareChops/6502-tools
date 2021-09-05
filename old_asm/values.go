package old_asm

import (
	"github.com/SnareChops/6502-tools/old_asm/lang"
)

// Parser represents a parser function
type ModeParser = func(string) (string, []byte)

type ValueParser = func(string) []byte

// A parses an absolute memory address
func A(inp string) (string, []byte) {
	if match := Submatch(inp, `^\$(.+)`); match != nil {
		if value := Uint16(match[1]); value != nil {
			return lang.A, value
		}
	}
	return "", nil
}

// I parses an immediate literal value
func I(inp string) (string, []byte) {
	if value := EitherValue(inp, Char, Uint8); value != nil {
		return lang.I, value
	}
	return "", nil
}

// AX parses an absolute with x value
// ex: $1234,x
func AX(inp string) (string, []byte) {
	if match := Submatch(inp, `(?i)^\$(.+),x$`); match != nil {
		if value := Uint16(match[1]); value != nil {
			return lang.AX, value
		}
	}
	return "", nil
}

// AY parses an absolute with y value
// ex: $1234,y
func AY(inp string) (string, []byte) {
	if match := Submatch(inp, `(?i)^\$(.+),y$`); match != nil {
		if value := Uint16(match[1]); value != nil {
			return lang.AY, value
		}
	}
	return "", nil
}

// ZP parses a zero page address
func ZP(inp string) (string, []byte) {
	if match := Submatch(inp, `^\$(.+)$`); match != nil {
		if value := Uint8(match[1]); value != nil {
			return lang.ZP, value
		}
	}
	return "", nil
}

// ZPX parses a zero paged with x address
func ZPX(inp string) (string, []byte) {
	if match := Submatch(inp, `^(.+),x$`); match != nil {
		if _, value := ZP(match[1]); value != nil {
			return lang.ZPX, value
		}
	}
	return "", nil
}

// ZPY parses a zero paged with y address
func ZPY(inp string) (string, []byte) {
	if match := Submatch(inp, `^(.+),y$`); match != nil {
		if _, value := ZP(match[1]); value != nil {
			return lang.ZPY, value
		}
	}
	return "", nil
}

// ZPIX parses a zero page indirect indexed with x address
func ZPIX(inp string) (string, []byte) {
	if match := Submatch(inp, `^\((.+),x\)$`); match != nil {
		if _, value := ZP(match[1]); value != nil {
			return lang.ZPIX, value
		}
	}
	return "", nil
}

// ZPIY parses a zero page indirect indexed with y address
func ZPIY(inp string) (string, []byte) {
	if match := Submatch(inp, `^\((.+)\),y$`); match != nil {
		if _, value := ZP(match[1]); value != nil {
			return lang.ZPIY, value
		}
	}
	return "", nil
}

// R parses a relative value
func R(inp string) (string, []byte) {
	if match := Submatch(inp, `^\$(.+)$`); match != nil {
		println(match)
		if value := Int8(match[1]); value != nil {
			return lang.R, value
		}
	}
	return "", nil
}

// AI parses an Absolute Indirect value
func AI(inp string) (string, []byte) {
	if match := Submatch(inp, `^\((.+)\)$`); match != nil {
		if _, value := A(match[1]); value != nil {
			return lang.AI, value
		}
	}
	return "", nil
}

func LabelAddress(inp string) (string, []byte) {
	if match := Submatch(inp, `^[a-zA-Z][a-zA-Z1-9_]*$`); match != nil {
		return lang.Label, []byte(match[0])
	}
	return "", nil
}
