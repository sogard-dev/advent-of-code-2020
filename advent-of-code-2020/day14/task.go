package day14

import (
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type machine struct {
	memory  map[uint64]uint64
	mask    string
	bitsize int
	v1      bool
}

func numToPaddedRune(value uint64, bitsize int) []rune {
	asStr := strconv.FormatUint(value, 2)
	asStr = strings.Repeat("0", bitsize-len(asStr)) + asStr
	return []rune(asStr)
}

func replaceRunes(runes []rune, chars string, mask string) {
	for i, r := range mask {
		if strings.ContainsRune(chars, r) {
			runes[i] = r
		}
	}
}

func replaceValueByMask(mask string, value uint64, bitsize int) uint64 {
	out := numToPaddedRune(value, bitsize)
	replaceRunes(out, "01", mask)
	n, ok := strconv.ParseUint(string(out), 2, 64)
	if ok != nil {
		panic("ohh noes")
	}
	return n
}

func setFloatingMemory(m *machine, out []rune, idx int, value uint64) {
	if idx == m.bitsize {
		n, ok := strconv.ParseUint(string(out), 2, 64)
		if ok != nil {
			panic("ohh noes")
		}
		m.memory[n] = value
	} else if out[idx] == 'X' {
		out[idx] = '0'
		setFloatingMemory(m, out, idx+1, value)
		out[idx] = '1'
		setFloatingMemory(m, out, idx+1, value)
		out[idx] = 'X'
	} else {
		setFloatingMemory(m, out, idx+1, value)
	}
}

func (m *machine) set(address, value uint64) {
	if m.v1 {
		m.memory[address] = replaceValueByMask(m.mask, value, m.bitsize)
	} else {
		out := numToPaddedRune(address, m.bitsize)
		replaceRunes(out, "X1", m.mask)
		setFloatingMemory(m, out, 0, value)
	}
}

func (m *machine) setMask(mask string) {
	m.mask = mask
}

func (m *machine) sum() uint64 {
	s := uint64(0)
	for _, v := range m.memory {
		s += v
	}
	return s
}

func newMachine(bitsize int, part1 bool) machine {
	return machine{
		mask:    strings.Repeat("X", bitsize),
		bitsize: bitsize,
		memory:  map[uint64]uint64{},
		v1:      part1,
	}
}

func part1(input string) uint64 {
	m := newMachine(36, true)
	return solve(input, m)
}

func part2(input string) uint64 {
	m := newMachine(36, false)
	return solve(input, m)
}

func solve(input string, m machine) uint64 {
	for _, line := range strings.Split(input, "\n") {
		spl := strings.Split(line, " = ")
		if spl[0] == "mask" {
			m.setMask(spl[1])
		} else {
			nums := utils.GetAllNumbers(line)
			address, value := nums[0], nums[1]
			m.set(uint64(address), uint64(value))
		}
	}
	return m.sum()
}
