package day11

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput string = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestPart1(t *testing.T) {
	require.Equal(t, 37, part1(testInput))
	require.Equal(t, 2406, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 26, part2(testInput))
	require.Equal(t, 2149, part2(utils.GetInput(t, "input.txt")))
}
