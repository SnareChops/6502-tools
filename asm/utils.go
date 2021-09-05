package asm

import "regexp"

// Submatch runs a regex on the string and returns all submatches
func Submatch(inp string, reg string) []string {
	regex := regexp.MustCompile(reg)
	return regex.FindStringSubmatch(inp)
}

// Matcher returns a matcher function
func Matcher(pattern string) func(string) []string {
	regex := regexp.MustCompile(pattern)
	return func(inp string) []string {
		return regex.FindStringSubmatch(inp)
	}
}

// EitherInstruction runs a sequence of Instruction parsers returning the result
// of the first matching parser
func EitherInstruction(inp string, parsers ...InstructionParser) *Instruction {
	for _, parser := range parsers {
		instruction := parser(inp)
		if instruction != nil {
			return instruction
		}
	}
	return nil
}

func EitherMode(inp string, parsers ...ModeParser) (string, []byte) {
	for _, parser := range parsers {
		mode, value := parser(inp)
		if mode != "" {
			return mode, value
		}
	}
	return "", nil
}

func EitherValue(inp string, parsers ...ValueParser) []byte {
	for _, parser := range parsers {
		value := parser(inp)
		if value != nil {
			return value
		}
	}
	return nil
}
