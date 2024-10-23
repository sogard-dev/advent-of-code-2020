package day16

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 71, part1(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`))
	require.Equal(t, 23036, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 12, part2(`departure class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`))
	require.Equal(t, 1909224687553, part2(utils.GetInput(t, "input.txt")))
}
