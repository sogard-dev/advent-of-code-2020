package day17

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `.#.
..#
###`

func TestPart1(t *testing.T) {
	require.Equal(t, 112, part1(testInput))
	require.Equal(t, 313, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 848, part2(testInput))
	require.Equal(t, 2640, part2(utils.GetInput(t, "input.txt")))
}
