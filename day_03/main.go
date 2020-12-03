//advent of code 2020, day 03, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type pos struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := bytes.Split(bytes.TrimSpace(data), []byte{10})
	res := make([]int, 0)
	for _, p := range []pos{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		y, x, count := p.y, p.x, 0
		for y < len(grid) {
			if grid[y][x] == '#' {
				count++
			}
			y += p.y
			x = (x + p.x) % len(grid[0])
		}
		res = append(res, count)
	}
	p1, p2 := res[1], 1
	for _, val := range res {
		p2 *= val
	}
	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}
