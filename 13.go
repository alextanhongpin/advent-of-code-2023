package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	fmt.Println(solve(example, false)) // 405
	fmt.Println(solve(input, false))   // 35360
	fmt.Println(solve(example, true))  // 400
	fmt.Println(solve(input, true))    // 36775
}

func solve(in string, fix bool) int {
	var result int

	sections := strings.Split(in, "\n\n")
	for _, s := range sections {
		s = strings.TrimSpace(s)
		v := parse(s)
		if fix {
			result += fixSmudge(v)
		} else {
			n, ok := getReflect(v)
			if !ok {
				panic("hello")
			}

			result += n[0]
		}
	}
	return result
}

type Point struct {
	x, y int
}

type Range struct {
	min, max int
	dir      string
}

func fixSmudge(grid map[Point]bool) int {
	var maxX, maxY int
	for p := range grid {
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	// This is what we are supposed to get without smudge.
	o, ok := getReflect(grid)
	if !ok {
		panic("invalid")
	}

	// Go through each point, and update the smudge.
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			p := Point{x, y}
			newGrid := maps.Clone(grid)
			newGrid[p] = !newGrid[p]
			n, ok := getReflect(newGrid)
			if ok {
				for _, j := range n {
					if j != o[0] {
						return j
					}
				}
			}
		}
	}

	return o[0]
}

func getReflect(grid map[Point]bool) ([]int, bool) {
	var maxX, maxY int
	for p := range grid {
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}

	cacheByY := make(map[int]string)
	cacheByX := make(map[int]string)
	var rngs []Range

	// Find horizontal (y-min, y-max) reflection.
	for y := 0; y <= maxY; y++ {
		var xs []int
		for x := 0; x <= maxX; x++ {
			if grid[Point{x, y}] {
				xs = append(xs, x)
			}
		}
		cacheByY[y] = fmt.Sprint(xs)
		if cacheByY[y-1] == cacheByY[y] {
			rngs = append(rngs, Range{y - 1, y, "horizontal"})
		}
	}

	// Find vertical (x-min, x-max) reflection.
	for x := 0; x <= maxX; x++ {
		var ys []int
		for y := 0; y <= maxY; y++ {
			if grid[Point{x, y}] {
				ys = append(ys, y)
			}
		}
		cacheByX[x] = fmt.Sprint(ys)
		if cacheByX[x-1] == cacheByX[x] {
			rngs = append(rngs, Range{x - 1, x, "vertical"})
		}
	}
	result := []int{}

	for _, r := range rngs {
		if r.dir == "horizontal" {
			var valid = true
			dy := min(r.min, maxY-r.max)
			for y := 0; y <= dy; y++ {
				prev := cacheByY[r.min-y]
				curr := cacheByY[r.max+y]
				if prev != curr {
					valid = false
					break
				}
			}
			if valid {
				result = append(result, r.max*100)
				//return r.max * 100, true
			}
			continue
		}

		var valid = true
		dx := min(r.min, maxX-r.max)
		for x := 0; x <= dx; x++ {
			prev := cacheByX[r.min-x]
			curr := cacheByX[r.max+x]
			if prev != curr {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, r.max)
			//return r.max, true
		}
	}

	if len(result) > 0 {
		return result, true
	}

	return nil, false
}

func parse(in string) map[Point]bool {
	lines := strings.Split(in, "\n")
	res := make(map[Point]bool)
	for y, line := range lines {
		for x, c := range line {
			res[Point{x, y}] = c == '#'
		}
	}
	return res
}

var example = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

var input = `#.##....##.#.
#.##....##.#.
.#.#....#.#.#
.###....###..
#.#.####.#.#.
.####...###.#
#..######..#.
..#..##..#..#
.#........#.#

###......####
..##....##...
.#.######.#..
#.###..###.##
#.##.....#.##
.#.##..##.#..
.##......##..

...#.#.##.#.#....
..###..##..###...
#..##.#..#.##..##
#.###.####.###.#.
####..#..#..####.
####.##..##.#####
#..#.##..##.#..#.
..####.#######..#
......####......#
...#........#...#
...#........#...#

#..###..###
#..#.#..#.#
..###.##.##
....#....#.
.....####..
###.#.##.#.
#.#####.###
##.###..###
#...#....#.
#..#.####.#
#..#.####.#

##..######..#
###.##..##.##
###..####..##
...#.#..#.#..
..##.#..#.##.
###.#.##.#.##
##....##....#
####.####.###
...########..
..#.#....#.#.
###..####..##
..##......##.
###.#.##.#.##
####..##..###
#..#......#..
...##.##.##..
###.##..##.##

####....#...#
###.....#...#
.###.#.###.##
##.#...##.#.#
#.##.#.#.###.
#.#.#..##..#.
##..##..#...#
###...###.#..
###...###.#..
##..##..#...#
#.#.#..##..#.

.####..####..####
.#.######.#..#.##
##...##...####...
#.#.#..#.#.##.#.#
..#.####.#....#.#
#.#.####.#.##.#.#
.##..##..##..##..
##..#..#..####..#
...#.##.#......#.
##.#....#.####.#.
..###..###....###
.##########..####
..###..###....###
#.#.#..#.#.##.#.#
#....##....##....
.#...###..#..#..#
#..#....#..##..#.

.##..#..#........
.##..##.......#.#
#########.###....
..#..#..#....#.##
..#..#........#.#
#......##..###...
.######.#....####
#......###.######
..#..#..#..##....
..#..#..#..##....
#......###.######
.######.#....####
#......##..###...
..#..#........#.#
..#..#..#....#.##
#########.###....
.##..##.......#.#

..####.
.##..##
##.##.#
##.##.#
.##..##
..####.
.######
#......
###...#

##..##.
##...#.
##.#..#
##...#.
..#..#.
.#..#.#
.#...#.
.###.#.
.###.#.

#..####..######..
..##..##..####..#
..#....#...##...#
..#.##.#..#..#..#
.#.####.#..##..#.
.##....##..##..##
..........#..#...
..#....#.#.##.#.#
.##....##.####.##
.########......##
##.#..#.##....##.

.##.#..#.##.#....
#..#......###.#..
#..###.###.##....
.##.#...#....#...
.##...###...###..
#..####....##.#..
.##.####.########
#####.####.##..##
#..####.#####.###
###..###.#.##.###
.##....#.#.#...##
.......#.#.....##
#..#...#..##.####
.##.#..##.#.##.##
#..#.#.##.#.#....
#######.##..#.###
#####.##....#.###

..#..#..#..
#......#..#
.######.#..
#..##..#.#.
#..##..#.#.
.######.#..
#......#..#
..#..#..#..
..#..#....#
##....#####
##.##.#.#.#
#.####.#..#
..#..#..##.
..####.....
...##...###

.####.#.##.#.
#....###..###
#.##.#.#..#.#
......#....#.
.#..#.#......
.####.#.##.#.
......##..##.
.......#..#..
.#..#..####..

...##.#.##..#
#.#..######.#
#...##..###..
##...#.....##
..#...#.##.##
..#...#.##.##
##...#.....##
#...##..###..
#.#..######.#
...##.#.##..#
#..#..#.#....
#..#..#.#....
...##.####..#

#.#...#.##.#..#
##..##.####.##.
.#######..#####
#.#.##..##..##.
.#.##.##..##.##
#.#.#..#..#..#.
.....#.#..#.#..
..#..##.##.##..
..#..##.##.##..

.#..##.#..#..
#..#.####.##.
###.#.##.#.##
###.#.##.#.##
#..#.####.##.
.#..##.#..#..
.#...#.#.###.
##...##..##.#
.#...#.....##
.###..#..##.#
.###..#..##.#
.#...#.....##
##...##..##.#
.#...#.#.###.
.#..##.#..#..
#..#.####.##.
###.#.##...##

##..#..##.##.###.
#....#.#.##...#.#
#....#.#.##...#.#
##..#..##.##.###.
.##....##....#..#
.##..#..#.###....
..###...######.##
#.##.###..#..####
....#.#.###.###..
.##....#####.###.
..##.##..##.#..##
..##.##....#.....
..##.##....##....

###.###..###.
###.###..###.
.#...#.##.#..
.#.#..####..#
.#...##..###.
#.###########
#####......##

##..#.##.
##..#.##.
###.#####
.#.#..##.
####..##.
#....####
.########
..#.##..#
.#.####.#
#.###....
.##.#.##.
###.#####
#####....

#.#..##
...#.##
..#..#.
...#.##
###.##.
#.....#
##....#
#####..
#####..
##....#
#.....#

.##....##..
..#..#.##.#
..#.##.##.#
.##....##..
.....#....#
#..#...##..
..#..##..##
.###..#..#.
##.##......
#.##..#..#.
...#.......
.##........
.##..######

#.####.#.##.####.
#......##.#.#..#.
...##.....##....#
#.####.###.#.##.#
#.#..#.#.##..##..
.##..##.###.###..
...##.....#..##..
.#.##.#.##.......
#.####.#..#......

..##..#..#.#...
..#####....####
...##..##.#..##
...##..##.#..##
..#####....####
..##..#..#.#...
#..#....##.....
#.##.##.##...##
..##.......#.##
.#.##.#...#####
##....###.##..#
#.####.###...##
.#...##.##.....
###.#.#.#.##...
##....#....####
....#.##..##.##
#...######.####

#..#.##.#..#..#
#..#.##.#..#.##
..########...#.
.####..####....
##.##..##.##.#.
#..........###.
#####..######..
....####......#
#.#..##..#.##..
.#.#.##.#.#.#.#
..###..###..##.
#.###..###.###.
###.####.####..
##.#.##.#.####.
####.##.#####..

#.#..#.#.#.#..###
#.####.###.#.##.#
..####...##..#..#
##....##.#.###.##
..............#..
.#....#...#..###.
#.#..#.#######...
#.#..#.#.....#...
.#....#...######.
########.....###.
.##..##.#.#.....#
##....##..#..##.#
.#.##.#...#..##.#
.#....#..##...#..
#.#..#.#.#.#.#..#
#.#..#.##.##.###.
#.#..#.##.##..##.

#.#....##....#.##
#.#####..#####.##
#.....####.....##
.#.#.#....#.#.#..
...#.#.##.#.#....
.##.#......#.##..
#.#...#..#...#.##
###.########.####
.#.#..#..#..#.###
.#.###.##.###.#..
###..#.##.#..####
#.#.##.##.##.#.##
#.#..........#.##
.###.#....#.###..
####.#.##.#.#####
#.#.#.####.#.#.##
.#.###.##.###.#..

......#..
#.#.#..##
..#..#..#
..#..#..#
#.#.#..##
......#..
##.....#.
#.#.#.#..
##.#.#.##
....###.#
#..###.##
#..###.##
...####.#
##.#.#.##
#.#.#.#..

..##..##..#.#
###.##.#..#..
##.######.###
##....#.#.##.
##....#.#.##.
##.######.###
###.##.#..#..
..##..##..#.#
#...#...##...
#...#...##...
..##.###..#.#

#.###.##..#..##.#
##.##....#....##.
##.##....##...##.
#.###.##..#..##.#
#.#..#.#.#..#.###
#...#.###..#...##
.####..#.#.#.....
..#####...#####.#
....####..#....#.
.#.###..##.###.##
.#.###..##.###.##
....####..#....#.
..#####...#####.#
.####..#.#.#.....
#...#.###..#...##

..##.##...#
##########.
.#.##..#..#
..##..##.##
......#.#.#
......####.
###.#####..
###.#####..
......####.

##...#.#..#.#..
##.#.#..#.##.##
#.####.##..#.##
#.#######..#.##
##.#.#..#.##.##
##...#.#..#.#..
.#.#.###...##..
...##..###.##..
.#.######...###
..##.#..#.##.##
###..#...##.###
######.#.#.####
#..#..####.####

#......###.#..#.#
.##..##.##..##..#
########....##...
.........#..##..#
###..####..#..#..
##.##.##..##..##.
..#..#..##.#..#.#
###..####...##...
.#....###.##..##.
#......#.#......#
##.##.###........
...##.....#....#.
#.####.###.####.#

###..#.#.#...
#.#...#.#..##
..##......#..
..#.##..##..#
.##.##.##.#..
.#.###.###..#
.#.###.###..#
.##.#####.#..
..#.##..##..#
..##......#..
#.#...#.#..##
###..#.#.#...
#####...#....
.###.####.##.
.###.####.##.

#..##.#..#.
...##..##..
...##..##..
#...#.#..#.
.##.#.#..#.
.##.#.####.
#.#..#.##.#
#...#.####.
..#########

.....#.####.#..
.##.###.#..###.
....#.######.#.
....#..#..#..#.
.##.#........#.
#..#..##..##..#
####..........#
####.########.#
.##.####..####.
#..#..##..##..#
#..###.#..#.###
#..###.#..#.###
.##..#.####.#..
#..###.#..#.###
#..#.########.#

....#....#..#
.#.##.##.#.#.
###.###.#...#
###.###.#...#
.#.##.##.#.#.
....#....#..#
..###..###.#.
..#.#.###..#.
####..###.##.
####..###.##.
..#.#.###..#.
..###..#.#.#.
....#....#..#
.#.##.##.#.#.
###.###.#...#

..######.#..#.#
.###..#.#..#.##
.###..#....#.##
##.###..#.#..##
......###.#.###
##..##....##...
##.###.##.#.##.
####..##..#.#..
####..##..#.#..
##.###.##.#.##.
##..##....##...
......###.#.###
##.###..#.#..##
.###..#....#.##
.###..#.#..#.##

..#.#...##..#
.##.#...##..#
..#.#######.#
##.#.####..##
.#.#.#...##..
##....#..##.#
.#......##.##
..##.##..##.#
.....#####.##
###.....####.
##.#..#....##
##.#..#....##
###.....####.

...#.##....##.#
###.#.#.##...#.
###.#.#.##...#.
...#.##....##.#
...#.###.#.###.
#..##..##.....#
##.#....##.###.
###..#######.##
##..#..#....##.
##..##.###..###
..#..#.....#...
.#####.#..###..
####.#......#.#
####.#......#.#
.#####.#..###..
.##..#.....#...
##..##.###..###

##.#..##.#.##.#.#
#.##.#.####..####
.##.......####...
.#########....###
.#....#####..####
.#....#..#.##.#..
.##...#..##..##..
##.#.....#....#..
##.###...######..

#.##.#..###
##..###..##
.#..#.#.###
.......####
.####.#####
.####....##
#.##.##.##.
......#.#..
..##..##.##

#.......#
#.......#
##....###
..#.#..##
#..####..
...##..#.
###.#.##.
..###..#.
.....#.##
###.#.#..
.##.##...
#.##.##..
#.##.##..
.##.##...
###.#.#..
.....#.##
...##..#.

.#.#.###.#..#
.##....######
.##.#.#.#####
.##.#.#.#####
.##....######
.#.#.###.#..#
#.#.#.####..#
#.#...#...##.
#..#.###.....
..#.#.#.##..#
.######..###.
#...#.#.#####
.#..#.##.....
..###.#.#####
#.#######....

###...###.#####.#
#....#.##..##.#..
.#.#.#.#.#.......
###...#.....#####
.##...#.....#####
###.#.....#...##.
###.#.....#...##.
.##...#.....#####
###...#.....#####

..#..#..###
##.##.###..
#.####.####
#########..
###..######
..####.....
#..##..#.##
..#..#.....
..#..#.....
.#....#.#..
...##..####
########...
########.##

.#..#.##.####.#
######.########
.#..#..........
.#..#....####.#
.#..#..###..###
......#...##...
.#..#.#########
.#..#...######.
######.........

..##...#.
####.###.
#.##.#.#.
....#....
....#....
#.#..#.#.
####.###.
..##...#.
#.##....#
...##.#..
.###.....
.###.....
...##.#..
#.##....#
..##...#.

##....###.#.#.#
#.####.#.#..#..
.#....#..###..#
..####..###...#
..####...##.#..
##....###.#...#
.#.##.#.#.##.##
.#.##.#.#..####
.........#.#..#
#.#..#.#..#.#.#
.######...##.##
.######...##.##
#.#..#.#..#.###

#..#....#..#....#
#..#.#..#####.#.#
.##...###..##.###
####.#.#.#.#..###
####.#.#.#..#..#.
#..##..#.###.#...
#..#.#.###..##...
.##.....###.....#
.##...#..##.#.#..
.##...#..##.#.#..
.##.....###....##

#.##.##...#.###
####.##.#.##.#.
...######.#..##
..#..#.....##..
.####.#.####.#.
#.####....##.#.
.####.#.##.##.#
#.#...#.#.#....
#.#...#.#.#....
.######.##.##.#
#.####....##.#.
#.####....##.#.
.######.##.##.#
#.#...#.#.#....
#.#...#.#.#....
.####.#.##.##.#
#.####....##.#.

#...#..##.#.##..#
#..#.#...###.#..#
#..#.#.#.###.#..#
#...#..##.#.##..#
#...#.......##..#
.#..#.#.#.###.##.
..###....#.......

.####...##..#
.####....#..#
####.#...####
...#.#......#
#.....##.#.##
#...#..#.##..
#...#..#.##..

.#..########..#
#.###.#..#.###.
...#.#.##.#.#..
#.############.
.#####....#####
.......##......
#....#....#....
...##.#..#.##..
#.####....####.
##..##.##.##..#
##..##.##.##..#
..####....####.
...##.#..#.##..
#....#....#....
.......##......
.#####....#####
#.############.

......##.
.#..##..#
..##.##..
###.##..#
#....####
#.##.##.#
#.##.##.#
#....####
###..#..#
###..#..#
#....####
#.##.##.#
#.##.##.#
#....####
###.##..#

.#.###.###..###
....#.#..#..#..
...#..#.#....#.
#....####....##
.#....##......#
##.#.#.##....##
#.#..##########
#.##.#.#.####.#
#.##.#.#.####.#
#.#..##########
.#.#.#.##....##
.#....##......#
#....####....##

###.####.
##...#.#.
..###.###
##..###..
##..###..
..###.###
##...#.##

.#....#.#..##
#.####.####.#
###..####.###
###..####.###
#.####.####.#
.#....#.#..##
#.####.##....
.........###.
##.##.##....#
#.#..#.#...##
.#.##....#.##
###..###..#.#
.#.##.#..#..#
.#....#.#....
#......##..##
..####...#.##
.#.##.#.##..#

.###.###.#.
#.#..###..#
...###.#..#
...##.###..
...##.##...
...###.#..#
#.#..###..#
.###.###.#.
###..##....
###..##....
.###.###.#.
#.#..###..#
...###.#..#

.#.....
###.#..
#.....#
##.###.
.#....#
.#....#
##.###.
#.....#
.##.#..
.##.#..
#.....#
##.###.
.#....#
.#....#
##.###.
#.....#
###.#..

##...#..#...#
..##..##..##.
...#.#..#.#..
##.#.#..#.#.#
##..#####...#
###..####..##
..#...##...#.
....#....#...
....#....#...
###...##...##
###........##

#...##...##..##
.#......#.#....
..#....#...##..
####..######.##
###....###..#..
#.##..##.######
#........###...
##......##.##..
.########....##
..........###..
..#...##.....##
..##..##..#.###
..#....#....###

..#.####.#..#
..#.####.#..#
...#....#....
##........##.
#.#.####.#.#.
#.#.#..#.#.#.
#.##....##.#.
#.#.####.#.#.
..########..#
.####.#####.#
.##.#..#.##..

#.#....#.####
..####.##....
..#...#######
..#....######
..####.##....
#.#....#.####
#...###.#....
#.##.##..####
#..#..#......
.#..##...#..#
######.#..##.
##..##..#....
#..###.#.....
...###.#.####
#.#..#.##....
##..##.##.##.
.##.#.#..####

.###...
...#...
...#...
.###...
...##..
##..#..
#.#.###
....#..
##.#...
....#..
...#.#.
.##.#..
##..#..
###.###
####...

######...#..##.
##..#######..#.
.######.#####.#
#.##.###..##..#
##..##.###.#..#
.#..#.###...###
#....##.....#..
.........##.###
..##..###..#...
........#####..
..##..#...#.##.
......#.####...
..........###.#
.......##.#####
##..##.########
##..##..###...#
##..##..###...#

.##..#.#..#.#..
.##.##..##..##.
....###....###.
.##.#.#....#.#.
.#..#..####..#.
#####.##..##.##
....#.#.##.#.#.
.##...#.##.#...
#..#..#....#..#
#####.##..##.##
.......#..#....

#.####...
###.###..
....#..##
##.......
##..#....
##..#.#..
###.#...#
.##....##
..##.....
#........
#..##....
#..##....
#........

#.##..###...##.
...##..#.###..#
#.#.##...#..##.
.##...##....##.
#..#....#..####
#..#....#..####
.##...##....##.
#.#.##...#..##.
.#.##..#.###..#
#.##..###...##.
##.....###..##.

.##.###..###.
.##.###..###.
...##.#...#..
.##...##..###
######...#.#.
.##..##.#.###
.##.##..#....
####.#.#.#...
.....#...##..
.##.####..#..
.##.....#...#
.......#.###.
#####..#..##.
########.##.#
.##....#..#.#
.......##.#..
#..#.###..###

###.#.###.#
###.###..#.
#..###.##.#
#....##.#..
#....##.#..
#..######.#
###.###..#.
###.#.###.#
#..########
##.#.######
....#..#.##
.#.##...###
######.#.#.
..##.###..#
..##.###..#
######.#.#.
.#.##...###

......#.#
.##.#....
....#..##
....#.#..
#..###...
#..###...
....###..
....#..##
.##.#....
......#.#
#..##..##

...#.##.#..#.##.#
.#...##..##..##..
###.####.##.####.
..###..##..##..##
..##.....##.....#
..#.#..#....#..#.
.###.##.####.##.#
.##......##......
.################
..#..............
.#...##......##..
##..####.##.####.
#..##..######..##
..##....####....#
#................

#...#.#######
#..###.......
###.##..#..#.
.#..#########
..##.##.####.
#...##.......
.#.#..##.##.#
#######.#..#.
#.#......##.#
#.####.#....#
....###.#..#.
##.#.###....#
..#....######
.#..#..#.##.#
.#..#..#.##.#

###...#..
#.####.#.
#.####.#.
###...#..
#..#..#.#
.##.##...
.##..#...

#..#....#..##..
#.#.####.#.##.#
.####..####..##
..##..#.##....#
..########....#
..########....#
#.#..##..#.##.#
.##......##..##
##........####.

####.#.#.###.
##.####....##
...#.#...##..
...#.#...##..
##.####....##
####.#.#.###.
..#..###.####
..#.#....#.##
##.#.#.##..#.
...#...#....#
.##..#.#...##
....###..#.##
##...#.#.####

###....####
.#.#..#.#..
##.#..#.###
.##.##.##..
###.##.####
..##..##...
#...##...##
...#..#....
...####....
##......###
.##.##.##..
#.##..#####
..#....#...

#.#.#.##..##.
#.#.#.##..##.
#####.#.##.#.
#..#...#..###
##.....#.##..
##.##.#..#.#.
.#.#...###.##
....####....#
#.....###....
#....#.#...#.
#....#..##.#.
.###.#.###...
##.###...####
##.###...####
.###.#.###...
#....#..##.#.
#......#...#.

#.##.#.###.
.####.#....
.####.#....
#.##.#.#.#.
######..#..
#.##.##...#
#....#..#..
.#..#..##..
..##..##.##
......####.
.#..#.#..#.

##.#..##..#.###
.#..........#..
.#...#..#...#..
...##....##.#..
###........####
#..##....##..##
##...#..#...###

##...#.##.##.##
##.#..##..##..#
#.#..#.#......#
#.#...##....#.#
.#...###.#..#.#
.##.####..##..#
###..##.#....#.
#.#.#.....##...
.....#....##...
....#.###....##
....#.###....##

###.#.###..
....#.###..
....###....
##.##.#.###
##..####.##
####.##.###
##..###....
###..#...##
......###..
.....###...
..##..#.#..
##...###.##
...#......#
##.####.###
..##..###..
..##.##.###
........#..

#.###.....#####
#.###.....#####
....#..#.##.#..
##..#.##.#.....
#..#...#....##.
#...####.#.#.##
.#..#.#.###..#.
##.#.#.#....##.
.###...####.#.#
.###...####.#.#
##.#.#.#....##.
.#..#.#.###..#.
#...####.#.#.##
#..#...#....###
##..#.##.#.....

#..##..###.
.######.##.
########.##
.##..##..##
#########..
#..##..###.
#..##..#...
#..##..##..
.##..##.#.#
.##..##.#..
........###
#..##..#..#
.##..##.#.#

....#..###.##.#
.#..#..###....#
.#..#..###....#
....#..###.##.#
#..###.##......
#..##......###.
.#.##.#.##....#
.##.#.#####..##
.#.#.#..###..##
.#..##.###.##.#
.#...#.##......

..#.###
###..##
#.#.#..
####...
.#.#.##
#.##.##
###.#..
.#.####
##...##
##...##
.##....
.##....
##...##
##...##
.#.#.##
###.#..
#.##.##

.....##..
..#######
##.######
##...##..
..#######
.#.......
#..#....#

...#..#....
##.####.##.
#...##...##
.#......#..
##..##...#.
...####...#
..#....#...
####..#####
###.##.###.
#...##...##
##......##.
##......##.
#...##...##
###.##.###.
####..#####
..#....#...
...####...#

..##.####
.#.#.#...
.#.#.##..
####.#.#.
##.####..
.##.##..#
.#...#...
.####.##.
.##.#.##.
.##.#.##.
.####.##.

###.###
...#..#
##...##
..#..##
...####
##..##.
##..##.
...####
..#..#.
##...##
...#..#
###.###
..#....
##.....
###..#.

....#...#.#..
.##..#..##..#
....#.#......
######..#.###
.##.##..#..#.
####.##..####
.##..##..###.
#..#..###....
####.....##..
#..###..#..##
#..###.....##
####.....##..
#..#..###....
.##..##..###.
####.##..####
.##.##..#..#.
######..#.###

..#.###.###..
###...##...#.
....#.##.....
....#.##.....
###...##...#.
..#.###.###..
...######.##.
#...###.#...#
...#.###..#..
##.###.##.##.
..####.#.....
...#.####.#..
####.#..##...
....#..#.....
..#.#.#.#..#.

#..#.##..###.#.
.#.####.#.#..##
..#.##.#..##..#
..##..##..#....
#...##...##.#..
.########.....#
###....#####...
##..##..##.....
.#.#..#.#.#..##
##......##.####
...#..#...#....
.#.####.#..###.
.#.####.#..###.

#.#.#....#..#..##
...##..##.....##.
....#..##.....##.
#.#.#....#..#..##
.##..#.#..#.....#
..#.#.#..###...##
####..#.###....##
.##.###.###.#.##.
..##.####......##
#..###.##...#..#.
###..#.#####.###.
#..#.##...#####..
.#.#...#.#.##...#
.#.#...#.#.##...#
#..#.##...#####..

.....#..#..
.#..#.##.#.
###........
###..#..#..
.###......#
.#.###..###
..#.######.
..##......#
.#.........
..#########
##.#.#..#.#
##.#......#
.#.########
....#.##.#.
..#.######.
#...######.
##..######.

#.##.#.#.#..#####
...#.##.#..#.....
#..#..####.......
#.###.#..#...####
.###....#.#######
#..##.##.#.......
.###...###.##....
#...##.##..######
#....######..#..#
###.##.###.######
#.########..#....
.##..##....######
...#.###..##.....
.#.#####.#.......
..#..##..#...####

..##..#.#..
#.##.##..##
.#..#.#.###
..##....#..
#.##.######
#....###.##
##..##.#.##
..##..#....
######.#...
.####.#..##
.####...#..
..##..##.##
.####.##...
#######.###
.####.#.#..
..##..#..##
.#..###..##

..##.#...##..
#.####..#..#.
...#..#####..
..##.##..#...
..##.##......
##....#.#.##.
.##.#..#.#.##
##..######.#.
#...#..#.#.##
#...#.#...#.#
#...#.#...#.#
#...#..#.#.##
##..######.#.
.##.#..#.#.##
##....#.#.##.
..##.##......
..##.##..#...

#.##.#.##.#.#
.#..#......#.
.####.####.##
.#..#..##..#.
##..########.
.#..#.####.#.
.####.#..#.##
######.##.###
##..##....##.
.........#...
.#..#..##..#.`
