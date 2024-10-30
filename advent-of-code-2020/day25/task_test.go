package day25

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 14897079, part1(`5764801
17807724`))
	require.Equal(t, 0, part1(utils.GetInput(t, "input.txt")))
}
