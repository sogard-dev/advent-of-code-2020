package day23

import (
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

func parse(input string) (circle, int) {
	c := newCircle()
	first := -1
	for _, line := range strings.Split(input, "") {
		n := utils.GetAllNumbers(line)[0]
		if first == -1 {
			first = n
		}
		c.insert(n)
	}
	return c, first
}

func part1(input string, moves int) string {
	cubs, first := parse(input)
	solve(&cubs, first, moves)

	ret := cubs.removeAfter(1, cubs.size-1)

	str := ""
	for _, v := range ret {
		str = str + strconv.Itoa(v)

	}

	return str
}

func part2(input string) int {
	cubs, first := parse(input)
	for c := cubs.size + 1; c <= 1000000; c++ {
		cubs.insert(c)
	}

	solve(&cubs, first, 10000000)
	ret := cubs.removeAfter(1, 2)
	return ret[0] * ret[1]
}

func solve(cubs *circle, first int, moves int) {
	currentCub := first
	TO_TAKE := 3
	MAX := cubs.size
	for move := range moves {
		if move%10000 == 0 {
			println("Move ", move)
		}

		pickUp := cubs.removeAfter(currentCub, TO_TAKE)

		destination := currentCub - 1
		for !cubs.insertAfter(destination, pickUp) {
			destination--
			if destination < 1 {
				destination = MAX
			}
		}

		currentCub = cubs.valueAfter(currentCub)
	}
}

func newCircle() circle {
	return circle{
		lookup: map[int]*node{},
	}
}

type circle struct {
	head   *node
	size   int
	lookup map[int]*node
}

func (c *circle) insert(v int) {
	n := node{
		v: v,
	}
	if c.head == nil {
		c.head = &n
		n.next = &n
	} else {
		n.next = c.head.next
		c.head.next = &n
		c.head = &n
	}
	c.size++
	c.lookup[v] = &n
}

func (c *circle) removeAfter(v int, occurences int) []int {
	ret := make([]int, occurences)
	n := c.lookup[v]

	nextHead := n
	for r := 0; r < occurences; r++ {
		nextHead = nextHead.next
		ret[r] = nextHead.v
		delete(c.lookup, nextHead.v)
	}
	n.next = nextHead.next
	c.head = n
	c.size -= occurences
	return ret
}

func (c *circle) insertAfter(v int, values []int) bool {
	if n, exist := c.lookup[v]; exist {
		insertAfter := n
		for _, toInsert := range values {
			newNode := node{
				v:    toInsert,
				next: insertAfter.next,
			}
			insertAfter.next = &newNode
			insertAfter = &newNode
			c.lookup[toInsert] = &newNode
		}
		c.size += len(values)

		return true
	}

	return false
}

func (c *circle) valueAfter(v int) int {
	n := c.lookup[v]
	return n.next.v
}

func (c *circle) print() {
	var visited *node
	for n := c.head; n.next != nil; n = n.next {
		if visited == nil {
			visited = n
		} else if visited == n {
			break
		}

		print(" ", n.v, " ")
	}
}

type node struct {
	v    int
	next *node
}
