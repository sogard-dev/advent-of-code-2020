package day12

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type pos struct {
	N int
	E int
}

func (s pos) manhatten() int {
	return max(s.N, -s.N) + max(s.E, -s.E)
}

func part1(input string) int {
	dirs := map[int]pos{
		0: {
			N: 1, E: 0,
		},
		1: {
			N: 0, E: 1,
		},
		2: {
			N: -1, E: 0,
		},
		3: {
			N: 0, E: -1,
		},
	}

	dir := 1
	s := pos{}
	for _, line := range strings.Split(input, "\n") {
		op := line[0:1]
		num := utils.GetAllNumbers(line)[0]
		switch op {
		case "N":
			s.N += num
		case "S":
			s.N += -num
		case "E":
			s.E += num
		case "W":
			s.E += -num
		case "L":
			dir = (dir + (360-num)/90) % 4
		case "R":
			dir = (dir + num/90) % 4
		case "F":
			s.N += dirs[dir].N * num
			s.E += dirs[dir].E * num
		}
	}
	return pos{N: s.N, E: s.E}.manhatten()
}

func part2(input string) int {
	s := pos{}
	w := pos{E: 10, N: 1}

	rotateRight := func() {
		w.E, w.N = w.N, -w.E
	}

	for _, line := range strings.Split(input, "\n") {
		op := line[0:1]
		num := utils.GetAllNumbers(line)[0]
		switch op {
		case "N":
			w.N += num
		case "S":
			w.N += -num
		case "E":
			w.E += num
		case "W":
			w.E += -num
		case "L":
			num = 360 - num
			fallthrough
		case "R":
			for i := 0; i < num/90; i++ {
				rotateRight()
			}
		case "F":
			s.N += w.N * num
			s.E += w.E * num
		}
	}
	return s.manhatten()
}
