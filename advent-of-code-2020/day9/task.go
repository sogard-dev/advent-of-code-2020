package day9

import (
	"math"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func part1(input string, preamble int) int {
	q := utils.NewQueue[int](preamble)
	numbers := utils.GetAllNumbers(input)

	for _, num := range numbers {
		if q.Len() == preamble && !hasPreamble(&q, num) {
			return num
		}
		q.Push(num)
	}

	return -1
}

func hasPreamble(q *utils.Queue[int], num int) bool {
	for i1 := q.Iterator(); i1.HasNext(); {
		n1 := i1.Next()
		for i2 := q.Iterator(); i2.HasNext(); {
			if n2 := i2.Next(); n1 != n2 && n1+n2 == num {
				return true
			}
		}
	}
	return false
}

func part2(input string, preamble int) int {
	num := part1(input, preamble)
	q := utils.NewUnboundedQueue[int]()
	numbers := utils.GetAllNumbers(input)

	for _, n := range numbers {
		q.Push(n)

		for queueSum(&q) > num {
			q.Pop()
		}

		s := queueSum(&q)
		if s == num {
			a, b := findMinMax(&q)
			return a + b
		}
	}

	return -1
}

func findMinMax(q *utils.Queue[int]) (int, int) {
	minRet := math.MaxInt
	maxRet := math.MinInt
	for i := q.Iterator(); i.HasNext(); {
		n := i.Next()
		minRet = min(minRet, n)
		maxRet = max(maxRet, n)
	}
	return minRet, maxRet
}

func queueSum(q *utils.Queue[int]) int {
	sum := 0
	for i := q.Iterator(); i.HasNext(); {
		sum += i.Next()
	}
	return sum
}
