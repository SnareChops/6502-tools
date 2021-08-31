package parse_test

import (
	"testing"

	"github.com/SnareChops/6502-tools/parse"
	"github.com/stretchr/testify/require"
)

func TestInclude(t *testing.T) {
	file := parse.Include(".include something.asm")
	require.Equal(t, "something.asm", file)

	file = parse.Include(`.include "another.asm"`)
	require.Equal(t, "another.asm", file)

	file = parse.Include(".InCLudE insensitive.asm")
	require.Equal(t, "insensitive.asm", file)

	file = parse.Include(`.include "file with spaces.asm"`)
	require.Equal(t, "file with spaces.asm", file)

	file = parse.Include(`.include /absolute/path/to/file.asm`)
	require.Equal(t, "/absolute/path/to/file.asm", file)

	file = parse.Include(".include")
	require.Empty(t, file)

	file = parse.Include("include nodot.asm")
	require.Empty(t, file)
}
