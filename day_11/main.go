//advent of code 2020, day 11, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type pos struct {
	y, x int
}

var adjacent = []pos{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func count(grid [][]byte, i, j int, f bool) int {
	res, n, m := 0, len(grid), len(grid[0])
	for _, a := range adjacent {
		y, x := i, j
		for {
			y, x = y+a.y, x+a.x
			if y < 0 || x < 0 {
				break
			}
			if y >= n || x >= m {
				break
			}
			if grid[y][x] == 'L' {
				break
			}
			if grid[y][x] == '#' {
				res += 1
				break
			}
			if !f {
				break
			}
		}
	}
	return res
}

func occupied(grid [][]byte) (res int) {
	for _, row := range grid {
		res += bytes.Count(row, []byte("#"))
	}
	return
}

func change(grid [][]byte, taken, free []pos) {
	for _, p := range taken {
		grid[p.y][p.x] = 'L'
	}
	for _, p := range free {
		grid[p.y][p.x] = '#'
	}
}

func list(grid [][]byte, first bool) (taken, free []pos) {
	b := map[bool]int{false: 4, true: 5}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			c := grid[y][x]
			if !(c == 'L' || c == '#') {
				continue
			}
			k := count(grid, y, x, first)
			if c == 'L' && k == 0 {
				free = append(free, pos{y, x})
			}
			if c == '#' && k >= b[first] {
				taken = append(taken, pos{y, x})
			}
		}
	}
	return
}

func conway(grid [][]byte, first bool) int {
	before, after := 0, 1
	for before != after {
		taken, free := list(grid, first)
		change(grid, taken, free)
		before, after = after, occupied(grid)
	}
	return after
}

func getInput() [][]byte {
	data, _ := ioutil.ReadFile("input.txt")
	return bytes.Fields(data)
}

func main() {
	fmt.Println("part 1:", conway(getInput(), false))
	fmt.Println("part 2:", conway(getInput(), true))
}
