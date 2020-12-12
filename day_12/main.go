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

func (p point) r0() point {
	return point{p.y, p.x}
}

func (p point) r90() point {
	return point{p.x, -p.y}
}

func (p point) r180() point {
	return point{-p.y, -p.x}
}

func (p point) r270() point {
	return point{-p.x, p.y}
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

var puzzle = []string{}

var turns = map[string]int{
	"L90":  3,
	"L180": 2,
	"L270": 1,
	"R90":  1,
	"R180": 2,
	"R270": 3,
}

var nesw = map[byte]point{
	'N': {-1, 0},
	'E': {0, 1},
	'S': {1, 0},
	'W': {0, -1},
}

var dirs = []byte{'N', 'E', 'S', 'W'}

var rots = []func(point) point{point.r0, point.r90, point.r180, point.r270}

func part1() int {
	var (
		p, q point
		ok   bool
		d    = 1
	)
	for _, s := range puzzle {
		key := s[0]
		val, _ := strconv.Atoi(s[1:])
		if _, ok = turns[s]; ok {
			d = (d + turns[s]) % 4
			continue
		}
		if q, ok = nesw[key]; !ok {
			q = nesw[dirs[d]]
		}
		for ; val > 0; val-- {
			p = p.add(q)
		}
	}
	return abs(p.y) + abs(p.x)
}

func part2() int {
	var (
		w    = point{-1, 10}
		p, q point
		ok   bool
	)
	for _, s := range puzzle {
		key := s[0]
		val, _ := strconv.Atoi(s[1:])
		if _, ok = turns[s]; ok {
			w = rots[turns[s]](w)
		} else if q, ok = nesw[key]; ok {
			for ; val > 0; val-- {
				w = w.add(q)
			}
		} else {
			for ; val > 0; val-- {
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
