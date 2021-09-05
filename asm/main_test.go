package asm_test

import (
	"testing"

	"github.com/SnareChops/6502-tools/asm"
	"github.com/SnareChops/6502-tools/asm/lang"
	"github.com/stretchr/testify/require"
)

func TestParseLine(t *testing.T) {
	program := &asm.Program{}
	result := asm.ParseLine("label:", program)
	require.Nil(t, result)
	require.Equal(t, program.Labels["label"], 0)

	program = &asm.Program{}
	result = asm.ParseLine("LDA $0xff12", program)
	require.Nil(t, result)
	require.Equal(t, program.Instructions[0].Name, lang.LDA)
	require.Equal(t, program.Instructions[0].Value, []byte{0x12, 0xff})

	program = &asm.Program{}
	result = asm.ParseLine("sfoksdf wefoh sdf", program)
	require.NotNil(t, result)
}
