package day23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, "92658374", part1(`389125467`, 10))
	require.Equal(t, "67384529", part1(`389125467`, 100))
	require.Equal(t, "26354798", part1("284573961", 100))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 149245887792, part2(`389125467`))
	require.Equal(t, 166298218695, part2(`284573961`))
}
