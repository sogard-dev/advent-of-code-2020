package day14

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 165, int(part1(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)))
	require.Equal(t, 15172047086292, int(part1(utils.GetInput(t, "input.txt"))))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 208, int(part2(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)))
	require.Equal(t, 4197941339968, int(part2(utils.GetInput(t, "input.txt"))))
}
