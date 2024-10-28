package day21

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2020/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestPart1(t *testing.T) {
	require.Equal(t, 5, part1(testInput))
	require.Equal(t, 2061, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, "mxmxvkd,sqjhc,fvjkl", part2(testInput))
	require.Equal(t, "cdqvp,dglm,zhqjs,rbpg,xvtrfz,tgmzqjz,mfqgx,rffqhl", part2(utils.GetInput(t, "input.txt")))
}
