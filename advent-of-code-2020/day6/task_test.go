package day6

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 11, part1(`abc

a
b
c

ab
ac

a
a
a
a

b`))
	require.Equal(t, 6549, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 6, part2(`abc

a
b
c

ab
ac

a
a
a
a

b`))
	require.Equal(t, 3466, part2(utils.GetInput(t, "input.txt")))
}
