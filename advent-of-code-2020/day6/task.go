package day6

import (
	"strings"
)

func part1(input string) int {
	sum := 0
	for _, group := range strings.Split(input, "\n\n") {
		seen := map[rune]bool{}
		for _, line := range strings.Split(group, "\n") {
			for _, r := range line {
				seen[r] = true
			}
		}
		sum += len(seen)
	}
	return sum
}

func part2(input string) int {
	sum := 0
	for _, group := range strings.Split(input, "\n\n") {
		seen := map[rune]int{}
		members := strings.Split(group, "\n")
		mlen := len(members)
		for _, line := range members {
			for _, r := range line {
				seen[r] += 1
				if seen[r] == mlen {
					sum += 1
				}
			}
		}
	}
	return sum
}
