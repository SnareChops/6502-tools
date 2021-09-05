package asm_test

import (
	"testing"

	"github.com/SnareChops/6502-tools/asm"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	match, result := asm.A("$1234")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.A("$65535")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = asm.A("$0xff12")
	require.Equal(t, `A`, match)
	require.Equal(t, []byte{0x12, 0xff}, result)

	match, _ = asm.A("1234")
	require.Empty(t, match)
}

func TestI(t *testing.T) {
	match, result := asm.I("12")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.I("'c'")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0x63}, result)

	match, result = asm.I("0xff")
	require.Equal(t, `I`, match)
	require.Equal(t, []byte{0xff}, result)
}

func TestAX(t *testing.T) {
	match, result := asm.AX("$1234,x")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.AX("$1234,X")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.AX("$0xffd1,x")
	require.Equal(t, `AX`, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, _ = asm.AX("$'c',x")
	require.Empty(t, match)

	match, _ = asm.AX("1234,x")
	require.Empty(t, match)

	match, _ = asm.AX("$1234,y")
	require.Empty(t, match)
}

func TestAY(t *testing.T) {
	match, result := asm.AY("$1234,y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.AY("$1234,Y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.AY("$0xffd1,y")
	require.Equal(t, `AY`, match)
	require.Equal(t, []byte{0xd1, 0xff}, result)

	match, _ = asm.AY("$'c',y")
	require.Empty(t, match)

	match, _ = asm.AY("1234,y")
	require.Empty(t, match)

	match, _ = asm.AY("$1234,x")
	require.Empty(t, match)
}

func TestZP(t *testing.T) {
	match, result := asm.ZP("$12")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.ZP("$0xff")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = asm.ZP("$255")
	require.Equal(t, `ZP`, match)
	require.Equal(t, []byte{0xff}, result)

	match, _ = asm.ZP("$256")
	require.Empty(t, match)

	match, _ = asm.ZP("26")
	require.Empty(t, match)

	match, _ = asm.ZP("$'c'")
	require.Empty(t, match)

	match, _ = asm.ZP("")
	require.Empty(t, match)
}

func TestZPX(t *testing.T) {
	match, result := asm.ZPX("$12,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.ZPX("$0xff,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = asm.ZPX("$255,x")
	require.Equal(t, `ZPX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, _ = asm.ZPX("$256,x")
	require.Empty(t, match)

	match, _ = asm.ZPX("255,x")
	require.Empty(t, match)

	match, _ = asm.ZPX("$'c',x")
	require.Empty(t, match)

	match, _ = asm.ZPX("")
	require.Empty(t, match)
}

func TestZPY(t *testing.T) {
	match, result := asm.ZPY("$12,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.ZPY("$0xff,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = asm.ZPY("$255,y")
	require.Equal(t, `ZPY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, _ = asm.ZPY("$256,y")
	require.Empty(t, match)

	match, _ = asm.ZPY("255,y")
	require.Empty(t, match)

	match, _ = asm.ZPY("$'c',y")
	require.Empty(t, match)

	match, _ = asm.ZPY("")
	require.Empty(t, match)
}

func TestZPIX(t *testing.T) {
	match, result := asm.ZPIX("($12,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.ZPIX("($0xff,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = asm.ZPIX("($255,x)")
	require.Equal(t, `ZPIX`, match)
	require.Equal(t, []byte{0xff}, result)

	match, _ = asm.ZPIX("$255,x")
	require.Empty(t, match)

	match, _ = asm.ZPIX("($256,x)")
	require.Empty(t, match)

	match, _ = asm.ZPIX("(255,x)")
	require.Empty(t, match)
}

func TestZPIY(t *testing.T) {
	match, result := asm.ZPIY("($12),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xc}, result)

	match, result = asm.ZPIY("($0xff),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, result = asm.ZPIY("($255),y")
	require.Equal(t, `ZPIY`, match)
	require.Equal(t, []byte{0xff}, result)

	match, _ = asm.ZPIY("($256),y")
	require.Empty(t, match)

	match, _ = asm.ZPIY("$255,y")
	require.Empty(t, match)

	match, _ = asm.ZPIY("(255),y")
	require.Empty(t, match)
}

func TestR(t *testing.T) {
	match, result := asm.R("$12")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x0c}, result)

	match, result = asm.R("$0x40")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x40}, result)

	match, result = asm.R("$127")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x7f}, result)

	match, result = asm.R("$-128")
	require.Equal(t, `R`, match)
	require.Equal(t, []byte{0x80}, result)

	match, _ = asm.R("$128")
	require.Empty(t, match)

	match, _ = asm.R("26")
	require.Empty(t, match)

	match, _ = asm.R("$'c'")
	require.Empty(t, match)

	match, _ = asm.R("")
	require.Empty(t, match)
}

func TestAI(t *testing.T) {
	match, result := asm.AI("($1234)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0xd2, 0x04}, result)

	match, result = asm.AI("($65535)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0xff, 0xff}, result)

	match, result = asm.AI("($0xff12)")
	require.Equal(t, `AI`, match)
	require.Equal(t, []byte{0x12, 0xff}, result)

	match, _ = asm.AI("(1234)")
	require.Empty(t, match)
}
