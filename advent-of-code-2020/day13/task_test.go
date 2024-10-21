package day13

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `939
7,13,x,x,59,x,31,19`

func TestPart1(t *testing.T) {
	require.Equal(t, 295, part1(testInput))
	require.Equal(t, 2406, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 3417, part2(`q
	17,x,13,19`))
	require.Equal(t, 1068781, part2(testInput))
	require.Equal(t, 225850756401039, part2(utils.GetInput(t, "input.txt")))
}
