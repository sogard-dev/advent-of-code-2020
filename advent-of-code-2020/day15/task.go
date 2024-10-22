package day15

import (
	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string) int {
	return solve(input, 2020)
}

func part2(input string) int {
	return solve(input, 30000000)
}

func solve(input string, turns int) int {
	numbers := utils.GetAllNumbers(input)
	memory := map[int]*utils.Queue[int]{}

	addToMemory := func(idx, val int) {
		if _, exist := memory[idx]; !exist {
			q := utils.NewQueue[int](2)
			memory[idx] = &q
		}
		memory[idx].Push(val)
	}

	previous := 0
	for i, n := range numbers {
		addToMemory(n, i+1)
		previous = n
	}

	for turn := len(numbers) + 1; turn <= turns; turn++ {
		if entry, exist := memory[previous]; exist && entry.Len() == 2 {
			it := entry.Iterator()
			a := it.Next()
			b := it.Next()
			previous = b - a
		} else {
			previous = 0
		}

		addToMemory(previous, turn)
	}

	return previous
}
