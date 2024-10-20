package day10

import (
	"slices"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string) int {
	nums := utils.GetAllNumbers(input)
	slices.Sort(nums)

	jumps := make([]int, 4)
	prev := 0
	for _, n := range nums {
		jumps[n-prev] += 1
		prev = n
	}
	jumps[3] += 1
	return jumps[1] * jumps[3]
}

func part2(input string) int {
	nums := utils.GetAllNumbers(input)
	slices.Sort(nums)

	return count(nums, 0, 0, map[entry]int{})
}

type entry struct {
	idx        int
	lastPicked int
}

func count(nums []int, idx int, lastPicked int, cache map[entry]int) int {
	device := nums[len(nums)-1]
	if lastPicked == device {
		return 1
	}
	if idx >= len(nums) {
		return 0
	}
	n := nums[idx]
	if n-lastPicked > 3 {
		return 0
	}

	e := entry{idx: idx, lastPicked: lastPicked}
	if p, exist := cache[e]; exist {
		return p
	}

	options := 0
	options += count(nums, idx+1, n, cache)
	options += count(nums, idx+1, lastPicked, cache)
	cache[entry{idx: idx, lastPicked: lastPicked}] = options
	return options
}
