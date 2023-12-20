// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(solve(example, 1000))  // 8000 4000 32000000
	fmt.Println(solve(example2, 1000)) // 4250 2750 11687500
	fmt.Println(solve(input, 1000))    // 17913 50172 898731036
	fmt.Println(solve(input, 0))       // 229414480926893

}

type state struct {
	from  string
	to    string
	pulse rune
}

func solve(in string, epoch int) int {
	clear(flipFlopState)
	clear(conjunctionState)

	m := parse(in)
	mods := m["broadcaster"]

	var lo int
	var hi int
	var iter int
	//target := epoch
	for _, v := range m {
		for _, s := range v {
			if _, ok := m["&"+s]; ok {
				// Is conjuction.
				// Find all the nodes.
				for kk, vv := range m {
					if slices.Contains(vv, s) {
						if _, ok := conjunctionState[s]; !ok {
							conjunctionState[s] = make(map[string]bool)
						}
						kk = strings.TrimPrefix(kk, "%")
						kk = strings.TrimPrefix(kk, "&")
						conjunctionState[s][kk] = false
					}
				}
			}
		}
	}

	var points = make(map[string]int)
	var cond func() bool
	var do func()
	if epoch == 0 {
		cond = func() bool {
			return true
		}
		do = func() {}
	} else {
		cond = func() bool {
			return epoch > 0
		}
		do = func() {
			epoch--
		}
	}
	for cond() {
		iter++
		lo++ // Initial pulse from the button is always low.
		do()

		var pulses []state
		for _, mod := range mods {
			pulses = append(pulses, state{from: "broadcaster", to: mod, pulse: 'l'})
		}

		isFlipFlop := func(s string) bool {
			_, ok := m["%"+s]
			return ok
		}

		isConj := func(s string) bool {
			_, ok := m["&"+s]
			return ok
		}

		for len(pulses) > 0 {
			var h state
			h, pulses = pulses[0], pulses[1:]
			switch h.pulse {
			case 'l':
				lo++
			case 'h':
				hi++
			default:
				continue
			}

			if epoch == 0 && (h.from == "dc" || h.from == "rv" || h.from == "vp" || h.from == "cq") && h.pulse == 'h' {
				if len(points) == 4 {
					var i = 1
					for _, p := range points {
						i *= p
					}
					return i
				}
				points[h.from] = iter
			}

			mods, ok := m["%"+h.to]
			if !ok {
				mods = m["&"+h.to]
			}

			pulse := h.pulse
			if isFlipFlop(h.to) {
				pulse = flipFlop(h.to, h.pulse)
			} else if isConj(h.to) {
				pulse = conjunction(h.from, h.to, h.pulse)
			}

			for _, mod := range mods {
				pulses = append(pulses, state{from: h.to, to: mod, pulse: pulse})
			}
		}
	}
	fmt.Println(lo, hi)

	return lo * hi
}

func parse(in string) map[string][]string {
	m := make(map[string][]string)

	lines := strings.Split(in, "\n")
	for _, line := range lines {
		lhs, rhs, ok := strings.Cut(line, " -> ")
		if !ok {
			panic("invalid")
		}
		m[lhs] = strings.Split(rhs, ", ")
	}

	return m
}

var flipFlopState = make(map[string]bool)
var conjunctionState = make(map[string]map[string]bool)

func flipFlop(to string, pulse rune) rune {
	if pulse == 'h' {
		return '0'
	}
	flipFlopState[to] = !flipFlopState[to]

	if flipFlopState[to] {
		return 'h'
	}

	return 'l'
}

func conjunction(from, to string, pulse rune) rune {
	if _, ok := conjunctionState[to]; !ok {
		conjunctionState[to] = make(map[string]bool)
	}

	conjunctionState[to][from] = pulse == 'h'

	for _, v := range conjunctionState[to] {
		if !v {
			return 'h'
		}
	}

	return 'l'
}

var example = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var example2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

var input = `%vh -> qc, rr
&pb -> gf, gv, vp, qb, vr, hq, zj
%zj -> kn, pb
%mm -> dj
%gp -> cp
&dc -> ns
%qc -> gp
%dx -> fq, dj
%tg -> nl, ks
%pr -> nl
%gx -> xf
%hd -> lt, nl
%dq -> dj, jc
%ht -> jv
%bs -> pb, rd
&nl -> ks, cq, tc, xf, gx, hd, lt
&dj -> dc, fq, jz, ht, zs, jc
&rr -> gp, rv, jt, qc, sq
%vr -> qb
%jz -> dj, ht
%hq -> nx
%cf -> jg, rr
%hj -> cf, rr
%mt -> rr
%sq -> rr, vh
%jg -> rr, pd
%gf -> gv
%xv -> dj, dx
%rh -> nl, gx
broadcaster -> hd, zj, sq, jz
%jv -> dj, zs
%rd -> vs, pb
%pd -> rr, mt
&rv -> ns
&vp -> ns
%vs -> pb
%nx -> pb, bs
%zp -> mm, dj
&ns -> rx
%lt -> rh
%pf -> pr, nl
%tc -> qz
%xz -> dj, zp
%qb -> hq
%rl -> pf, nl
%fq -> xz
%kn -> pb, xn
%xf -> tg
%qz -> nl, rl
%ks -> tc
%jt -> kb
%jc -> xv
%kb -> hj, rr
%zs -> dq
%gv -> vr
&cq -> ns
%cp -> rr, jt
%xn -> pb, gf`
