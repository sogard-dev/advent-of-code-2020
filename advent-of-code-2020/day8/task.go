package day8

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type machine struct {
	ins int
	acc int
}

type singleArg struct {
	val int
}

type acc singleArg
type jmp singleArg
type nop singleArg

func (a acc) execute(m *machine) {
	m.acc += a.val
	m.ins += 1
}

func (j jmp) execute(m *machine) {
	m.ins += j.val
}

func (n nop) execute(m *machine) {
	m.ins += 1
}

type instruction interface {
	execute(m *machine)
}

func parse(input string) []instruction {
	parseLine := func(line string) instruction {
		spl := strings.Split(line, " ")
		op := spl[0]
		num := utils.GetAllNumbers(spl[1])[0]
		switch op {
		case "jmp":
			return jmp{val: num}
		case "acc":
			return acc{val: num}
		case "nop":
			return nop{val: num}
		}
		return nil
	}

	i := []instruction{}
	for _, line := range strings.Split(input, "\n") {
		i = append(i, parseLine(line))

	}
	return i
}

func runUntilSeen(m *machine, instructions []instruction) {
	seen := map[int]bool{m.ins: true}
	for m.ins < len(instructions) {
		instructions[m.ins].execute(m)
		if _, exist := seen[m.ins]; exist {
			break
		}
		seen[m.ins] = true
	}
}

func part1(input string) int {
	m := machine{}
	runUntilSeen(&m, parse(input))
	return m.acc
}

func part2(input string) int {
	org := parse(input)
	for idx, ins := range org {
		instructions := append([]instruction{}, org...)
		switch inner := ins.(type) {
		case jmp:
			instructions[idx] = nop(inner)
		case nop:
			instructions[idx] = jmp(inner)
		}
		m := machine{}
		runUntilSeen(&m, instructions)
		if m.ins == len(org) {
			return m.acc
		}
	}

	return -1
}
