//advent of code 2020, day 12, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	y, x int
}

func (p point) add(q point) point {
	return point{p.y + q.y, p.x + q.x}
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

var (
	puzzle = []string{}
	turns  = map[string]int{"L90": 3, "L180": 2, "L270": 1, "R90": 1, "R180": 2, "R270": 3}
	dirs   = []point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func part1() int {
	p, d := point{}, 1
	for _, s := range puzzle {
		if _, ok := turns[s]; ok {
			d = (d + turns[s]) % 4
			continue
		}
		i := strings.Index("NESW", s[:1])
		if i == -1 {
			i = d
		}
		n, _ := strconv.Atoi(s[1:])
		for ; n > 0; n-- {
			p = p.add(dirs[i])
		}
	}
	return abs(p.y) + abs(p.x)
}

func part2() int {
	p, w := point{}, point{-1, 10}
	rot := func(i int, p point) point {
		return map[int]point{
			1: point{p.x, -p.y},
			2: point{-p.y, -p.x},
			3: point{-p.x, p.y},
		}[i]
	}

	for _, s := range puzzle {
		if _, ok := turns[s]; ok {
			w = rot(turns[s], w)
			continue
		}
		n, _ := strconv.Atoi(s[1:])
		i := strings.Index("NESW", s[:1])
		if i != -1 {
			for ; n > 0; n-- {
				w = w.add(dirs[i])
			}
		} else {
			for ; n > 0; n-- {
				p = p.add(w)
			}
		}
	}
	return abs(p.y) + abs(p.x)
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	puzzle = strings.Fields(string(data))
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
