package day24

import (
	"strings"
)

type direction struct {
	str    string
	dx, dy int
}

func getDirections() []direction {
	return []direction{
		{str: "e", dx: 2},
		{str: "w", dx: -2},
		{str: "nw", dx: -1, dy: -1},
		{str: "ne", dx: 1, dy: -1},
		{str: "sw", dx: -1, dy: 1},
		{str: "se", dx: 1, dy: 1},
	}
}

func parse(input string) []direction {
	ret := []direction{}
	directions := getDirections()
	for i := 0; i < len(input); {
		for _, d := range directions {
			dir := d.str
			if i+len(dir) <= len(input) && input[i:i+len(dir)] == dir[:] {
				ret = append(ret, d)
				i += len(dir)
			}
		}
	}

	return ret
}

type pos struct {
	x, y int
}

func part1(input string) int {
	return len(solve(input))
}

func part2(input string) int {
	blacks := solve(input)
	days := 100
	directionDeltas := getDirections()

	for range days {
		adjacents := map[pos]int{}
		for p := range blacks {
			for _, d := range directionDeltas {
				adjacents[pos{
					x: p.x + d.dx,
					y: p.y + d.dy,
				}]++
			}
		}

		newBlacks := map[pos]bool{}
		for k, v := range adjacents {
			if _, exists := blacks[k]; exists {
				if v == 1 || v == 2 {
					newBlacks[k] = true
				}
			} else {
				if v == 2 {
					newBlacks[k] = true
				}
			}
		}

		blacks = newBlacks
	}

	return len(blacks)
}

func solve(input string) map[pos]bool {
	tiles := map[pos]bool{}

	for _, line := range strings.Split(input, "\n") {
		directions := parse(line)
		x, y := 0, 0
		for _, dir := range directions {
			x += dir.dx
			y += dir.dy
		}
		p := pos{x: x, y: y}
		tiles[p] = !tiles[p]
	}

	for k, v := range tiles {
		if !v {
			delete(tiles, k)
		}
	}

	return tiles
}
