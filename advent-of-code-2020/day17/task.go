package day17

import (
	"strings"
)

func part1(input string) int {
	return solve(input, 0)
}

func part2(input string) int {
	return solve(input, 1)
}

type hypercube struct {
	x, y, z, w int
}

func solve(input string, wdiff int) int {
	bootCycles := 6

	activeCubes := map[hypercube]bool{}

	for c, line := range strings.Split(input, "\n") {
		for r, l := range line {
			if l == '#' {
				activeCubes[hypercube{
					x: c,
					y: r,
				}] = true
			}
		}
	}

	for cycle := 1; cycle <= bootCycles; cycle++ {
		activeCubes = runCycle(activeCubes, wdiff)
	}

	return len(activeCubes)
}
func runCycle(activeCubes map[hypercube]bool, wdiff int) map[hypercube]bool {
	neighbours := map[hypercube]int{}
	for k, _ := range activeCubes {
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				for dz := -1; dz < 2; dz++ {
					for dw := -wdiff; dw <= wdiff; dw++ {
						if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
							neighbours[hypercube{
								x: k.x + dx,
								y: k.y + dy,
								z: k.z + dz,
								w: k.w + dw,
							}]++
						}
					}

				}
			}
		}
	}

	nextActiveCubes := map[hypercube]bool{}
	for k, v := range neighbours {
		if _, exist := activeCubes[k]; exist {
			if v == 2 || v == 3 {
				nextActiveCubes[k] = true
			}
		} else {
			if v == 3 {
				nextActiveCubes[k] = true
			}
		}
	}

	return nextActiveCubes
}
