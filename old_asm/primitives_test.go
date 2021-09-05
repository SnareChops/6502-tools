package old_asm_test

import (
	"testing"

	asm "github.com/SnareChops/6502-tools/old_asm"
	"github.com/stretchr/testify/require"
)

func TestUint8(t *testing.T) {
	result := asm.Uint8("6")
	require.Equal(t, []byte{0x06}, result)

	result = asm.Uint8("255")
	require.Equal(t, []byte{0xff}, result)

	result = asm.Uint8("256")
	require.Nil(t, result)

	result = asm.Uint8("65535")
	require.Nil(t, result)

	result = asm.Uint8("x")
	require.Nil(t, result)

}

func TestUint16(t *testing.T) {
	result := asm.Uint16("6")
	require.Equal(t, []byte{0x06, 0x00}, result)

	result = asm.Uint16("255")
	require.Equal(t, []byte{0xff, 0x00}, result)

	result = asm.Uint16("65535")
	require.Equal(t, []byte{0xff, 0xff}, result)

	result = asm.Uint16("127236748")
	require.Nil(t, result)

	result = asm.Uint16("h")
	require.Nil(t, result)
}

func TestChar(t *testing.T) {
	result := asm.Char("'c'")
	require.Equal(t, []byte{0x63}, result)

	result = asm.Char("'C'")
	require.Equal(t, []byte{0x43}, result)

	result = asm.Char("' '")
	require.Equal(t, []byte{0x20}, result)

	result = asm.Char("'''")
	require.Equal(t, []byte{0x27}, result)

	result = asm.Char("'9'")
	require.Equal(t, []byte{0x39}, result)

	result = asm.Char("'")
	require.Nil(t, result)

	result = asm.Char("")
	require.Nil(t, result)

	result = asm.Char("'😀'")
	require.Nil(t, result)
}
