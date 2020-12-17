//advent of code 2020, day 17, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type vec3i struct {
	x, y, z int
}

type vec4i struct {
	x, y, z, w int
}

var (
	xyz      = []int{-1, 0, 1}
	puzzle3i = map[vec3i]int{}
	puzzle4i = map[vec4i]int{}
)

/*
*   TODO: remove repetitive code
 */

func neighbors3i() map[vec3i]int {
	res := map[vec3i]int{}
	for k := range puzzle3i {
		for _, x := range xyz {
			for _, y := range xyz {
				for _, z := range xyz {
					c := vec3i{k.x + x, k.y + y, k.z + z}
					if c != k {
						res[c] += 1
					}
				}
			}
		}
	}
	return res
}

func part1() int {
	for i := 0; i < 6; i++ {
		nextgen := map[vec3i]int{}
		for k, v := range neighbors3i() {
			_, exists := puzzle3i[k]
			if v == 2 && exists {
				nextgen[k] = 0
			} else if v == 3 {
				nextgen[k] = 0
			}
		}
		puzzle3i = nextgen
	}
	return len(puzzle3i)
}

/*
*
 */

func neighbors4i() map[vec4i]int {
	res := map[vec4i]int{}
	for k := range puzzle4i {
		for _, x := range xyz {
			for _, y := range xyz {
				for _, z := range xyz {
					for _, w := range xyz {
						c := vec4i{k.x + x, k.y + y, k.z + z, k.w + w}
						if c != k {
							res[c] += 1
						}
					}
				}
			}
		}
	}
	return res
}

func part2() int {
	for i := 0; i < 6; i++ {
		nextgen := map[vec4i]int{}
		for k, v := range neighbors4i() {
			_, exists := puzzle4i[k]
			if v == 2 && exists {
				nextgen[k] = 0
			} else if v == 3 {
				nextgen[k] = 0
			}
		}
		puzzle4i = nextgen
	}
	return len(puzzle4i)
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	grid := bytes.Split(data, []byte("\n"))
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '#' {
				puzzle3i[vec3i{x, y, 0}] = 0
				puzzle4i[vec4i{x, y, 0, 0}] = 0
			}
		}
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 1:", part2())
}
