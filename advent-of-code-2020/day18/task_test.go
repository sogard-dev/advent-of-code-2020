package day18

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 26, part1(`2 * 3 + (4 * 5)`))
	require.Equal(t, 437, part1(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	require.Equal(t, 12240, part1(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	require.Equal(t, 13632, part1(`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`))
	require.Equal(t, 21993583522852, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1445, part2(`5 + (8 * 3 + 9 + 3 * 4 * 3)`))
	require.Equal(t, 51, part2(`1 + (2 * 3) + (4 * (5 + 6))`))
	require.Equal(t, 669060, part2(`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`))
	require.Equal(t, 122438593522757, part2(utils.GetInput(t, "input.txt")))
}
