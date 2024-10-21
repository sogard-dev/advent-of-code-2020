package day13

import (
	"math"
	"slices"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string) int {
	nums := utils.GetAllNumbers(input)
	earliest := nums[0]
	busses := nums[1:]

	waitingTime := math.MaxInt
	closestBus := 0
	for _, bus := range busses {
		time := bus - (earliest % bus)

		if time < waitingTime {
			waitingTime = time
			closestBus = bus
		}
	}

	return waitingTime * closestBus
}

type number struct {
	num    int
	offset int
}

func part2(input string) int {
	line := strings.Split(input, "\n")[1]
	entries := strings.Split(line, ",")
	numbers := []number{}
	for i, e := range entries {
		n := utils.GetAllNumbers(e)
		if len(n) == 1 {
			numbers = append(numbers, number{
				offset: i,
				num:    n[0],
			})
		}
	}

	slices.SortFunc(numbers, func(a, b number) int {
		return b.num - a.num
	})

	base, mult := findCommonT(numbers[0:min(4, len(numbers))])

	for i := 0; i < math.MaxInt; i++ {
		t := base + mult*i

		good := true
		for _, e := range numbers {
			if (t+e.offset)%e.num != 0 {
				good = false
				break
			}
		}

		if good {
			return t
		}
	}

	return -1
}

func findCommonT(numbers []number) (int, int) {
	biggest := number{}
	for _, e := range numbers {
		if e.num > biggest.num {
			biggest = e
		}
	}

	base := 0

	for i := 1; i < math.MaxInt; i++ {
		t := biggest.num*i - biggest.offset

		good := true

		for _, e := range numbers {
			if (t+e.offset)%e.num != 0 {
				good = false
				break
			}
		}

		if good {
			if base == 0 {
				base = t
			} else {
				return base, t - base
			}
		}
	}
	panic("donkey")
}
