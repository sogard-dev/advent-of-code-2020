package day13

import (
	"math"
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

	d := 0
	m := 1
	for _, e := range numbers {
		for {
			d += m
			if (d+e.offset)%e.num == 0 {
				m = m * e.num
				break
			}
		}
	}

	return d
}
