//advent of code 2020, day 12, part 1 and 2
package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

var (
	puzzle = []string{}
	turns  = map[string]int{"L90": 3, "L180": 2, "L270": 1, "R90": 1, "R180": 2, "R270": 3}
	dirs   = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} /* URDL, NESW */
)

func part1() int {
	p, d := image.Point{}, 1
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
		p = p.Add(dirs[i].Mul(n))
	}
	return abs(p.Y) + abs(p.X)
}

func part2() int {
	p, w := image.Point{}, image.Point{10, -1}
	rot := func(i int, p image.Point) image.Point {
		return map[int]image.Point{
			1: image.Point{-p.Y, p.X},
			2: image.Point{-p.X, -p.Y},
			3: image.Point{p.Y, -p.X},
		}[i]
	}

	for _, s := range puzzle {
		if _, ok := turns[s]; ok {
			w = rot(turns[s], w)
			continue
		}
		n, _ := strconv.Atoi(s[1:])
		if i := strings.Index("NESW", s[:1]); i != -1 {
			w = w.Add(dirs[i].Mul(n))
		} else {
			p = p.Add(w.Mul(n))
		}
	}
	return abs(p.Y) + abs(p.X)
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	puzzle = strings.Fields(string(data))
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
