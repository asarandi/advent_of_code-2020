//advent of code 2020, day 24, part 1 and 2
package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

var (
	input = []string{}
	tiles = map[image.Point]uint{}
	steps = map[string]image.Point{
		"e":  image.Point{0, 2},
		"se": image.Point{1, 1},
		"sw": image.Point{1, -1},
		"w":  image.Point{0, -2},
		"nw": image.Point{-1, -1},
		"ne": image.Point{-1, 1},
	}
)

func count() uint {
	var res uint
	for _, v := range tiles {
		res += v
	}
	return res
}

func part1() uint {
	for _, line := range input {
		t := image.Point{}
		for i := 0; i < len(line); {
			if s, ok := steps[line[i:i+1]]; ok {
				t = t.Add(s)
				i += 1
			} else {
				t = t.Add(steps[line[i:i+2]])
				i += 2
			}
		}
		tiles[t] ^= 1
	}
	return count()
}

func part2() uint {
	for i := 0; i < 100; i++ {
		missing := map[image.Point]uint{}
		for k, v := range tiles {
			if v == 1 {
				for _, s := range steps {
					t := k.Add(s)
					missing[t] = tiles[t]
				}
			}
		}
		for k, v := range missing {
			tiles[k] = v
		}
		nextgen := map[image.Point]uint{}
		for k, v := range tiles {
			var u uint
			for _, s := range steps {
				u += tiles[k.Add(s)]
			}
			if (v == 1 && (u == 0 || u > 2)) ||
				(v == 0 && u == 2) {
				v ^= 1
			}
			nextgen[k] = v
		}
		tiles = nextgen
	}
	return count()
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	input = strings.Fields(string(data))
	fmt.Println("part 1:", part1())
	fmt.Println("part 1:", part2())
}
