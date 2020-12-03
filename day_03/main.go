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

var grid [][]byte

func part1() (res int) {
	n, m := len(grid), len(grid[0])
	for y, x := 1, 3; y < n; {
		if grid[y][x] == '#' {
			res++
		}
		y, x = y+1, (x+3)%m
	}
	return
}

func part2() (res int) {
	slopes := []vec2i{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	n, m := len(grid), len(grid[0])
	res = 1
	for _, v := range slopes {
		y, x, count := v.y, v.x, 0
		for y < n {
			if grid[y][x] == '#' {
				count++
			}
			y, x = y+v.y, (x+v.x)%m
		}
		res *= count
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid = bytes.Split(bytes.TrimSpace(data), []byte{10})
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
