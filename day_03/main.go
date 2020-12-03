//advent of code 2020, day 03, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type vec2i struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := bytes.Split(bytes.TrimSpace(data), []byte{10})
	part1, part2 := 0, 1
	slopes := []vec2i{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for i, vec := range slopes {
		y, x, count := vec.y, vec.x, 0
		for y < len(grid) {
			if grid[y][x] == '#' {
				count++
			}
			y += vec.y
			x = (x + vec.x) % len(grid[0])
		}
		if i == 1 {
			part1 = count
		}
		part2 *= count
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
