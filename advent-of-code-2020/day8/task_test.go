package day8

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 5, part1(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	require.Equal(t, 1816, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 8, part2(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`))
	require.Equal(t, 1149, part2(utils.GetInput(t, "input.txt")))
}
