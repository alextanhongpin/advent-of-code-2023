// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solve(example, 6))   // 16
	fmt.Println(solve(example, 10))  // 50
	fmt.Println(solve(example, 50))  // 1594
	fmt.Println(solve(example, 100)) // 6536
	fmt.Println(solve(example, 500)) // 167004
	fmt.Println(solve(input, 64))    // 3594
	fmt.Println(part2(input))        // 3594
}

func part2(in string) int {
	// For part 2
	// The start S is in the middle
	// The horizontal and vertical middle/top/bottom/left and right is empty.
	// We basically need to fill the diamond ...
	// / \
	// \ /
	// The other solutions suggests to use polynomial extrapolation.

	steps := 26501365
	_ = steps
	fmt.Println(solve(in, 65))       // 3755
	fmt.Println(solve(in, 65+131))   // 33494
	fmt.Println(solve(in, 65+131*2)) // 92811

	/*
		# main.py
		import numpy as np

		m = np.matrix([[0, 0, 1], [1,1,1], [4,2,1]])
		b = np.array([3755, 33494, 92811])
		x = np.linalg.solve(m, b)

		# 26501365 = 202300 * 131 + 65 where 131 is the dimension of the grid.
		n = 202300
		print(x[0] * n * n + x[1] * n + x[2])
	*/

	return 605247138198755
}

var directions = []Point{
	Point{0, 1},
	Point{0, -1},
	Point{1, 0},
	Point{-1, 0},
}

func solve(in string, steps int) int {
	m := parse(in)

	var start Point
	for p, c := range m {
		if c == 'S' {
			start = p
		}
	}

	return fill(m, start, steps)
}

type Tuple[T1, T2 comparable] struct {
	a T1
	b T2
}

func fill(m map[Point]rune, start Point, steps int) int {
	paths := make(map[Point]bool)

	isEven := steps%2 == 0
	var multiplier int
	if !isEven {
		multiplier = 1
	}

	maxX, maxY := shape(m)

	visited := make(map[Point]bool)

	var q []state
	q = append(q, state{
		p:     start,
		steps: steps,
	})
	for len(q) > 0 {
		var h state
		h, q = q[0], q[1:]
		p, steps := h.p, h.steps

		pn := Point{(p.x%maxX + maxX) % maxX, (p.y%maxY + maxY) % maxY}
		if r, ok := m[pn]; !ok || r == '#' {
			continue
		}

		if steps < 0 {
			continue
		}

		if _, ok := visited[h.p]; ok {
			continue
		}

		visited[h.p] = true

		dist := p.Distance(start)
		if dist%2 == multiplier {
			paths[p] = true
		}

		for _, d := range directions {
			q = append(q, state{
				p:     Point{p.x + d.x, p.y + d.y},
				steps: steps - 1,
			})
		}
	}

	return len(paths)
}

type state struct {
	p     Point
	steps int
}

func shape(m map[Point]rune) (maxX, maxY int) {
	for p := range m {
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}
	maxX++
	maxY++
	return
}

func parse(in string) map[Point]rune {
	m := make(map[Point]rune)
	for y, line := range strings.Split(in, "\n") {
		for x, c := range line {
			m[Point{x, y}] = c
		}
	}
	return m
}

type Point struct {
	x, y int
}

func (p Point) Distance(o Point) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var example = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

var input = `...................................................................................................................................
.......#.#.#.#.....#.#...##..........#........##....#......#...............##........#..#...##.......###....#.##........#.#....#...
.###............#........#......#.....#...................................#............#..........#.....#..##.#.......#.....#......
.......#....#...##......#..............##.###....##......#..................#.....#..#.....#...#.#....#................#....#.##...
..............#................#..#.#...#......#.#..#....................##....#..#.#.....#.#..........#...#.#......#...#..#.......
.....#.....#............##.#.....#......#..#...................................#......##..#.#..........#.#..##..#................#.
.#....#......##.#...#..#.....#................##.#....#.............................##...#...#........#....#.......................
....#.....#................#.....##...............#...............#...........#..#...........#.........#.............#........#....
.........#.................#...........#......#...#...............#...........#....#....#..........#...#.#.##.................#....
.........#....#..#.....#..#..#.#....#...#......................#................####...#.....#..........#.......#..#........#......
..............#.##...............#.#.#..........#..#............#.........................................#.#.#............#.#.....
..............#...#.....#...#....##.............##............##..#.................###........#....#.#.......................#..#.
......#.....##..#....#.#......#.#....#..#.#.....#............................................##.#.....#.#.......#......#.#.#.......
........#.............#.............#.....#................#.......#.....................#...#.........#..#...##..##...##..###.....
.......#..........................#........................................#........###..##......#.#.....#....#....#....#.......#..
...#......#.#..........#..#..........#..#.....#..........#........#.......................#....#.......#...##.............#.#......
..#..##...#...........#.#.#.##........##..#...................#....##..##..##...............#....#...........#......##..#.......#..
....................#.#...#.........##....................#.....#.....#..#....................................#.#.#........#.#.....
............##.#..#.#..#............#..#..#...........#.........#...#...............................#...............#...#........#.
.......##......#..............###......#............#.....#..........#..#.#.....#............#.#.......#..#...............#.....#..
.......###.......#......#..#..#...#...#...........#......#.........#...........#...................#.....##..........#....#........
..#...#.#..................##........................#..............##.....#.#.......................#...#..#...#..................
.........#.......#...#...#.....#......#..............##.......##....#.............##..............##.....#....................#....
..#....#.....#.....#............................#.#.......#.............#.....................#.........#.......#....##.#..........
...#.#...#........#....#........#.#...............#..##......#........#.........#...................#...#.......#...#......#...#...
.......................#..#..#.....#............#.#...#..#...........#.......##....#.#................#...................#....#...
.........#....#.#..........#.#....#..................#..............#...#...#..........#..........#.............##..#.......#..#...
...#......#.#......#.....#.##.#...#.....................#....##............#......#...#.........##...............#...............#.
........#.......#...................................#....#...............#....#.##..#............#...#........................#....
.......................#.......#..............#........#.#............#............................#...#......#..#..#.##.#.........
......#...#.##.............................##.#...................#..#......#.#...#...##..#.........##......#..#....#........#..#..
.........#.#............#................#...#.##......#..#.....#.##....##.#...........#....................#.............#........
.....#...#.........#...#.#...#.........#...#....#......#......##...#..#........#.....#.....#.............##.........#..........#...
.....#.......#............#.#............###..#..#.#.....#.##.......#.#.............##..##.....................#.#.......#..##.....
.#...#.....#...#.......................#...#.##...#....#........#.....#...#..................#............##..#...#....#.##.#...#..
......##.............#...#.........#.....#.....#..#........#......#..#......#..##...#..##.....#.#.................#..#.........#...
.........#..........#.#............###.................#..#.....#.....#..#...#..........#........#..........#.................#....
..#..........#....................#.....#..#...#....#.....#..#............#....#..#.........##..##..........#....#......#..........
.........#..........#..#.......#............#...........#.......#..#...........#.......##.....................#.......#............
...#...........................#..#.........#.#.....#.........#..........#........##.......#...#...............#....##.#.#....#....
....##.....#...#.......................#..#...#......#....#.#.....#..##.......#.........##.....................................#...
.#......#....##...#............#.....#..#........#......#...##.......#...#...#.#...#.#.....##..........................#..#........
.#.......##...................#..#......#..#..................##....#..#....#..##.......#..#.....###................#.#............
..............#.............#..#....#.....#.......#..#....##.........#..#.#.....###.......#.....##...##..........#..........#......
..#.#..#.................#......#...#.#...#......#...#.............#...#..##...............#........#............##.#.....#........
........#.....#.................#.....##.....#.#...#..#...........#..............#...#........#........##..............#...........
.#......#........................#..##.#........##......##....#...#........#..#...#..##.........#....#...............#.............
....#.....##..#............##.........#...#......#.#.....#.........#.#.#....###......................#...#.........................
.........###............##..........#.#.#.....#.....#..###.....................##.....#.....#....#.#...............................
...#.#.#............#..#.........###.....####.......#.....#...................#..##..##.......#....................................
.#.#.....#.#...............................#..#.......##........................#...#....#...##.#...#....................##........
...#.....................#...#...............#...#..#...#....#......##....#.#..#.#......#...#........#.....##.................##...
.........#.........#..#...#.#.#...#...#..........#......####.#.......#................#.........###...#..#.........................
....##..........#.........#..................#....#....#........#...#.........#..##...#....##.....#.#..........#..#..........#.....
...#............#...#..#.#.#.....#..##.#..#.......................#..#.....#....#...#.....#.....#..........#...#...#...............
.#..................................#..#..#..............#...#.....#...#................#.....#...#...........##..#................
....##................#....#.......#..#....#....#............##..........##...#.......#.......#...........#....#...................
....#..........#..........#..##......#.......###...........#............#..........#...#.#.#........#.#.#.#.#.#......#..........#..
............#...........................................#.......#.##.#.....................#..#..#..#.............#..............#.
...........#...#.........#.#....#................#.....#........#..#...#.#.....#..........#....................#...#..#............
...........#.#...#....#.....#..#..........#...#.#..#..#.#.........#.#.##........#......#................#...........#..............
................................................#..#..#....#.........#.............#....###..#...............#...#.......#.........
............................#...#..#......##......#...#........#....#.....#.....#...#..#...#..........................###.#........
..............#............#...#.#...##.......##.....#.##..........#....#............#.......#......##.#....#.......#..............
...........#.##.##..#.#.#.........###........#...#...#................#.#...............#..#.#........##........#.##.......#.......
.................................................................S.................................................................
.............#.#......#..##..#.............#...........#..........#..##.#...............#..###..............#...#.##.##....#.......
......#......##..#...............#..#............##.......#.#.......##..#...#.....##........#.................#....#......#........
..........#..........#.#........#....#......#........#.....#.............#....#..............#.#..#...#.............#..............
...........#........##.#.......#...#....#........#.#.....#........#..#..........#......##.....#...............#....#....#..........
...........#............#.....#..#.............#..#....#.#.....##............#.......#.......#............#.#..##..................
...........##..#...#.....##.........##.......#.......#.............#...#......#...##........#..#....#.......##..........#..........
.............#.......#......#..#...##..#.................................####...#......#.............#..#..#.......#...............
................#............#........#...##......##...................#..#...#.........#..#.........................#.............
.........................#...##.#......#...........#........#........#...............###....#.#........#............#..............
...#...........##........#..##.....#.....#.##......#....#..#........#.................#.#.........#........##..#..#................
..#.............#.##....#......#...#.#.......#.......#.................#.....#.....#.............###......#...#............##......
........................#.....#..##.........#.....#.....#..#........#...#.....#.....#....#......#...##..#...#.............#..#.....
......#........................#.....#..#...#....#......#....#..#.#.#...#.....#..#......#..#..#...#...#..#..#...............#......
..........#...........#......##.#.#............#...................##.....#..#...##....#..........#........#.............#.........
.#.......................#.....#..........#..##....#....#...............##......#.#..#.###..............................#..#..#....
.......................#........##......#.#......#......####..........#..#........#.#...........#..#.#..#.....#....................
....##.....................#..#.......###.................##.......#...#..#.###..#...#..###........#....#.#.............##....#....
........#..#............#...#.........#........#....#.##.......#...#.......#..#.#.......................................###........
.......##.....#..............#..#....#..#...............#......#................#................##......#.............#.#.........
....#......##...........###.....##................#...#....##.#......#....................#...##.#.......#............###...#......
.#...#.........#..........#..#.##.#.#...........#.#.........##....#.#......#........#.....#..#..........##.............#...#....##.
.........#.#.#....#..........#.....#...#..#.#...............#..........#....................##........................#.##....#..#.
........###..#..................#...#.....#.................#.#....#........#................#...#...........................#...#.
.....#.#..#........#.............#.........#.............#..###......#..#.#....#................##.....................##..........
.#......###....#....#........#.#..#...#.#..#.#...#........................#.........#...#.........#..........#.......#.........#...
..#.....#...........###........#.............#......#......#..........##.......#.#.#....#...#.................#...........#........
.....#........#..#..#................#......#....#.........#.##.......#...........#...........#.#............#.#....#......##......
.....#...##.#.....................#......................................#.#..#.#.....#.......#.#............####.#................
.......#...#......................####..#..#.............#..#.................#.....#...........#............#..#..........#.......
.##.............#...#.#...........................#..........##............#.......#.#......#............#........#................
.......#.......#.......#................#......#.#.#...#.............#........#........#.....#.#..............#....................
.#....#......#.#....#.................#.................##..............................#...............#.....###....#......#....#.
........##...............................#........#.............................#.....#..####.............#...#..#.........##.##...
.#.#............#..#..........#.......#......##.#.#..........#.........##.....#...........#..................#......##.......#.....
.................##.....#.#..............#.#.#.###...#.............#.#....#.....##...#.##..#...........#.#.#....#..................
..#.#..#..#................................#.#.......##.............#.#..........###................#.....##....#..................
........#........#.##..#.....##.#........#.#...#..#..#.....#...#....#.#.......#...#.....#..............#.........#.#...#...........
........#.....##..#..#......................#................#...............#...#.#.#..........#.#.##......................#......
........#......#.....#.#.......#............#...........##...............#............#...........#..........#.#..........#.#......
......###.....#.................###................#......#........#...#.#.#...#....#.#.........#...#...##........#...#............
..#.##..#..#....#.....##..........#............#................#.#..#.........................#.##..##..........##....#...#.....#.
....#............#.#..#...#......#..............#..#....##...#..#.....#.........#...#...........#..........#.#......#.#...##.###...
....#................#.........#...............##..#...........#.......#.#..#.......................#....#...##........#...........
....#....................##......#...#...........#.......#.#..#..............##.#....................#....#...................#....
...#..............##......#..............#...........#.......#......#........##.............#.........#...#.#..##..................
....#........#.#....#....#.#.....#........#..........................#..#.#.#.#.........#..........#.##...#.##..###.....#..........
.#.......#.#............#..#....#.......#.#............#..........#.#........##.............#...#..#....#..........#..........#....
...#...#......#..............#...##.......#..........#.....#......##...#..............#................#..............#...##..##...
....#......#.##..##....#.....................#..........#....##......#................#.......#.#..................#......#.#......
...#...............#..#..........#..#.....#..##............#.......###...........................##................#...#...........
........#....#........#.......#....#..#.#..................#.............##.............#..#.......##..........................#...
.#....................#......#.#........#..................#.##.#..#.....#.........##................................#.............
.......#....................#..#................#......................................##..........#.....#..##..#..........#.#..#..
..#.....#..#....#.....#......#..##...#.........#.........................................#...#....#.........#...##.......#.#.....#.
...........#..........#..#.#...#...##..#.........#.#...........#....#..............#.............#.......##..#...#....##.......#...
..#.........#.#..#...#.....#......................#..........#........#.......#....................#...##......#.................#.
.#......#...#.#.#......#.#........#.#......#....#....#............................#......####............#.......#..#..........##..
.#......##..........#...#...........#....#...#.#.................................#.##...#......................#..#..........#.#...
.....#...#.......#...................##.#....#...##............#..##.........#..................#..#......#..........#....#.....##.
.#..................#..........#.....................#.#.........................#...........#...##...#...#...##.#.................
.#..#..#..#..#......#........#........#.......#.#...#.#..................#.......#.......#..#.#......#...#...............#.....#...
.#........#......................#...##.......#..##......................###......#.....##....#..#................#................
.#...........#.............#..#.....##.........#.##........#.............#....#.#...........#...............#.........#........#...
......#....#.#...###................#....#......#..#........#.........................##.##...#..#..#.#......#.....................
...................................................................................................................................`
