// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solve(example).steps / 2)  // 4
	fmt.Println(solve(example2).steps / 2) // 8
	fmt.Println(solve(input).steps / 2)    // 6697
	fmt.Println(part2(example3))           // 4
	fmt.Println(part2(input))              // 423
}

func part2(in string) int {
	s := solve(in)
	s.edges = append(s.edges[len(s.edges)-1:], s.edges...)
	// Pick's theorem.
	return shoelace(s.edges) - s.steps/2 + 1
}

func shoelace(p []Vec2D) int {
	// Shoelace formula
	// https://en.wikipedia.org/wiki/Shoelace_formula
	var sum int
	for i := 0; i < len(p)-1; i++ {
		sum += p[i].x*p[i+1].y - p[i].y*p[i+1].x
	}
	if sum < 0 {
		sum = -sum
	}

	return sum / 2
}

func solve(in string) state {
	m := NewMap(in)

	start := m.Start()
	q := []state{
		{start, Vec2D{0, 1}, 0, nil},
		{start, Vec2D{0, -1}, 0, nil},
		{start, Vec2D{1, 0}, 0, nil},
		{start, Vec2D{-1, 0}, 0, nil},
	}

	for len(q) > 0 {
		var s state
		s, q = q[0], q[1:]

		pos, dir, ok := move(m, s.pos, s.dir)
		if !ok {
			continue
		}

		// When the direction change, there's an edge.
		if isEdge := dir != s.dir; isEdge {
			s.edges = append(s.edges, pos)
		}
		s.pos = pos
		s.dir = dir
		s.steps++

		if m[pos] == 'S' {
			s.edges = append(s.edges, pos)
			return s
		}
		q = append(q, s)
	}

	return state{}
}

type state struct {
	pos   Vec2D
	dir   Vec2D
	steps int
	edges []Vec2D
}

func move(m Map, pos, dir Vec2D) (newPos, newDir Vec2D, ok bool) {
	newPos = Vec2D{
		x: pos.x + dir.x,
		y: pos.y + dir.y,
	}

	r := m[newPos]
	switch r {
	case '/':
		if dir.x != 0 {
			newDir = Vec2D{0, -dir.x}
			return newPos, newDir, true
		}
		if dir.y != 0 {
			newDir = Vec2D{-dir.y, 0}
			return newPos, newDir, true
		}
	case '\\':
		if dir.x != 0 {
			newDir = Vec2D{0, dir.x}
			return newPos, newDir, true
		}
		if dir.y != 0 {
			newDir = Vec2D{dir.y, 0}
			return newPos, newDir, true
		}

	case '|':
		if dir.y != 0 {
			return newPos, dir, true
		}
	case '-':
		if dir.x != 0 {
			return newPos, dir, true
		}
	case 'S':
		return newPos, dir, true
	default:
	}

	return pos, dir, false
}

type Vec2D struct {
	x, y int
}

type Map map[Vec2D]rune

func (m Map) Start() Vec2D {
	for p, c := range m {
		if c == 'S' {
			return p
		}
	}

	panic("no start found")
}

func NewMap(in string) Map {
	m := make(Map)

	// We replace the complicated symbols with simpler ones.
	mappings := make(map[rune]rune)
	mappings['L'] = '\\'
	mappings['J'] = '/'
	mappings['7'] = '\\'
	mappings['F'] = '/'

	for y, line := range strings.Split(in, "\n") {
		for x, c := range line {
			p := Vec2D{x, y}
			if v, ok := mappings[c]; ok {
				m[p] = v
			} else {
				m[p] = c
			}
		}
	}

	return m
}

var example = `.....
.S-7.
.|.|.
.L-J.
.....`
var example2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`
var example3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`
var input = `F7---|FJ-L-|7-7F-7.F7|-|F-L7F7--7F---FJ7-F-F.FF.J---|77--7FF-7.-7-|.7FF-77F|77.|7-FF7FJFJ7-7FF77L77F7-F.|FFJ.-|-77.JFF-|J-|F|77F77.F7-F7-FF7
LF7J.L7J7LFJF.FL-|-|LL.L7LL7-|7JFL|FJ.|.7.LL7J--JJ.LL-J7|JF|FJ7.F7---7J-|F|J77..7FJJ||F-7|.LFJ-F7--F7-L.LFJF77|LJF7--7.L7LF7L-7L777JFJ|FFL7-
|J|.FL7F7.L-7-77.F.|..LFL.LJ-J-77.L-|-..F7LLJF|.F.FF--7FFF-JL7FF7J.|L-77L7L|.FFLLJ7.JFL-F77..FLJ||.LF7J7F.F7.F--FL77LJ7F|.J7L|L7LLF-7-F7FJ|L
|-JFJ--777.F|FLL-|-7-F-|JL.L7|7|L-J||.LFL7...FL7||LJ.F7|||F-7L-J|..77-JJ.|7JFL7J7FJ7|FJ7L77FF7JFJF-77J77L.L|LJJ-L7|L|J|F-7-J-FL--|LJ|-7L|LF.
L7.LF|JLJ---|-JL-J.J-F.F-JF|FJJL-J.F7J7|.J-F|7FJ-F-J7F77LLJ.|F--JJFLL7JL-JJ-LJ|FFL|-||.7J|JFJL7.LL7L-7F|-L.LJ..LFF7---JJJF7-7FFJFLJL|L-F--JL
|F-JLL--J|L|JL7L||F..L|J.|FL77L7|7-||7FJ.FLLJ|7|L|L|-7.7F|FFJL-7FL7|||F|L|J||FL-J7|-|LF77F7|F7|77FJF-J7J7-LJ|F---FJLLF.|.|JFJ7|JFJ-FJ..|L||.
FJ7|.FJ-||-7F||7LFL.L-JJ7F-J||LLJJ.L|7||F-.J-|FJ-|J|-F-77F7L7F7L7L|FFLF77L--F.LLJ-|7JFLJF||LJ||77L7|F777F-J-J|.J7L..L|F-|.F-LLL77J.L777|7LF-
|.F|7JJ7L|.||F-J-J.-J.|FLFJFLJ-|LF-.||77JFJJL-J7..F7-L7L7|L-J|L-JF-77J.F-7LJ|777..L--LLFFJL7FJL---J||L7F||L7.|7.JJ-|LJ-.J.|-|.|LJ.-FFLLL---.
F-FFF.J|F7--|7|FL7F-J7FFLJ|.|FFLFJLLF|J..|-FF|L-|-LF--JFJL--7L--7L7|FF7|FJ.F77F|.FL|7LFFJF7|L--7F--J|FJ-F7JFJL77F|-L7|J-FFJJ.LJ7...JL|LFL-F|
.FJ7L7.F|J|.|7J7JLJ7JF7|-FJ.7-F7L.FF|F777|.L7FF7.-.L-7FJF7F7|F--JL||FJ|||J|-L7|..7FL7--L-J|L--7||7F7|L7|||.|JF7F-7.L|F-FJJL-L7-|-F..-7-|JFLJ
7--.FL7-J7.F.--F7F-|7F7|FJFFJ-L|JF-F|J||FJ-LF7J|7.FF7|L7|LJ||L7F7J|||FJ|L-7LLJ-|.|7FFF7F-7L--7||L7|||FJFJ|F7FJ||FJ7.FJ-|F|..L--7F--F|JF7-F-7
LLJF-JL7|-7JF|F-J-7L-L7LJJ-F7J|.||FF7|L-7|-LJF77-F7|||FJL-7||FJ||FJLJL-JF7L7J.||FJLF7F7L7L-7FJ|L7||||L7|FJ||L7LJ|JF-J|FJ7F7-LL-JJ.F-JF|J7F-L
L-7L|.LJJ7LJJL7J-|LJF7J-||||F---7F-JL7.F7F7.FJ|JJ|LJ|||F--J|||FJ|L-7F7F-JL-JFF7F7-7||||-L-7|L7L7||||L7LJL-J|FJF-J7--F||FJJJ.7-|7LF|L-J.LL|F|
F7J-77.F-|JJ---77-.FF7LF7F|-L-7FJL7F-JFJLJL7|FJ|FL-7|||L--7LJLJFJF7||LJF7.LF-J|||-FJ||L---JL-JFJ|LJL7|F----JL7|LF7-F||F|77.---.-7FL7LL7|-L-7
F-J.L7-7LJF7FLL|JJFFJ|F||J7FF-JL7FJL-7L-7F-J||LF7-FJLJL---JF---JFJLJL7FJL7JL-7LJL7|FJ|F------7L7|F-7||L-7F7.FJL-JL--7---F7FJLJ7LJ.FJJ.J777JJ
||--.LLLF.--7-|J|JFL7|FJ|7FFJF7FJL--7L--J|F-JL7||FJF7F--7F7L--7FJF-7FJ|F-JF77L7F7LJL7LJF7-F7FL7|LJFJ||F-J||FJF7F7F7FJ.L-L-7J7LJ7F-7J7F.LFJ..
7J.|-FFLLFJF-FF--7F-J||FJF7L7|LJF7F7|F---JL7F7LJ|L-JLJF7LJL7F-JL7|FLJ7|L--J|F7LJL7F7L7FJL7||F7||F7L7LJL--JLJFJLJLJLJ..|F||L7F-7|J7L7-77.J.FL
.LL|-F7FLJ-7J|L-7|L-7||L7|L-JL7FJLJ|||F-7F7||L--JF7F--JL7F-JL---JL---7L7F--J|L7F7||L-JL-7|||||||||L|F7F----7|F-7FF7J7FFJ.F--.LJLJ7L7|L-7.L7|
-.|.FJLFF.FJLLLFJL-7|||FJL-7F7|L--7||||FJ|||L7F7F||L---7|L-7F---7F---JJ||F7-L7|||||F7.F-JLJLJLJ||L7||LJF--7LJ|FJFJL---7J-|-F-F|FFJLL|-FF-7J7
.FJ7.FFL7--..LFL--7LJ||L-7-LJ|L7F-JLJ|||FJLJFJ||FJL7F--J|F7||FF7LJF7F--J|||F-J||LJ||L7L-------7|L7||L-7|F7L-7|L-JF----J.|.|L-J7F7JL7JJL|FJ.|
.7.F7F|7|7J-FLJF-7L-7||F-JF7FJFJL---7|||L7F7L7||L7FJL--7LJLJL7|L-7||L-7FJ||L-7||F-JL7L-7-F7F7FJL-J||F-J||L7FJ|F--JF7F-77777J.JJFLFFJ-FF-7.F7
7LFLL-JJJ||F|-FL7L--J|||F7||L7L-7F7FJ||L7LJL7||L7|L7F7-L7F-7FJL-7LJ|F7|L7||F7||||F7FJF7|FJLJ|L---7LJL--JL7|L7|L---J||FJF7J.F|.FF-|L7-|JFJF-J
JLF.||.LLJFFF77FJF7F7LJLJ||L7L-7|||L7|L7L7F-JLJFJL7||L-7LJFJL7F7|F-J|||FJ|||||||LJ||FJLJL--7L7JF-JF------JL-J|F----J|L-J||.|7|7L7.-|L|F|-LF-
.FJFF-F--L-FJL7L7|LJL7F-7|L7|F7|LJL7|L7|FJL-7F-JF-J|L-7|7FJF7||||L7FJ|||FJ||LJLJF-J|L7LF7F7L7L7|F7L-------7F-JL--7F-JF7FJ7FF777F77|.|LJLJ-|J
.|LLJJ|J.LJ|F7L7LJ|F7LJFJL7||||L-7FJ|FJ||F7FJL--JF-JF-JL7L-J|LJ||FJ|FJ|||FJL--7FJF7L7L-J||L-JFJLJ|F-------JL7F7F-J|F-JLJF-7|L7FJL7J-J||FJFL7
L|7--FJFFJ7LJ|FJF7FJL-7L-7|||||FFJ|FJ|FJ||LJF--7FJF7L-7FJJF-JF-J||FJL7|||L-7F-JL7||7L--7|L--7|F-7|L-----7-F7LJ||F7||F7|JL7|L7||F-JFL|F----FJ
|L-JLF7-7-F.F||7||L7F7L--JLJLJL7|FJL7|L7|L-7L-7LJ-|||FJL-7L-7|7FJ|L-7LJLJF-JL-7FJ||F-7FJ|F7L||L7||F-7F-7L-JL7FJ||||||L-7FJ|FJLJL--7.-7FJJ.LJ
L7|FLLJ.F777FJL-JL7LJ|F----7F--J||LL||FJL7FJF7|F-7|L7L7F-JF7|L7L7L-7L7F--JF7LFJL7|||FJL7LJ|FJL-JLJL7LJLL---7LJFJ|||||F-JL7|L-7F-7FJ-|.LJFF.|
FLL--J-FJL-7L---7FJF-J|7F--JL7F7|L7FJ|L7FJ|FJLJL7||FJFJL7FJLJFJ||F-JF|L7F-JL7L-7|||||F7|F-JL7F7F---JLF7F--7L--JFJLJ||L7|FJL--JL7||7-L7J7F.7|
F77L|J.|F-7L-7F7|L7L--JFJF--7LJ|L7|L7L7|L-J|F7-FJ||L7L-7|L--7|F7|L7F-JFJL7F-JF-J||||||LJL-7|LJ||F-7F7|||F-JLF7FJF--J|FJFJF-7F7FJ||-7JJFJL-J|
LL-F77FLJFJF7LJ||FJF7F7L-JF7|F-JFJL7|FJL--7LJL7L7LJFJ-FJL-7.||||L7|L-7|F7|L-7L-7LJLJ|L7F--JF-7|||FJ||||||F-7|||FJF7FJL-JFJFJ|LJJLJF7--J||.|7
F|-|L--F-JFJL-7LJ|L|LJL7-FJLJ|F7L-7|||F-7J|F--JFL-7L7FJF--JFJ||L7||7FJ||||F-JF7L--7FJFJL7F7L7LJLJL7||||||L7LJ||L7|LJF7F7L7|FJ.FF|7L|JL-77-F-
|.|..|7L7FJF7-L-7L7L--7L7L-7FJ||F-J|||L7|FJL7F7F7FL7|L7|F-7L7||FJ|L-JFJ|||L7J||-F7||LL7FJ|L7L-7F--J|LJLJ|FJF-J|FJ|F-JLJL-JLJLFF-77.J7|JL7FL7
F|-F|F7-||FJL7F7L7L--7L7||FJL7||L-7|||FJLJF-J|LJ|F7||FJLJFJFJ||L7L--7L7|||FJFJ|FJLJ|F7||FJFJ.FJ|F-7|F---JL7L--JL-JL7F-----7F77|FJ-JF-7.LFJFL
L|.F----J|L-7LJL7|F7FJL|L7L-7|||F7||||L--7|F7L7FJ||||L7F-J-L7||FJ.F7L7|||||JL7|L--7||||||FJF7L7|L7|||F7F7-|F------7LJF----J|L-JL7JJJ7|FL|-FJ
L|.L-----JF7L--7|LJ||F7L7|F7|||||LJLJ|F7FJ||L7||FJ|||FJL7F7FJ|||F7||FJ||LJ|F7||F7FJ||||||L7||FJL7|||LJ||L-JL7F---7|F7L7F-7-L7F--J.|-7JFL--J7
..7|FLF---JL7F7||F-JLJL-JLJLJLJ|L---7|||L7|L7|||L7LJ|L7FJ||L7||||||LJFJL7LLJ||||||FJ|LJ||FJ|||.F|||L-7LJLF7FJ|F--J|||FJ|FJF7||-LL---|.77L|F7
|J--FJL----7LJLJLJF-7F--7F7F7F7L----J|||FJ|FJ|||FJ-FJFJ|FJ|FJ||||||.FJF7L--7||LJ||L7L-7|||FJ|L7FJ|L7FJF-7||L-J|F7L|||L-JL-JLJL7JLLL7|7|JLJ..
LJ|FF-7F7F7|F-7F7FJL||F-J|||LJL-----7||LJ-LJFJ|||F-JFJF||FJ|FLJLJ|L7L7|L7F-J|L-7||FJF7||||L7|FJL-J||L7L7LJL--7LJL7LJL7F-7F---7L-7J.LL77.FL-J
|-F-L7LJLJLJ||LJ||F-J||F7|LJLF7F----JLJF---7L7||||F7L7FJ|L7L-7JFFJFJFJL7|L7LL7FJLJL7|||||L7||L7F---JFJLL----7|F-7L--7LJFJL--7|F-JLLJLJ||F||7
|F7J-L---7F7L7F7LJL-7|||LJ|F-J|L----7F7|F--JF||||||L-J|FJ7L-7L7FJFJFJF-JL7L-7|L7F7FJ||||L7LJ|FJL-7F7|F-----7|||FJF7FJF7L-7F7|||F7.|.FL--LL7J
LLJF7|LLFJ||FJ|L-7F7LJLJF--JF7L-----J|LJL---7LJ||LJF--JL-7F-JFJ|FJJL7L7F-JF-JL7LJ|L7||||FJF-JL-7FJ||||F---7LJLJL-JLJ|||F7LJLJ|LJL7.F7J-|-J.|
|J.7||.FJFJLJ-|F7LJL7F7FJF7FJL-7F7F7FJF-----JLFJL7L|F7F--J|F7L7|L7F-JFJL-7L-7FJF7L7||LJ|L7L-7F-JL7|LJLJF7FJF-7F7F7F7FJLJL-7-LL7F7L7---7|L|JJ
777LLF-JFJ7F--J|L--7LJ|L-JLJF--J|LJLJFJF---7F7|F7L7LJ|L7F-J||FJL7||F-JF--JF-J|FJL7||L7|L-JF-JL-7FJ|F---JLJ||FJ|||LJ||F----J7.F|||FJ7|FLL7|..
|-F77L-7|F7L-7FJF7-|F7L-7F7FJF-7|F---JFJF-7LJ|||L-JF-JFJL7FJ|L-7||LJF-JF-7L-7|L-7||L7L7F7-L-7F-J|FJ|F7F7F7FJL-JLJF7LJL-----7J-|||||7|-.LFL-J
|-FF-7L||||F7LJFJL-J||F7LJLJFJJLJL----JFJ.L--J|L-7.L-7L-7|L7|F-JLJF-JF7|FJF-JL7FJ|L7L7||L---JL-7||FJ|LJLJLJF7F7F7|L7F------JJ-||LJ--J.|-||L|
7.FJFJLLJ|LJL--JF---JLJL----JF7F----7F7L-----7|F7|F7FJF7|L-J||F-7LL-7|LJL7L7||LJL|FJFJ|L-7F7F7FJLJL-JF7F---JLJLJLJ7|L----7F7|FLJLL-F-.7-LL7L
|-|J-JFFF|F7F7F7|F-----------J|L---7||L---7F-J||LJ||L7|||F--JLJFJF--JL--7L7L7F--7|L7|FJF-J||||L7JF7F-J|L------77F-7L-----J|L--7L7LFLJFJ|.L|J
F.JFLFLLFJ||||||||F------7F7F7L----J|L---7|L-7LJF-JL-J||||F--7FJ|L7F7F-7||L7LJF-JL-J||FJF7|LJ|FJFJ|L-7L------7L7L7|F7F----JF--JF7L--FF--7.|7
L-.J.||LL-JLJLJLJLJF-----J||||F7F7F7L--7FJL--JF7L-7F--J||LJ-FJL7F-J||L7LJF-JF-JF--7FJ||FJ||F-JL7L7|F7L------7L7L7|||||F7F--JF--J|-.F|.-.F.F7
JF--7J---L.LF---7F7|F---7FJ||LJLJLJL--7LJF7F7J|L7FJL--7||F--JF-J|F-JL7L-7L-7L-7|F-JL-JLJFJ||F--JFJLJL-------J.L7LJLJ|||LJF-7|F7FJ|F-JFL-JLJ|
|FFLJLL|||.FL7F7LJ|LJLF7LJ.||F---7F--7L--JLJL-JFJL----JLJL7F7L7FJL-7FJF-JF-JF7LJ|F-7F---JFJ|||F7L---7F---7F77F7L----J|L-7L7|||LJ-|J|FJJL7.77
LJFJ7.--J77FFLJ|F7|F7FJL7F-J|L--7|L-7|F7F7F7F-7L--7F-7F7F-J||FJ|F7FJL7L-7L7FJ|F7LJFJ|F7F7L7LJFJL----J|F--J|L-JL-7|F--JF7L7|LJL77-|JLJFFFJFL7
|FJJL7FJ|LLF--7|||LJLJF7|L--JJF-JL--JLJLJ||LJ-L7F7LJFJ||L-7|||FJ||L7FJF-JFJL7|||F-JFJ||||FJF-JF--7F--JL---JF7F-7L-JF--JL-J|F--JF-7.|FL7J-7.|
L7FF-L-|-F7L-7|LJL-7F-J|||F7F7L--7F-7F7F7LJJF7-LJL-7L-JL-7||||L7|L7|L7|F7L7FJLJ||F7L7|||LJFJF7|-FJ|F7F-7F--JLJJL-7FJF-7-F-JL7F-JFJ-L77.||JF|
7|LJ|F.-7|FF-JL--77LJF-JL-J||L---J|FJ||||F7FJL----7L-7F7FJLJ|L7|L7||FJLJ|L|L--7|LJL7|||L7FJFJ||FJFJ|||FJL-------7|L7|FJFJF--J|F7L-7.LJ7L-77|
--J.||.|LLFL----7L-7FJF7F-7|L-----JL-JLJLJLJF-----JF7LJLJ7F-JFJ|FJLJL7F-JFJF--J|F-7|LJL-JL7L7||L7|FJLJL---------JL-J|L7L7|-F7||L--J7JLLL7L-7
|J|LF7F|-LFF--7LL-7|L-JLJFJL7F-7FF---7F-7F-7L------JL--7F7L-7|LLJF---JL-7L7L--7LJJLJF7F7|FJFJ|L7LJL--7F------7F7JF-7L7L-JL-JLJL--77FF-..LL|7
L.JF|FF7.F-JF7L7F7||F7LF7L7FJL7L7L--7|L7|L7|F-7F7F7F7F-J||.FJL--7L-7F-7FJL|F-7L7F7F-JLJL7L-J-L-JF---7||F--7F7LJL7|FJJ|F----7F7F-7|-J-|-F--|7
FJLLF7||.L--JL7|||||||FJL-JL7-|FJF7FJL-JL7|LJLLJLJLJ|L--JL7L-7F7L-7LJFJL-7|L7|FJ||L---7FJF-7F--7L7F7LJ|L-7||L---J|L-7|L-7F7LJLJJ||L|FL-JL|.|
J7F.|LJL7.F7FFJLJLJLJLJF---7L-JL7||L-7F-7LJF7F7F7F7JL-7F-7L7J|||F7L-7L7F7|L7|||FJL--7FJL7L7||F-JFLJL-7|F-JLJF7JF7|F-J|F-J|L----7LJJLF.L.F7F|
L-7FL7F7L-JL-JF-----7F7|F--JF--7LJL--J|FJF7|LJLJ|||F-7|||L7L7||||L7FJFJ||L7||LJL-7F7LJF7|FJLJL7F----7|||F7-FJL-JLJL7F|L--JF7F--J-|JF|FL-J7J.
L-JF-LJL7F-7F-JLF-77LJLJL7F-JF-JF7F---JL-JLJF7F-J|LJFJ||F-JFJ||LJ7LJFL7||FJ||F---J||F-J||L7F--JL---7|||LJL-JF-7F7F7L-JF-7FJLJF77JJFLJ7.|L777
.LJ.|.|LLJFJ|F-7L7|F7LF77LJF-JF7|LJF--7F7F-7||L--JF7L7LJ|F7L7LJJ.FLJ-|LJ||LLJL---7|||F-JL7||JF7F7F-J|LJF7F--JJLJLJL-7FJ|LJF--JL7--LFJ|F|7F7J
L7F|77F7|JL-J|FJFJLJL-J|F7-L--JLJJFJF7LJ||7LJL7F7FJL-JF|LJL7L777|J|7.|FFLJLF----7LJLJL--7LJ|FJLJ|L-7|F7|LJ7F--7F7F77||F---JF7F-J.F-|.F-|-L.|
LFJ.|FLF-----JL-JF----7|||F7F7F7F7L-JL-7LJF7F7LJLJF7FLF7F--JFJ7-7|FJ-|F-|--L7F-7L----77FJF7LJF--JF-JLJ|L-7FJF7LJLJ|FJLJF7F7||L--7-.L--..FJ-L
L7..FJ.L--7F--7F-JF---JLJLJLJLJLJL-----JF-JLJL---7|L7FJ|L-7FJ-F7LJLJ-LF7||JFLJFJF-7F7L-JFJL7FJ7F7L---7|F7LJFJL----JL--7|LJLJL--7|77J|7FF|7J|
||L.|L7LJJLJF-J|F-JF--7F-7F-7F7F7F------JF-7F----J|FJ|FJ|-LJJFLJ-J-|J||7---J7-L-JFLJL7F-JF7||F7||F7F-JLJL--JF7F7F-7F7LLJF7F7F7|LJ-L7||-FL7FF
FJ|F|7F7.|F7L--JL-7|F-J|-|L7||||||F-7F7F7L7LJF77F-JL-JL-7F77LLJ|.|JF.|J|.L||L-F---7F7LJF-JLJLJ|||||L--7F---7|||||FJ|L7F7|LJLJ|F-7|F-JJ-|||.|
..L7|F7FFF|L---7F7LJL7FJFJFJ||||||L7LJLJL7L--JL7|F--7F--J||-77|-7J.L7F.F-.F-..L7F7LJL--JF7F7F7LJLJ|F7.LJFF7LJ|||||7|FJ|LJF---J|FJ-7J.F-77J-F
--7|-FLLLFJF--7LJL7F-JL7L7|.LJLJLJFL---7.L7F7F7LJL-7|L-7FJ|7-JJJ.-F--J-LL7|-7.LLJ|F7F7F7|||||L7F-7LJ|F7F-JL--J|LJL-JL7|F-JF7F-J|FLF.-|JF7---
||L77JL||L-JF7|F-7LJF7FJLLJF-7F7F------JF7LJLJL----JL--J|FJJ-J7JLJFJLL..||7L-LFF7LJLJLJLJ||||J||-L-7|||L---7F7|F-----J||F7|||F-J7|F7FL.L7|L|
--.L|7J|F---JLJ|FJF7|LJF7F7L7|||L-------JL7F7F----7F7F--JL7J7.-7L||7.F|7LJF7-|FJL--7F----J|LJFJ|F--J|||F7F-J|LJL------JLJLJLJL--77|JF77|.|7J
FLF.LF.-L------JL-JLJF-JLJL-J||L7.F---7F7FJ|LJF---J|LJF---JL-7|FF-F7.|L7J.LJF-L---7|L---7FJF-JFJL--7LJLJLJ|FJF7F------7F--------J7|F-77J-L7J
L77.L|FFJLF-7JF7F7F7JL--7F--7|L7L7L--7LJLJFJF-JF7F7L7FJ|F7|.|LJ7JJ|JF|F|7--L|.LF--JL----J|FL7FJF---JF7F---7|FJLJF7F-7FJL-7F--7F--7-L7L7JF.|7
FJ-7.J7|J|L7L-JLJ|||F7F7LJF-JL7L7L--7L----JFJFFJLJL-JL--JL--77.|F-|7|7F-...|L7.L-----7F7FJF-J|JL----JLJF--J||F7FJLJFJL7F7||F-J|F-JF7|FJ|LFL7
-J-7-L7L-7F|F7F-7LJLJLJ|F7L--7L-JF-7L---7F-JF7|F----7F-7F---J77-||.L.J7.F7-JJFFF7F---J||L7L--JF-7F-7F7LL---JLS||F--JF7LJLJ|L7.||F7||||-7-7--
L|L7FL|FFF-J|||FJF7F--7LJL--7L---JJL---7LJ|FJLJL---7LJ.LJ-F7F77F|-7FL---J77.FF7||L---7|L-JF--7|FJL7LJL-----7F7LJL--7|L-7F7L7L-JLJLJLJL-7F7J7
.|FLF-|F-|F7|LJL-J||F7L--7F7L-----7F7F-JF-7L7F-----JF7-FF7||||FF7-77JFJFLF7F7||||F---J|JF7L-7LJ|F7|F-7F---7LJ|F7F--J|F-J||FJF7F7F7F-7F7LJ|.F
FF|J.LLFJLJLJF77F-J||L---J||F7F--7LJLJF-JFJFJL----7.|L7FJLJLJL7||-|7-|---|LJ||LJ|L-7F7|FJL-7L-7LJ|LJFJ|F--JF7LJ|L--7|L-7|LJFJ||||LJFJ|L--J--
-JJ.|7|L|JJFFJL-JF-JL7F7F7|||LJF7L--7FJF-J7L7F----JFJFJL7F--7FJ||JLJ-7|7LL7FJL-7|F7LJLJL--7|F7L-7L--JFJL---JL-7|F7FJ|F7LJF7|JLJ|L-7L7L-7LL.|
.|-JJLJ||L--L7F7FJF7LLJLJLJLJF-JL-7FJ|FJF7F7||F7F77L7L--JL-7LJFJL-777|FF--JL7|FJLJ|JF7F7F7||||F7L---7|F-------JLJ|L-J|L7FJ|L7F-JF7L7L7FJ-|F|
7-JL7-LL|LLJLLJ||FJ|F7F7F----JF--7|L-JL7|LJLJLJLJL7-L7F---7L7LL--7|F7-FL---7L7|F-7L7|||LJLJLJLJL7F7FJ||F--------7L7F7L7|L7|FJL7FJ|FJFJ|J.|FJ
L|7LL-7J.|.|FFFJ|L7LJLJ|L-----JF7|L----J|F--------JF7|L--7L-J-.F-JLJ|F7F7F7|FJLJ-L7LJLJF7F-7F7F7LJLJFJLJF--7F--7|J|||FJL7||L-7LJFJ|.L7|.FF7|
JL|.JFJ.77.7.FL-JLL7F-7L--7F---JLJF--7F7|L-7JF7FF7FJLJF--JF7J|FL-7F-J|LJ||LJL7LFF7L--7FJLJ|LJ||L7F-7|F7.L-7|L7FJL7||||F7|||F-JF7L7L-7LJ---|J
F7L-J|L-JF777F.F7F7LJ.L--7|L------JF-J|LJF7L-JL-JLJF--JF7FJL77F7FJ|F7L7FJL--7L7F||F-7LJF7F7F7LJFJ||LJ|L---JL-JL-7||||LJ|LJLJF7||FL7FJ|||J|J|
7.L|FFJJFFJF---JLJ|F7F7F-JL--------JF7|F-J|F7F7F-7FJF-7|||F7|FJ|L7|||FJ|F7F-JFJFJ|L7|F7|||||L7LL7|F--JF7F7F7F7F7LJLJ|F7L----JLJL-7LJL|-7L7-J
JFFF-|-FF-7L---7F7LJLJLJF7F-7F-7F---JLJL7FLJLJ|L7|L-JFJ||LJ|LJFJFJ||||FJ||L-7L7L7L7|||LJLJLJFJF7LJL---JLJ||LJLJ|F-7FJ|L7F7F7F7F-7L777JJ|7|.7
FFLJJF7..LF|F|FJ|L--7F-7|||FJ|LLJF7F----JF--7-L7||F-7L7|L-7L-7|LL7||||L7||F7|FJFJFJ||L----7FJ-|L----7F--7LJF---J|FJL7L7LJLJLJLJ-L7L7J7FLF-77
F7L--.|.--F---JFJF--J|FJ|||L-JF--JLJF----JF7L-7|||L7L-J|F-JF7||F7|||||FJ|||LJ|-|FJFJ|F7F7FJL-7L----7LJF7L--JF7F7||F7L7L7F7F-----7L-JL7.J.J|.
7J|FL-FJLLL7F7FJ-L---JL-J||F-7L-----JF7F7FJL7FJLJL7L7F7|L-7|LJ||||LJLJL7|||F7L7|L7L7|||||L7F-JF----JF7|L----JLJLJ||L7L7LJLJF----J|.|-77LL.77
|7-|-|FJ.|.LJLJF7F------7LJL7|FF-----JLJLJFFJL--7-L-J|LJF-JL-7LJ|L7F---J||LJL7|L7|FJLJLJL-JL77L7F---JLJF--7F-----J|FJFJF7F7L-----7.LFJJ7LJ..
LL.|L-J|F-JJF7FJLJF7F--7|JF-JL-JF---------7|F7F-JF--7L-7|F7F7L-7|FJL7F-7|L7F-J|L||L--7F---7FJF7LJF77F7-L-7LJF-----JL7L7||||F--7F-J--F7.FF|.L
FL-F-J|7-||F|LJF7FJLJF-JL7L7F7F7|F--------JLJ|L-7L-7L7L|||||||FJ||F-J|FJ|FJL-7|FJL--7LJF7.|L-JL7FJ|FJL---JF7|F-7F7F7L7|||||L-7LJJ..-77.-7F-7
.LFL|LFF7JF7L--JLJF-7L-7FJFLJ||LJL-7F------7JL-7|F7L7L-JLJLJL7L7||L7FJ|FJL7F-JLJF-7FJF7||FJF-7FJL7|L------J||L7LJLJ|FJ|||||F7L77777LLF-.--|7
F.L7|FF7|FJL7F7F7-L7|F7|L7F--JL7-F7LJF-7F--JF7LLJ|L7L---7F-7FJFJ|L7|L7|L7FJ|F7F7L7||FJLJLJFJJLJF7||F--7F7F7LJFJJF7JLJLLJ||LJL7L-777.FF7.||JL
J.|L7L|L7L-7LJ||L--JLJLJFJL----JFJL--JFLJF7FJL-7FJFJF7F7LJFJ||L7|FJL7|L-JL7LJ|||FJLJL----7|F7F7||||L-7LJLJ|F7|F-JL-----7|L-7FJF-J|-|-J-FJL7J
FF7||7L7|F7|F7LJF7F-7F-7L7F--7F-JF-7F-7F-JLJF7FJL7|FJLJ|F7L7L77|LJF7|L--7FJF7LJ|L-7FF7F--JLJ||||LJL-7L7F-7LJ|LJF-------JL--JL-JL|-7J7|.JFJ.|
FJFF-7FJLJ|LJL7FJLJ7LJ7L7|L-7|L7FJFLJFJL7F--JLJF-J|L--7|||FJFJFJF7|||F7FJL-J|F-JF-JFJ||F-7F-J||L-7F-JJLJ7L-7|F7|F7F7F7F---7F-7-L.F7L|.F-LJ7.
|LFL7|L--7L7F7LJ.F------JL--JL7LJF7F7L-7|L-7F-7L-7|F77|||||FJ-L7|LJLJ|||F---JL-7L-7L7|LJFJL7||L7FJ|F7F7JF--J|||LJLJLJLJF--J|FJ-|-LJF|-JJ|-77
J-F-J|F7FJFJ||F-7|F--7F------7|FFJLJL--J|F-J|FJF7||||FJLJ|||F--JL7F7FJ|||F7F7F7L7FJFJL7FJF7L7|FJL7||LJL7L7F7LJL7F-7F---JF7F|L77|.L.-|77FFF-7
LFJF7LJ|L7|FJ|L7|LJF7LJF7F---JL-JF-7F7F7|L-7|L7||||||L7F7||LJF---J||L7||LJ||||L7|L7|F-JL-JL7||||FJ|L7F-J7LJL--7|L7|L---7||FJFJ-J7--.LF7--JLJ
.L-J|F7L7||L7|FJL7FJL--JLJF7F-7F-J7LJ||||F-JL7|||||||J||LJL7FJ7F7FJ|FJ||F-J||L7||FJ|L7F-7F-JLJL7L7L7||-F------JL-JL7F7FJ||L7|F7J77|-FL--J..7
-JJJ||L7LJL-JLJF-JL------7||||||F-7F-J|||L7F7|||LJLJL7|L7F7|L7FJ||FJL7|||F7|L-J|||FJFJL7|L-7F--JFJFJ|L7L---7F-----7||LJFJL-JLJL77|F7L-J7F|7|
||7L||JL7F7F7F7L7F7F7F7F-J|||FJ|L7LJF7|||FJ||||L---7FJ|FJ|||FJL7|||F7||||||L-7FJ|||FJ7.||F7||7F7|FJL|FJF7F7LJF----J|L7FJF7F7F-7|L7LJJJLF|J|J
L--FLJF7||LJLJL7||||LJLJF7|||L7|-L7FJ||LJL7|||L7F-7|L7|L7|||L7FJ|||||||||||F7||FJ|||F7FJ||||L7||||F7|L7||||F7L--7F7L-J|FJLJLJ.LJ7F7.|-L-L7||
FJ.LF-JLJL-7F--JLJLJF7F7|LJ||FJ|F-JL7|L-7FJ||L7|L7LJ|LJFJ||L7LJ|||||||||LJ||||||FJ||||L7||||FJ|||||||FJ|LJ||L7F-J|L-7J|L---7J-|L|L-77...FL|7
LF7JL--7F-7|L---7F7FJLJ||.FJ|L-JL7F7||F-JL7||FJL7L----7|FJ|FJF--J|||||||F-J|||||L7|||L7|LJ||L7|LJLJ||L7L-7LJLLJF7|F-JFJF---J|JJL.7L|L7|-7J||
.L7FLLFJL7||F--7LJLJF--JL7|FJF---J|LJ||F--J||L-7|F7F7FJ||.|L7L-7FJ||||LJL-7|||||FJ|||FJL7-LJFJL-7.FJ|F|F7|F----JLJ|F7|FJF7-F777.F7|F|7L.|7F|
L-L-L.L7FJLJL-7L----JF7F-J||JL-7F7L7JLJL7F7|L7FJ|||||L7|L7|FJF7||.||LJFF--J|LJLJL7|||L7FJF--JF-7L7|FJFJ||LJF7F----J|LJL-JL7|L--7LL--7JJ-FFF|
|.L.FFL||F---7|F--7F7||L77LJF-7LJL-JF7F-J||L7|L7||||L7|L-JLJJ|LJL7||F--JF-7L--7F7|||L7|L7L7F7L7L7||L7L7||F-J|L-----JF7F--7LJF--J-L.FL||.7.LJ
JJ.|7FFJ|L7F7LJ|F-J|||L7L--7|FJF----JLJF-J|FJL7|||||FJL---7|FJF7FJLJ|F7FJFL7F-J|LJ|L7||FJ|LJ|FJJLJL7L7|||L-7L--7F7F7||L-7L--JFJJJJFL7L7-L7..
|JF|--L-JJLJL-7||F7||L7L7F7||L7L----7F7|F-J|JFJ|LJ||L7F7F7L7L7||L--7||||F--JL7JL7FJFJ||L7F--JL7F---JFJLJL--JF7L|||||||F7L-7JJJ-JJ-||-LJFL77L
F--FJJ||||F---J||||||FJFJ|||L7L----7||LJ|F7L7L7L77LJF||||L7L7LJL7F-JLJ||L7F-7L-7|L7L7||FJL7F7FJ|F7F7L-------JL7LJLJLJ||L7FJ7LL|L|JLLJ|FL7JLJ
7-|J|.7L--|F7F7|||||||-|FJLJ.L----7LJL-7LJL7|FJFJF---J||L7|FJF--JL---7LJFLJ7|F7||FJFJ|LJJ-LJ||JLJ|||F7F7F7F7F-JF7F--7LJ|||FJJFF-7|-F-F--L7LL
-7FF-7|-L-LJ||||||LJLJFJL-7F-7F---JF--7L7JFJ||FJ|L7F7FJ|FJLJ.|F7F7F--JF-7F-7||||||-L7L7F7F7FJL7-FJ|LJ|||||||L--JLJF7L-7FJL7LFJL-|-77.L.LL--7
LLFJLL-7F7.F||||||F-7FJF7FJL7||F-7FJF7L7L7L7||L--7LJ||||L-7F7|||||L---JFJL7LJ|||LJF-JFJ|LJLJF-JFJFJ7FJ|||||L7F-7F-JL--JL7FJ7JL7-L-LF--|JJ|L-
|L|.F7.-|77-LJLJ||L7LJFJLJF7|LJ|FJL-JL7L7L7|||F-7|F-J|FJF7LJ|||||L7F---J7JL7FJ||F7|F7L7|F--7L7FJFJF-JFJLJLJFJL7|L------7LJLJ7|.7JL|.L-J..7J.
L7L-LF7JL.|.L|LJLJFJF7|F--JLJF-JL7F-7FJFJFJLJ|L7||L-7|L7|L7FJLJ|L7|L--7F77FJL7|||LJ||FJ||F-JFJL7L7L-7|F7|F-JF7|L7F----7|F|.L|FFL7.-7-J.F7J|L
|JF-FL7.|F.F-JF7LLL-J|||F----JF7FJL7|L7L7L--7|FJ||F-J|FJL7|L--7L-JL7F7LJL7L7FJ|||F7||L7||L7FJF-JFJ-FJLJL7L7FJ|L7|L--7-LJ-7--.|J.L7-|-|7FJ.||
|FF-JL-7-L|-.FF7.F7F-J|||F----J|L7FJL-JFJF7FJ|L7||L7FJL7FJL7F7L--7J|||F-7|FJL7||LJ|||FJLJFJ|||F-JF7|F-7FJ.|L7L7|L7F-J7-LL|-.LJFF.-J.FJJL7.|7
F-J|J.L7-JJ..LJ7.LFL7FJLJL--7F7L-J|F-7FJFJLJF|FJLJFJ|F-JL-7||L7F7L7LJ||.LJL7FJ||F-J||L--7L7L7|||FJLJL7|L-7|FJFJ|FJL--7.L-JL|.LL|7LF7|L7.|-J|
.LF|.-.J.|F77-|LJ-FFJL7J-F--J||F-7LJFJL7|F7F-JL--7L7|L-7F-J|L7||L7L-7|L---7|L7LJL7FJ|F--J7L-JLJFJF-7FJ|F-JLJFJFJL7F7FJ7.L77.F|FLJ--F-.J-|JL7
|7FJJJ..FF-||FJ.LJLL-7|F-JF--J|L7|F7|FFJLJ|L7F---JJ||.FJL7FJFJ||||F7|L-7F-JL7L7F-JL7|L-7LF7F7F7L7||LJFJL---7L7L-7||LJJ|7---|7.F|F7|.FF|7FF7.
|F|7|F7FF-7F-7.L77|.|||L--JF--JFJ|||L7L7F-JFJL----7LJFJF7|L7L7||FJ||L7FJL7LFJFJL-7FJ|F-JFJLJLJL-JL7|FJF-7F7L7|F7||L7JLFJ-J.|77-LLJFLFJJFLJJ7
L77J-J|F.LFJ.F-FJJL|FLJF---JF-7|-|||FJFJL-7L7F----JF7L7||L7L7|LJ|FJL7||F-JLL-J|F-J|FJL-7L7F7F7F-7FJFJFJFJ|L7||||||FJJ.|-LLFL|JJ7..FLL7.||JF|
L|FJ|L|L-7JF-7.7J77|L|-L7F7FJFJL7||||LL7F-JL|L----7|L-J||FJ||L7-||FL|||||.|-FL-L-7||F7FJFJ|LJLJFJL7L7|FJFJ|||LJLJLJ|.FJ7L--J|FL|FJ-.FL7L7.||
.L77|-.FF|-J7J.|.L-L-7JFJ|||.L-7||||L7FJL7JFJF7F--JL7F7|LJ7FJFJ-||-7|||L7JJ7F|JFFJ|||LJ7L-J-F--JF-JF||L7L-7|L-7LLJ.J-L-7.L--7|.LJJJ-JL--J.--
JF|--JFF-|L|L7-77.FF7J.L-J||F--J||||FJ|F-JFJFJ|L--7|LJ|L-7FJFJJFJ|.-LJL-J7|L77.FL7||L7--|J||L7F7L7FFJL7|F-JL7FJ-J-L.L||.7F7|L77..|-7|7.L|F|J
L-J|L--JF7.-JL-.777LJLF.LFLJL7F-J||||FLJJJL-JL|F--JF77|F-JL-J.LL7|..LJ||LF7JL77F-|||FJ|||JL|FLJL7L7L7FJLJ|-LLJ.|..LF.F77FF|7FJ7--7|FJF-FL-L.
--L|LF|7J|7.|7FLLLFJ|7L|LF-|L|||7LJ||J-L||LF--JL--7||FJ|.|-|F7.|LJ77.L|-LJLF7L|L-LJLJ.F7..FJJ.F-JFJF|||JF--LJ7JJF.-F7J77FJL-|JL7J.|L7J--JJFJ
LF-7F7L-FJ|7LFF.|J.|||.LFLF7.LJ7-F-J|JFLF-JL7F7F-7LJ|L-J-J.LFJ-||7.--.|LLF.||-|7.LFJ-FJF7F-..FL--JFF|L7JFJ7.FL.LFL-JJ-L--J-F|-LJ7FF|JF-7|7J.
||-LF7JL|7F-.FJL-.|LL-7FFJ||F7LJLL7FJFLFJ.--LJ||F|F-JJFL|-F|.-7.L77-L7J.||FLJF7777L7|.-7FJ|LL7|FFJFLL-J7L77F7LF-|7|J-7|LJ.FLJ7.|F|JL|J.7JF.F
JF-7LFJ.||7|7|7JF-JL|L-FJ-J-LJ.F|FLJL7J|F7F|LFJ|FJL-7-L7FJLFJLF77|7F7J7F77-JL||F-J--L.LJJL7FL-L-7.L|F|.|7LJJ7-|-L|7J7|7.|FF...--|||F7JFL.L-L
LF-F7.FF|L|FFL-.-.||.FL|-FLJ.L7L-77|-L7-J|LF-JFJL7F7L7|F|-7L|JF7-777.|LJJ||LJL7FJ.|L7LJJ77.F-JLLL---7-JLF.|L|--.|LF7LF|7JF--|JJLJJ-7LFJ|F.||
LJ|L|F77L-LF7LJF.FF7-J.L.|--7.||FL7JFJ.LFL.L-7L7FJ|L-J-F7|--|.L--LJ|.-7|.|FL7FLJ7-|JF|7LFJ---LFF7.FFJJ--.-FJLF|-F-J7|FLJFJ7F-JFJ|.-JJL-JJ7J7
|.|.L|L7J.|L77F-7|-||.F|-FFJ.|J-7LJ7|F7-J.FJ-L-J|FJJL|L---7J|-7|FJJ7|FJ-FJFLJ-|JJ7L--FFJLJ7J--7JLL7|J.F-|7|.FFL.|FFLL7L|JFLJ|7J.F-F-7|-7.F.7
-.F7-LJ|J.|.J7JLFJL|F7|-FF.F-J.L..LF7|-.|.7.FJJ|||.F|.F|JL-7F-J-J|.|777-7-F.L.L|.7-77||L|L.F-LL7FLFJ-||FL-|F7-|7|77L||7|||.-F---L-|L-JF7F|FJ
|FL|7..|-FFF7JF7.77|JF|.J.-J|LL--7JLJJLL-77F||.FLJ7F7--.L|.F--JJ-FLJFF-7J-L|J.FLJ-77--77FJFFJ77F|-7.7L-JL.LFL-77JLJ-J77|7.FJ|LLJL7.L|.JL|-|.
JLLLJ.L-F7-JJJL--|-|-JLFJ7-F7-L-LL.JJJJ.J.L---JLL|LL--L-.J-L|-.--LL-|JJ|-----JLL|.L7-LLL.|-F--7-7-J|JLJJ|J-J-LL|.LLL.|JJF.|..LJ-|J-.J.L7JJLL`
