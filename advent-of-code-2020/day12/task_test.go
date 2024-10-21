package day12

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `F10
N3
F7
R90
F11`

func TestPart1(t *testing.T) {
	require.Equal(t, 25, part1(testInput))
	require.Equal(t, 420, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 286, part2(testInput))
	require.Equal(t, 42073, part2(utils.GetInput(t, "input.txt")))
}
