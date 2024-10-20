package day10

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput string = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestPart1(t *testing.T) {
	require.Equal(t, 22*10, part1(testInput))
	require.Equal(t, 2400, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 19208, part2(testInput))
	require.Equal(t, 338510590509056, part2(utils.GetInput(t, "input.txt")))
}
