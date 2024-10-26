package day20

import (
	"math"
	"strings"

	"github.com/sogard-dev/advent-of-code-2020/utils"
)

type tile struct {
	id      int
	pattern *[][]string
}

func (t tile) top() string {
	return strings.Join((*t.pattern)[0], "")
}
func (t tile) bottom() string {
	return strings.Join((*t.pattern)[len(*t.pattern)-1], "")
}
func (t tile) left() string {
	s := len(*t.pattern)
	p := make([]string, s)
	for r := range s {
		p[r] = (*t.pattern)[r][0]
	}
	return strings.Join(p, "")
}
func (t tile) right() string {
	s := len(*t.pattern)
	p := make([]string, s)
	for r := range s {
		p[r] = (*t.pattern)[r][s-1]
	}
	return strings.Join(p, "")
}
func (t tile) rotateClockwise() tile {
	s := len(*t.pattern)
	p := make([][]string, s)
	for i := range s {
		p[i] = make([]string, s)
	}

	for r := range s {
		for c := range s {
			p[r][c] = (*t.pattern)[c][s-r-1]
		}
	}

	return tile{
		id:      t.id,
		pattern: &p,
	}
}
func (t tile) flipVertical() tile {
	s := len(*t.pattern)
	p := make([][]string, s)
	for i := range s {
		p[i] = make([]string, s)
	}

	for r := range s {
		for c := range s {
			p[r][c] = (*t.pattern)[r][s-c-1]
		}
	}

	return tile{
		id:      t.id,
		pattern: &p,
	}
}

type pos struct {
	x, y int
}

func parseTiles(input string) []tile {
	tileBlocks := strings.Split(input, "\n\n")

	tiles := []tile{}
	for _, tileBlock := range tileBlocks {
		tileNumber := utils.GetAllNumbers(tileBlock)[0]

		tileLines := strings.Split(tileBlock, "\n")[1:]
		pattern := [][]string{}
		for _, s := range tileLines {
			pattern = append(pattern, strings.Split(s, ""))
		}
		tiles = append(tiles, tile{
			id:      tileNumber,
			pattern: &pattern,
		})
	}
	return tiles
}

func createTileMap(tiles []tile) map[pos]tile {
	unusedTiles := map[tile]bool{}
	for i, tile := range tiles {
		if i > 0 {
			unusedTiles[tile] = true
		}
	}

	tileMap := map[pos]tile{}
	tileMap[pos{x: 0, y: 0}] = tiles[0]
	tileMapTodo := map[pos]bool{}
	tileMapTodo[pos{x: 0, y: 0}] = true

	testAndApply := func(thisPos pos, orgUnusedTile tile, pattern string, selector func(tile) string) {
		doWork := func(sourceTile tile) {
			t := sourceTile
			for r := 0; r < 4; r++ {
				t = t.rotateClockwise()
				if pattern == selector(t) {
					tileMap[thisPos] = t
					delete(unusedTiles, orgUnusedTile)
					tileMapTodo[thisPos] = true
				}
			}
		}

		doWork(orgUnusedTile)
		doWork(orgUnusedTile.flipVertical())
		doWork(orgUnusedTile.rotateClockwise().flipVertical())
	}

	for len(tileMapTodo) > 0 {
		for p := range tileMapTodo {
			sourceTile := tileMap[p]
			for orgUnusedTile := range unusedTiles {
				testAndApply(pos{x: p.x, y: p.y - 1}, orgUnusedTile, sourceTile.top(), func(t tile) string { return t.bottom() })
				testAndApply(pos{x: p.x, y: p.y + 1}, orgUnusedTile, sourceTile.bottom(), func(t tile) string { return t.top() })
				testAndApply(pos{x: p.x + 1, y: p.y}, orgUnusedTile, sourceTile.right(), func(t tile) string { return t.left() })
				testAndApply(pos{x: p.x - 1, y: p.y}, orgUnusedTile, sourceTile.left(), func(t tile) string { return t.right() })
			}

			delete(tileMapTodo, p)
		}
	}

	minX := math.MaxInt
	minY := math.MaxInt

	for p := range tileMap {
		minX = min(minX, p.x)
		minY = min(minY, p.y)
	}

	if minX != 0 || minY != 0 {
		newTileMap := map[pos]tile{}
		for p, t := range tileMap {
			newTileMap[pos{
				x: p.x - minX,
				y: p.y - minY,
			}] = t
		}
		tileMap = newTileMap
	}

	return tileMap
}

func part1(input string) int {
	tiles := parseTiles(input)
	mapTileWidth := int(math.Sqrt(float64(len(tiles))))
	tileMap := createTileMap(tiles)

	tl := tileMap[pos{x: 0, y: 0}].id
	tr := tileMap[pos{x: mapTileWidth - 1, y: 0}].id
	bl := tileMap[pos{x: 0, y: mapTileWidth - 1}].id
	br := tileMap[pos{x: mapTileWidth - 1, y: mapTileWidth - 1}].id

	return tl * bl * tr * br
}

func part2(input string) int {
	tiles := parseTiles(input)
	tileMap := createTileMap(tiles)
	mapTileWidth := int(math.Sqrt(float64(len(tiles))))
	tileLetterWidth := len((*tiles[0].pattern)[0])

	bigMapWidthPerTile := (tileLetterWidth - 2)
	bigMapWidth := mapTileWidth * bigMapWidthPerTile
	bigMap := make([][]string, bigMapWidth)
	for p := range bigMapWidth {
		bigMap[p] = make([]string, bigMapWidth)
	}
	for tx := range mapTileWidth {
		for ty := range mapTileWidth {
			t := tileMap[pos{x: tx, y: ty}]
			print(t.id, " ")
			p := *t.pattern
			for lx := 1; lx < tileLetterWidth-1; lx++ {
				for ly := 1; ly < tileLetterWidth-1; ly++ {
					bx := tx*bigMapWidthPerTile + lx - 1
					by := ty*bigMapWidthPerTile + ly - 1
					bigMap[by][bx] = p[ly][lx]
				}
			}
		}
		println()
	}

	seamonster := strings.Split(`                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `, "\n")

	seamonsterIndex := []pos{}
	for y, s := range seamonster {
		for x, v := range s {
			if v == '#' {
				seamonsterIndex = append(seamonsterIndex, pos{
					x: x,
					y: y,
				})
			}
		}
	}

	bigTile := tile{
		pattern: &bigMap,
	}

	findMonster := func(t tile) int {
		ret := map[pos]bool{}

		w := *t.pattern
		for y := range bigMapWidth {
			for x := range bigMapWidth {
				seamonster := true
				for _, m := range seamonsterIndex {
					wx := x + m.x
					wy := y + m.y
					if wx >= bigMapWidth || wy >= bigMapWidth || w[wx][wy] != "#" {
						seamonster = false
						break
					}
				}

				if seamonster {
					for _, m := range seamonsterIndex {
						wx := x + m.x
						wy := y + m.y
						ret[pos{
							x: wx,
							y: wy,
						}] = true
					}
				}
			}
		}

		return len(ret)
	}

	tryRotation := func(org tile) int {
		found := 0

		t := org
		for i := 0; i < 4; i++ {
			found = max(found, findMonster(t))
			t = t.rotateClockwise()
		}

		return found
	}

	hashes := 0
	for r := range bigMapWidth {
		for c := range bigMapWidth {
			v := (*bigTile.pattern)[r][c]
			if v == "#" {
				hashes += 1
			}
			print(v)
		}
		println()
	}

	try1 := tryRotation(bigTile)
	try2 := tryRotation(bigTile.flipVertical())
	try3 := tryRotation(bigTile.rotateClockwise().flipVertical())

	return hashes - max(try1, try2, try3)
}
