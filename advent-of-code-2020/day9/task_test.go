package day9

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput string = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestPart1(t *testing.T) {
	require.Equal(t, 127, part1(testInput, 5))
	require.Equal(t, 57195069, part1(utils.GetInput(t, "input.txt"), 25))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 62, part2(testInput, 5))
	require.Equal(t, 7409241, part2(utils.GetInput(t, "input.txt"), 25))
}
