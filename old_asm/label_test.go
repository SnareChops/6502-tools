package old_asm_test

import (
	"fmt"
	"testing"

	asm "github.com/SnareChops/6502-tools/old_asm"
	"github.com/stretchr/testify/require"
)

func TestLabel(t *testing.T) {
	match, err := asm.Label("test:", 0)
	require.True(t, match)
	require.Nil(t, err)
	require.Equal(t, "$0", asm.GetLabel("test"))

	match, err = asm.Label("t_16_ahs:", 12)
	require.True(t, match)
	require.Nil(t, err)
	require.Equal(t, "$12", asm.GetLabel("t_16_AHs"))

	match, err = asm.Label("test:", 42)
	require.False(t, match)
	require.NotNil(t, err)
	require.Equal(t, "Duplicate label 'test'", fmt.Sprint(err))

	match, err = asm.Label("test", 64)
	require.False(t, match)
	require.Nil(t, err)

	match, err = asm.Label("222:", 74)
	require.False(t, match)
	require.NotNil(t, err)
	require.Equal(t, "Label must start with a letter '222'", fmt.Sprint(err))

	asm.ClearLabels()
}

func TestReplaceFromLabel(t *testing.T) {
	asm.ClearLabels()

	asm.Label("test:", 255)
	result := asm.ReplaceFromLabel("test")
	require.Equal(t, "$255", result)

}
