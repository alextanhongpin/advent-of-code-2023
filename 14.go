package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(solve(example, 1, 1))          // 136
	fmt.Println(solve(input, 1, 1))            // 105784
	fmt.Println(solve(example, 1000000000, 4)) // 64
	fmt.Println(solve(input, 1000000000, 4))   // 91286
}

func draw(static map[Point]rune) string {
	var maxX, maxY int
	for p := range static {
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	var rows []string
	for y := 0; y <= maxY; y++ {
		var cols []string
		for x := 0; x <= maxX; x++ {
			if c, ok := static[Point{x, y}]; ok {
				cols = append(cols, string(c))
			} else {
				cols = append(cols, string('.'))
			}
		}
		rows = append(rows, strings.Join(cols, ""))
	}

	return strings.Join(rows, "\n")
}

type Point struct {
	x, y int
}

func cmp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func solve(in string, n int, iter int) (load int) {
	cache := make(map[string]int)
	loadByStep := make(map[int]int)

	var i int
	for i <= n {
		in, load = cycle(in, iter)
		if last, ok := cache[in]; ok {
			cycle := i - last
			cycleBegin := i - cycle

			left := (n - i - 1) % cycle
			return loadByStep[cycleBegin+left]
		}
		loadByStep[i] = load
		cache[in] = i
		i++
	}

	return load
}

func cycle(in string, n int) (out string, load int) {
	dir := "uldr"

	// We only use this to determine the boundary.
	grid := parseGrid(in)

	for i := 0; i < n; i++ {
		d := dir[i%len(dir)]
		static := parseStatic(in)
		rocks := parseRocks(in)

		slices.SortFunc(rocks, func(a, b Point) int {
			switch d {
			case 'u':
				// top to bottom, left to right
				if a.y == b.y {
					return cmp(a.x, b.x)
				}
				return cmp(a.y, b.y)
			case 'd':
				if a.y == b.y {
					return cmp(a.x, b.x)
				}
				// down first, then up
				return cmp(b.y, a.y)
			case 'l':
				// left first, then right
				if a.y == b.y {
					return cmp(a.x, b.x)
				}
				return cmp(a.y, b.y)
			case 'r':
				if a.y == b.y {
					// Right first, then left
					return cmp(b.x, a.x)
				}
				return cmp(a.y, b.y)
			}

			return 0
		})
		for len(rocks) > 0 {
			var keep []Point
			for _, ball := range rocks {
				x := ball.x
				y := ball.y
				switch d {
				case 'u':
					y--
				case 'd':
					y++
				case 'l':
					x--
				case 'r':
					x++
				}
				pos := Point{x, y}

				// The rock is out of bound.
				if _, ok := grid[pos]; !ok {
					static[ball] = 'O'
					continue
				}
				// There's an obstacle, the rock can no longer move.
				if _, ok := static[pos]; ok {
					static[ball] = 'O'
					continue
				}
				keep = append(keep, pos)
			}
			rocks = keep
		}

		out = draw(static)

		var maxY int
		for p := range grid {
			maxY = max(maxY, p.y)
		}
		maxY += 1

		load = 0
		for p, c := range static {
			if c != 'O' {
				continue
			}
			load += maxY - p.y
		}

		in = out
	}

	return
}

func parseStatic(in string) map[Point]rune {
	m := make(map[Point]rune)
	for y, line := range strings.Split(in, "\n") {
		for x, r := range line {
			if r == '#' {
				m[Point{x, y}] = r
			}
		}
	}
	return m
}

func parseRocks(in string) []Point {
	var rocks []Point
	for y, line := range strings.Split(in, "\n") {
		for x, r := range line {
			if r == 'O' {
				rocks = append(rocks, Point{x, y})
			}
		}
	}
	return rocks
}

func parseGrid(in string) map[Point]rune {
	m := make(map[Point]rune)
	for y, line := range strings.Split(in, "\n") {
		for x, r := range line {
			m[Point{x, y}] = r
		}
	}
	return m
}

var example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

var input = `..#..#......#...O.O#O#..#O.O#.O..O.O.#.O..OO.O......###...#..#....OO..O.#.....#.O...#.#OO#O#..O.O#O.
#.#..OOO..#..#...O......#.###O..O.O...#..........#..O......O.#O#......#....O..O#OO.........OO.##O...
..O.#....O.O...O#O...#.##O#OO.#OO.....#.O....O....#OO.O#....O.O.#.......O......O#.O.O...#.....#O.O.O
#OO.O...#.#......#.......##.O#...OO..O....#...#OO....O...#.#....O#.#.#.O.....#..O..O..#.#...O.#.O..O
O##O...O.#O....#.#O#........O....#.#.O.O#O..O...O.#OO.O###.#.#.O..........O##O#......O..##....#O##..
O..#..........O....O.O.O.OO.OO.OOOO...#..OOO..O..#O.##...#..OOOO.O####....O#....#O##.O.........O...O
.##...OO..O..#...O.O...O#.O.O.O..OO..O..##O.##...#.OOO.#O.....O...O....OOO.#O...#O#.......#.O.OO.O..
.O....O..#.OO...###....#.##...O.O#...O.#.O.#....O...#..#.#O.O.#.....O#O#O..OO..#.O.O...OO.O..##.#.O.
.#.##O#.OO....#..###.#.O...#.#O.OO.#..O..O......O...#....OO..##.#O#O.O..O.....#.#..#..O.#...#O....O.
...#.O..#OO...O#....O...#......O#.#..O.O#OOO#.#OO..O.....#.....#O.O#O...O.#.#..#O.#...##O..#..O.O#..
.##....O#.......##........#.#.O...O.O#.#.O........O......#.#O.O...O##.....OO..#O..O#.O....#...#O....
.#O..O.#..O#OO#..O.......#.O#O#...O...#O.O...OO#....#....O....#....O#O..OO....O.......#.....O.O#O##.
#....#..#.#..O#.#.OO....##...#O.O###.....O.##O#OO#......O.....O#.#O..O...#...#.O.O.O.#...#.....#.OO#
.#..OO.#...##.O...#......O...#O#.O.#O.#.#....#.#.....OO#.OO.......#O...OO.O.#....#.......#.#.##....O
...O....O.##O.O....#O...#....#OOOO......#..O..O..##O....#O.......O.#.O...O#..#O.O...#O.##..O....O...
#..#...##.##..#..O.#...#..O..#.....#...O.......#O#...O.#..#O#.#...#..O#....OO....O#..#O..O.O#.O#..#.
.OO.......##...O...#.#....#O#.#O..#.#O.#......O.O.#.......#O.O........#.#.O.#.OO.#..OO#..#O..O.#...#
...O..O..O#...#......O...#.OO.O....O.........O#O.O..#.#O..##..........O.O..#OO.O.#O..O...O...#......
...O.#...#.O##.#..OO#OO..#....O.#...#........#...O#...O..#..O..##.....O..##.O..#..#.OO...O.#.#.OO.#.
..O.#..O.O..#......O.....OO..O..OO.O.....O....#....##....OO....#..O#....O....#O#....#...##.....O#..#
OO.....#.##.OO#........O..##.....O.#.O#.....#.OO.OO.O...O.OO...#....#.#.O#.#...#OO...#..O#...O#O.O..
##...O#.#.O............O....#OOOOOO.....#.O#OO#O.OO....O#.#.O............#...OO.....#.....O#.....O..
..O...#....O##.O#.#.O#O...O..O...O.O....O..#..#...OO.O.OO..O.#.#..#..O#.....O.##........O......O....
...O.#.O...O.##..#.....O.#.........OO#....O......#OOOOO..........##.O......#.......##.##..O.#.....O#
#...#..#O..#.O.#....O..#.O.OOO#..O...#.#..O.OO.#...##..O..#.#..##.OOO....#..O.O.O.#.OO.O.O#..O#.....
..O.#......#..#.##.#.#.O.#.#.....O..#..O.#O#.#..#...O.O.O#.##....#.O.O.O....O..#.O.O......O..O......
..#O......O#.....#..#.#.....#.#.#.....#..O..........O...#.##.......O.#.....O...O##..O.O...#.....O.OO
..O...O.##.......#...#...#...O.........#...O...#.O.O....#.OO..#..OO...O...OOO#O...OO.O#..O.OO.....#O
.....O..#..OOO......##O..O.OO.#....O..O.##.#O..#..O.OOO...O#...OO.O..#.#.O#....#O..O..##.O....O.O..#
O#..O..........O....O##..OO....O.O..O.OOOO....O.......#...#O.#.###O.............#....O.O##..#O.....#
......#.OO#....O..O.........O..#.#O#O....O....OO#O..#.....#.O.#.#.O.OO...O.........#.##..OOO..#...#.
...#..O.....O##.....#O.O.O#....OOO.OO..O.OO.....#O.##O#.#...O.#..#....#..#..O.O...#O.O.##...O###O.#.
...O.#O#....#.#O..O....O.O..#...#.#..#......#....O.....#......##...O....#.#...#..O.....#....O#.....O
O.O#..#.O.O.....#..#O.....#..#O....#....O.O.O..#..O###O.#..O#...O##.O.......#........O..O#..OO..#...
#O..O..##OO....OO.O.....#...O...O...#..O#.#.......#O.OO##.....O.O...O...OOOO...OO....O......O.#....#
OOOO....O..##..OO....O#..##..OOO.#..O.OO.O#O....O.#.#.....OOO.##....O.O.O..#OO..O...#...#O.O.OO...O.
.#..OO.O...##..#......O..........#.O.....#........###..O#..#.O.O#.....O..O..O.#...O....#.O..O......O
....O.#O..#O.#..O...###..OO...O.OOO..#...O...OO.....#.O.#O..O........#..#...##.OOO.OO..#..O...#....O
..OO.....O#........OO#O.O.........#..#.....#.......#.O..#......#O#.O.O.O..OO..O.......O...O......#O.
##....O.....#..#.OOO..O.#.O.###..O...OOOO.#.......O..#...O.O#OO.O#...O....#.O.O..#.....O...O#...#..#
.O#..#..O..#O#.#.........O......O.O##...#..#...O....O..O.#....#...O..OOO..O....#..#.OO.O#.....#....#
.O#.OO.#.....#.O.#.O##O#......O......O.#O..#O#....#..OO.OO#.O.O.O..#OO#.#........###..........##....
...#O#....#....OO.....O.#..O..O....#..OOO##.......O......##..#...O#...#O#........O.....#..#.O.#.#.O#
.O.O....O.O....OO#.#...O#.OO..O.#...#...#...#.#.......#..#.O..#...O#..O...O..#..#.O.O.....O..#O.O..#
....O.O#...OO##OO....O.OO...O#...#.....O....OO..O.O...#....OO....O#.#.OO.O...........#.#.O..O.#.....
O.O.#.......#O.#O.....#...#OO....#.OO...#O.#..OO#.O....#.O#..........#OO#.......#O..#...#....O#.....
#.#..#.O...O...OOO.#.#..O....O..##O.O......#.O.#......O....O..#O......#O..#..O..O...O#..O......#O#.O
#.......#......O........O...##..#.#.O...O.......#..O......O#O.........#OO#...O.#...O..#.....#...#...
#...#..#.....O#.O#O..O#......#O...O..##....##....####..O......O..##..OO...........O#.......OO..O##..
..#O.#O...O.#....O#...#...#....#...#.......O.O.O..#...#.O.O....O..O.....O#O....###..O..O.#.OO##..O#.
O......#..........#.#..O###...O...#.#........#.O..O......#O#O.......#.#.O.#..OO#O.O..##.#.....O..O.O
.O.O.O#.O...O...O..#O.....O.#OOO..O.O.....O.#.#.#.#...#.......#O........O.#O.O...#.O.#..O...O.#O...O
.O....#O.O.#...###...O........O#O..#.O.OO#..OO...O.O......#..O.##..O#.#.#.#...O#O.......#..O.#.#O..#
..#...O..###O#..O..#...O...#OO...O.#..#.O.........OOOO.O...........#..O......#...#.#...##...O.......
....O#.OO...#....O#...#O.......OO.OOO..O......#......#...........####...O.##O.#.....O......#.......O
O...O..O.O.#O.##...OO#.O#O.#...#.#...OO...#O.O.OO.....#.................##.#..O.O.O..O.O.#.......O.O
O.O.#..#...O.O....O........##.#.....OO....O..O.......#..#O.#OO.OO.....##....O.O##....#.....#O.....OO
O.#.....#.......#O..O#......O..O..O#.O..O...#.O...##O#....O.#O....#O#...O......OO.O..O.#O#.#O..#....
#O.O#....#.....#O........#........#..O.OOO.O......#O#.....#.##.....##.#OO#.O......OO.OOOO.#.O#....O.
O.............#.....OO..OO.#.O#.#.......OO.......O#....#.#.#..........O..#...O.....O#..O...O.O#...#.
O.#O#.........#.#..O#O#..OOO#O#...#....##.#.O.##.O....OO....O......O.O#.O...O#O.#.##..........#OOO.O
...O.......OO...O#O#..O...O...O..OO..O..#..O.........##........O#....O..#.OO...OO...OO..#.##O#....#.
#....#O....O.#..O...O..#O#.....#OO.O#..O..#.....#..##........O..#O#O..........O....#OO..O#...O.O...O
O........O.O.O#..O....O###......O..##O.....O.....#O..O.#.....##.....#OO...OO...#..O#.#...#O.OO##.O..
.#O.........##.......#O......O#..OOOOOO#O....#..O#.#O#.......O.O...O...O#...#....O....#..#...#O.#.OO
.O.O...#.....OO.#.O..#OO.....#.........O....OO.....#.#.#.#.#..O...##...O..O...OO....#.O.......OO....
##OO.OOO#.#O#..O....O.O.O...O...O#..#O...O..#......OO.O..#...#O..O.#.OOOO.O...O......#.##O.#O.#...##
OO...O....#.#...#.#.O.#..#..O.O.O..#...#......O.......O......O.O.O......O.....OO##.#..O.............
....#.O.#......#..OO......O#.O.#...O.O.#.O...#O.....#..#............O..#O#.#.O..OO...O.OO..O.#...OO#
O#O..#O...#..OO............#.O..#..O.##OO.....OO.##..O.......O.#.O.....O...OO...OO.O.#......O...O#..
OO.##.#OO#O.......OO.O...#.OOO...O#....O.O#.O....OO.#..O..OO#OO.#....#..O##O.O.O....OOOO...###..#..O
....OO...#..O..........O.....O..O#.O.......##O.##OOO..OOO...O#.O.O...O..O#O..O.#O.#..O#.....O...OOO.
#..OOOO.#O......#...........O.##..##....#...O##O#......O.O..#...#O#....OO.#..OO..#..OO#.O...##......
.#O.....#O..O.#O.O.##.#.....OO....O#..##....O#..O.....#...#.O..OOO...##....OO...O..O...OO...OO...#O.
....O..O#.O#....OO#..#.#....OO##.O#O.##.....O........O#O...#.........O.....#.#.O.O...OO...O....O#.#.
...........O#OO#...#O.O..OO..O...O.....O.O..#.............O..O............#.O..O#...........#..#....
O.#.O........O.O..#..#O.#O.#......#.O#.O##..#O#O..O#O....#.##..#O....#........O.....#...#...O.......
.#.#O.O....OO.........O..#.....OOO.#...##.O.O.....#.OO..#...O.O........O#...O...O..#O...O.#......O..
....O..O...#O.O.OO.O.#...O#...O..#..O...#O#.#.....OO#..O..#...O#.O..#OO..#...OO.OO.#.#.....#....#.O#
..OOO.#......O..O.......#.......#...#....O.O#...O#.OO...#..#OO.O....#.#.#..O....O.....O.O.O#...O.O..
O......O.#.O.OO..O..#..O..#.....OO.O.OOOO#..O#....O.#..#...#.##.OO.....O.O.#.#.#.O.#O...OO.O.#.#...O
O...OO..###.O##........O..O.#.....OO..#....#.#.#..#.#O...O...###...#O..#O........#..#O.#.....###.OOO
.O..###O........OO###......#....##.....O.OO.....O.#...#.....O##O..OO.#..O...O#....OO#....O..#..#.OO.
.O..O..O#......O..O...###.....O..#.#....#........#....#.#..O#...O.#.#####O#O.......#O..O..##..O..O..
.OOO.OO.O###...........#.......O.O......O.....O...O.....OO.O#............#.OO.........OO...O#......#
....#.O#...#.#..OOO.#O.O.O...OO....#...##..O..OOO.......OOO....O#.O....O..O.###..O..O..#...##..O....
.......#O#.O..#...O#.O.O..O..O..O.O...OO.....O.O......#......#..##O.....O##...O...#.O..##..##.O...#O
OOOO..##............OO..O.....O...O.#.....#..#...O.#.O..O.#.##.....O..#..#...#O..OO....O..O...O.....
O..#OO.OO.....OO.O........#...O.O...OO.#.#......#....O...#OO...OO..#.O...#.O.......O...O...O#O##.O##
O.O.O.OO...##....#.......#...#......#..#O.......##.#.#O..#OO......O......#....OO...OO.#.....O.O..##.
O..O......##.....O...#...#.OO..OO.......###.#.......#.....O#OOO..#O.OO#......O..O.#.O..O..OO...#..#.
.....O...#.##.#.##...O#.#..O..OO.#O##O....O#O.O..OO.#O....O....#OO..O.O...OO...O....#....O.##O#...O.
.#.#....#..OOO.....#.....##.##..........#.O...##..O.#O.......OOOO#O.O.O#.O..O..O....O.....#......OO.
O..O##OO#....#O.....OO.O.O.##O.O............OOO..O#.....O...#O...#...........#OO..#.##.#.#.....O#...
.##.OO......#..##..O..OO...O..#...O..#O#.OO.O.OOO...#.OO..O.OO#..O.......#..#..##...O.O##O..OO...O..
......#....O..#O.#.O.#.#O.O......#.O.....O##.O......#....#.###O.O#..O#O.....O..O..O..##..#O#..O..O..
#O....#..O.#OO....#.....#....O.#...#.#..O..#.##..#.O.O.O......#.......#OO..O.#.O#OO.##.O.......OO..#
##....#...#.O..#.#..OOO...O...O..##...O.O..#..#...#.....O#.#O...O..#.#.....O....O#......O#O.O...O#O#
#..........O#...#.O#O.O.O.#..#O.O..#.#O.#.O......O.#......#.......O..#.#..#..##O..O.#...#..O#.##.O.O
O..#.O...O.........#...#........#.O..OO..O.#O.##.O..O..OO#.OO#O..OO...O#.O.#..#.#O#.#......O.#...#.O`
