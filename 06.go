// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var digits = regexp.MustCompile(`\d+`)

func main() {
	fmt.Println(parse(example))
	fmt.Println(parse(input))
	fmt.Println(parse2(example))
	fmt.Println(parse2(input))
}

func hold(t0, t1 int) int {
	return t0 * t1
}

func eval(times, dists []string) int {

	var total int = 1
	for i := 0; i < len(times); i++ {
		t := toInt(times[i])
		d := toInt(dists[i])

		var count int
		for i := 1; i < t; i++ {
			if hold(i, t-i) > d {
				count++
			}
		}
		total *= count
	}
	return total
}

func parse(input string) int {
	lines := strings.Split(input, "\n")
	times := digits.FindAllString(lines[0], -1)
	dists := digits.FindAllString(lines[1], -1)
	return eval(times, dists)
}
func parse2(input string) int {
	lines := strings.Split(input, "\n")
	times := strings.Join(digits.FindAllString(lines[0], -1), "")
	dists := strings.Join(digits.FindAllString(lines[1], -1), "")
	return eval([]string{times}, []string{dists})
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `Time:      7  15   30
Distance:  9  40  200`

var input = `Time:        61     67     75     71
Distance:   430   1036   1307   1150`
