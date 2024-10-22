package day15

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 436, part1(`0,3,6`))
	require.Equal(t, 517, part1(`5,2,8,16,18,0,1`))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 175594, part2(`0,3,6`))
	require.Equal(t, 1047739, part2(`5,2,8,16,18,0,1`))
}
