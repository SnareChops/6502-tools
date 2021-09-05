package asm

// type Parser = func(inp string) ([]rune, string)

// func Any() Parser {
// 	return func(inp string) ([]rune, string) {
// 		return []rune{rune(inp[0])}, inp[1:]
// 	}
// }

// func Where(f func(rune) bool) Parser {
// 	return func(inp string) ([]rune, string) {
// 		if f(rune(inp[0])) {
// 			return []rune{rune(inp[0])}, inp[1:]
// 		}
// 		return []rune{}, inp
// 	}
// }

// func Rune(char rune) Parser {
// 	return func(inp string) ([]rune, string) {
// 		if rune(inp[0]) == char {
// 			return []rune{char}, inp[1:]
// 		}
// 		return []rune{}, inp
// 	}
// }

// func Digit() Parser {
// 	return Where(unicode.IsDigit)
// }

// func Letter() Parser {
// 	return Where(unicode.IsLetter)
// }

// func Space() Parser {
// 	return Where(unicode.IsSpace)
// }

// func Alphanumeric() Parser {
// 	return Where(func(r rune) bool {
// 		return unicode.IsDigit(r) || unicode.IsLetter(r)
// 	})
// }

// func Digits() Parser {
// 	return Collect(unicode.IsDigit)
// }

// func Letters() Parser {
// 	return Collect(unicode.IsLetter)
// }

// func Whitespace() Parser {
// 	return Collect(unicode.IsSpace)
// }

// func Alphanumerics() Parser {
// 	return Collect(func(r rune) bool {
// 		return unicode.IsDigit(r) || unicode.IsLetter(r)
// 	})
// }

// func String(chars string) Parser {
// 	parser := Rune(rune(chars[0]))
// 	for _, w := range chars[1:] {
// 		parser = Then(parser, Rune(w))
// 	}
// 	return parser
// }

// func Then(p1, p2 Parser) Parser {
// 	return func(inp string) ([]rune, string) {
// 		r1, rest1 := p1(inp)
// 		if len(r1) == 0 {
// 			return []rune{}, inp
// 		}
// 		r2, rest := p2(rest1)
// 		if len(r2) == 0 {
// 			return []rune{}, inp
// 		}
// 		result := []rune{}
// 		for _, r := range r1 {
// 			result = append(result, r)
// 		}
// 		for _, r := range r2 {
// 			result = append(result, r)
// 		}
// 		return result, rest
// 	}
// }

// func Choice(parsers ...Parser) Parser {
// 	return func(inp string) ([]rune, string) {
// 		for _, p := range parsers {
// 			result, rest := p(inp)
// 			if len(result) > 0 {
// 				return result, rest
// 			}
// 		}
// 		return []rune{}, inp
// 	}
// }

// func Collect(f func(rune) bool) Parser {
// 	return func(inp string) ([]rune, string) {
// 		runes := []rune{}
// 		rest := inp
// 		for {
// 			char := rune(inp[len(runes)])
// 			if !f(char) {
// 				return runes, rest
// 			}
// 			runes = append(runes, char)
// 			rest = inp[1:]
// 		}
// 	}
// }

// type TokenParser = func(inp string) (*Token, string)

// func Instruction(str string) TokenParser {
// 	return func(inp string) (*Token, string) {
// 		Then(String(str)
// 	}
// }
