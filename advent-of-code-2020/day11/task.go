package day11

import (
	"strings"
)

type binaryGrid struct {
	rows    int
	columns int
	grid    []bool
}

type rc struct {
	r int
	c int
}

func newGrid(rows, columns int) binaryGrid {
	return binaryGrid{
		rows:    rows,
		columns: columns,
		grid:    make([]bool, rows*columns),
	}
}

func (b *binaryGrid) set(row, column int, on bool) {
	i := column + row*b.columns
	b.grid[i] = on
}

func (b *binaryGrid) get(row, column int) bool {
	i := column + +row*b.columns
	return b.grid[i]
}

func (b *binaryGrid) countTrue() int {
	trues := 0
	for i := range len(b.grid) {
		if b.grid[i] {
			trues += 1
		}
	}

	return trues
}

func part1(input string) int {
	return solve(input, true)
}

func part2(input string) int {
	return solve(input, false)
}

func solve(input string, part1 bool) int {
	lines := strings.Split(input, "\n")
	rows := len(lines)
	columns := len(lines[0])

	pattern := []rc(nil)
	patternGrid := newGrid(rows, columns)
	for r, line := range lines {
		for c, letter := range line {
			switch letter {
			case 'L':
				pattern = append(pattern, rc{
					r: r,
					c: c,
				})
				patternGrid.set(r, c, true)
			}
		}
	}

	current := newGrid(rows, columns)
	next := newGrid(rows, columns)

	changed := true
	for changed {
		if part1 {
			changed = iteratePart2(pattern, &patternGrid, &current, &next, 4, 1)
		} else {
			changed = iteratePart2(pattern, &patternGrid, &current, &next, 5, max(patternGrid.columns, patternGrid.rows))
		}
		current, next = next, current
	}
	return next.countTrue()
}

func iteratePart2(pattern []rc, patternGrid *binaryGrid, currentGrid *binaryGrid, next *binaryGrid, limit int, maxFromChar int) bool {
	changed := false
	for _, rc := range pattern {
		r := rc.r
		c := rc.c

		adjacent := 0

		for dr := -1; dr < 2; dr++ {
			for dc := -1; dc < 2; dc++ {
				if dr != 0 || dc != 0 {
					for k := 1; k <= maxFromChar; k++ {
						nr := r + dr*k
						nc := c + dc*k
						if nr < 0 || nc < 0 || nr >= patternGrid.rows || nc >= patternGrid.columns {
							break
						}
						if currentGrid.get(nr, nc) {
							adjacent += 1
							break
						} else if patternGrid.get(nr, nc) {
							break
						}
					}
				}
			}
		}

		v := false
		if currentGrid.get(r, c) {
			if adjacent < limit {
				v = true
			} else {
				changed = true
			}
		} else {
			if adjacent == 0 {
				v = true
				changed = true
			}
		}
		next.set(r, c, v)
	}

	return changed
}
