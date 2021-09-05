package asm_test

import (
	"strings"
	"testing"

	"github.com/SnareChops/6502-tools/asm"
	"github.com/stretchr/testify/assert"
)

func TestSplitWords(t *testing.T) {
	result := asm.SplitWords("This is a test string")
	expected := []asm.Word{{"This", 0}, {"is", 5}, {"a", 8}, {"test", 10}, {"string", 15}}
	assert.Equal(t, expected, result)
}

func TestLexChar(t *testing.T) {
	result, err := asm.LexChar(asm.Word{"'a'", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected := &asm.Token{asm.CHAR, asm.Location{"test.asm", 2, 4}, "a"}
	assert.Equal(t, expected, result)

	result, err = asm.LexChar(asm.Word{"'a", 3}, "test.asm", 2)
	assert.NotNil(t, err)
	assert.Nil(t, result)

	result, err = asm.LexChar(asm.Word{"'as'", 3}, "test.asm", 2)
	assert.NotNil(t, err)
	assert.Nil(t, result)

	result, err = asm.LexChar(asm.Word{"test", 3}, "test.asm", 2)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestLexInt(t *testing.T) {
	result, err := asm.LexInt(asm.Word{"12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected := &asm.Token{asm.INT, asm.Location{"test.asm", 2, 4}, "12"}
	assert.Equal(t, expected, result)

	result, err = asm.LexInt(asm.Word{"0x12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.INT, asm.Location{"test.asm", 2, 4}, "0x12"}
	assert.Equal(t, expected, result)

	result, err = asm.LexInt(asm.Word{"0b111", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.INT, asm.Location{"test.asm", 2, 4}, "0b111"}
	assert.Equal(t, expected, result)

	result, err = asm.LexInt(asm.Word{"test", 3}, "test.asm", 2)
	assert.Nil(t, err)
	assert.Nil(t, result)
}

func TestLexAddress(t *testing.T) {
	result, err := asm.LexAddress(asm.Word{"$12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected := &asm.Token{asm.ADDRESS, asm.Location{"test.asm", 2, 4}, "12"}
	assert.Equal(t, expected, result)

	result, err = asm.LexAddress(asm.Word{"$0x12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.ADDRESS, asm.Location{"test.asm", 2, 4}, "0x12"}
	assert.Equal(t, expected, result)

	result, err = asm.LexAddress(asm.Word{"$0b11", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.ADDRESS, asm.Location{"test.asm", 2, 4}, "0b11"}
	assert.Equal(t, expected, result)

	result, err = asm.LexAddress(asm.Word{"$abc", 3}, "test.asm", 2)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestLexWord(t *testing.T) {
	result, err := asm.LexWord(asm.Word{"test", 0}, "test.asm", 1)
	assert.Nil(t, err)
	expected := &asm.Token{asm.WORD, asm.Location{"test.asm", 1, 1}, "test"}
	assert.Equal(t, expected, result)

	result, err = asm.LexWord(asm.Word{"'a'", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.CHAR, asm.Location{"test.asm", 2, 4}, "a"}
	assert.Equal(t, expected, result)

	result, err = asm.LexWord(asm.Word{"$12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.ADDRESS, asm.Location{"test.asm", 2, 4}, "$12"}
	assert.Equal(t, expected, result)

	result, err = asm.LexInt(asm.Word{"12", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected = &asm.Token{asm.INT, asm.Location{"test.asm", 2, 4}, "12"}
	assert.Equal(t, expected, result)
}

func TestLexLabel(t *testing.T) {
	result, err := asm.LexLabel(asm.Word{"test:", 3}, "test.asm", 2)
	assert.Nil(t, err)
	expected := &asm.Token{asm.LABEL, asm.Location{"test.asm", 2, 4}, "test"}
	assert.Equal(t, expected, result)
}

func TestLexLine(t *testing.T) {
	result, err := asm.LexLine("LDA 12 'a' $0x12", "test.asm", 1)
	assert.Nil(t, err)
	expected := []*asm.Token{
		{asm.WORD, asm.Location{"test.asm", 1, 1}, "LDA"},
		{asm.INT, asm.Location{"test.asm", 1, 5}, "12"},
		{asm.CHAR, asm.Location{"test.asm", 1, 8}, "a"},
		{asm.ADDRESS, asm.Location{"test.asm", 1, 12}, "0x12"},
	}
	assert.Equal(t, expected, result)

	result, err = asm.LexLine("LDA 'a", "test.asm", 1)
	assert.NotNil(t, err)
	assert.Nil(t, result)
}

func TestLexReader(t *testing.T) {
	result, err := asm.LexReader(strings.NewReader("// test comment\nLDA 12\n0x12\nlabel:"), "t")
	assert.Nil(t, err)
	expected := []*asm.Token{
		{asm.WORD, asm.Location{"t", 2, 1}, "LDA"},
		{asm.INT, asm.Location{"t", 2, 5}, "12"},
		{asm.INT, asm.Location{"t", 3, 1}, "18"},
		{asm.LABEL, asm.Location{"t", 4, 1}, "label"},
	}
	assert.Equal(t, expected, result)
}
