package day22

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestPart1(t *testing.T) {
	require.Equal(t, 306, part1(testInput))
	require.Equal(t, 35202, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 291, part2(testInput))
	require.Equal(t, 32317, part2(utils.GetInput(t, "input.txt")))
}
