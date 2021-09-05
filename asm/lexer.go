package asm

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

const (
	WORD = iota
	CHAR
	INT
	ADDRESS
	LABEL
)

type Location struct {
	File   string
	Line   int
	Column int
}

type Token struct {
	Typ int
	Loc Location
	Val string
}
type Word struct {
	Val string
	Loc int
}

func LexError(file string, line, col int, message string) error {
	return fmt.Errorf("%s:%d:%d ERROR: %s", file, line, col, message)
}

func SplitWords(str string) []Word {
	result := []Word{}
	newWord := true
	for index, char := range str {
		if unicode.IsSpace(char) {
			newWord = true
			continue
		}
		if newWord {
			result = append(result, Word{string(char), index})
			newWord = false
			continue
		}
		result[len(result)-1].Val += string(char)
	}
	return result
}

func LexChar(word Word, file string, line int) (*Token, error) {
	if word.Val[0] != '\'' {
		return nil, nil
	}
	if word.Val[len(word.Val)-1] != '\'' {
		return nil, LexError(file, line, word.Loc+1, "Invalid char definition found. Missing closing quote")
	}
	if len(word.Val) != 3 {
		return nil, LexError(file, line, word.Loc+1, "Invalid char definition found. Char must only contain one character.")
	}
	return &Token{CHAR, Location{file, line, word.Loc + 1}, strings.Split(word.Val, "")[1]}, nil
}

func LexInt(word Word, file string, line int) (*Token, error) {
	if val, ok := strconv.ParseInt(word.Val, 0, 16); ok == nil {
		return &Token{INT, Location{file, line, word.Loc + 1}, fmt.Sprint(val)}, nil
	}
	return nil, nil
}

func LexAddress(word Word, file string, line int) (*Token, error) {
	if word.Val[0] != '$' {
		return nil, nil
	}
	if _, ok := strconv.ParseInt(word.Val[1:len(word.Val)-1], 0, 16); ok == nil {
		return &Token{ADDRESS, Location{file, line, word.Loc + 1}, word.Val[1:len(word.Val)]}, nil
	}
	return nil, LexError(file, line, word.Loc+1, "Invalid address definition found. Address must be a number")
}

func LexLabel(word Word, file string, line int) (*Token, error) {
	if word.Val[len(word.Val)-1] == ':' {
		if !unicode.IsLetter(rune(word.Val[0])) {
			return nil, LexError(file, line, word.Loc, "Invalid label definition found. Label must start with a letter")
		}
		return &Token{LABEL, Location{file, line, word.Loc + 1}, word.Val[0 : len(word.Val)-1]}, nil
	}
	return nil, nil
}

func LexWord(word Word, file string, line int) (*Token, error) {
	if result, err := LexChar(word, file, line); result != nil || err != nil {
		return result, err
	}
	if result, err := LexAddress(word, file, line); result != nil || err != nil {
		return result, err
	}
	if result, err := LexInt(word, file, line); result != nil || err != nil {
		return result, err
	}
	if result, err := LexLabel(word, file, line); result != nil || err != nil {
		return result, err
	}
	return &Token{WORD, Location{file, line, word.Loc + 1}, word.Val}, nil
}

func LexLine(contents, file string, line int) ([]*Token, error) {
	tokens := []*Token{}
	for _, word := range SplitWords(contents) {
		token, err := LexWord(word, file, line)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func LexReader(reader io.Reader, file string) ([]*Token, error) {
	scanner := bufio.NewScanner(reader)
	line := 0
	result := []*Token{}
	for scanner.Scan() {
		line++
		contents := scanner.Text()
		if len(contents) == 0 || (len(contents) >= 2 && contents[0:2] == "//") {
			continue
		}
		tokens, err := LexLine(contents, file, line)
		if err != nil {
			return nil, err
		}
		result = append(result, tokens...)
	}
	return result, nil
}
