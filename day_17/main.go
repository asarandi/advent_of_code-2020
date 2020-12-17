//advent of code 2020, day 17, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type vec [4]int

func neighbors(puzzle map[vec]int, n int) map[vec]int {
	res := map[vec]int{}
	for cube := range puzzle {
		for i := 1; i < n; i++ {
			newCube := cube
			for j, v := range []int{i % 3, i / 3 % 3, i / 9 % 3, i / 27} {
				newCube[j] += []int{0, -1, 1}[v]
			}
			res[newCube] += 1
		}
	}
	return res
}

func solve(puzzle map[vec]int, n int) int {
	for i := 0; i < 6; i++ {
		nextgen := map[vec]int{}
		for k, v := range neighbors(puzzle, n) {
			_, exists := puzzle[k]
			if v == 2 && exists {
				nextgen[k] = 0
			} else if v == 3 {
				nextgen[k] = 0
			}
		}
		puzzle = nextgen
	}
	return len(puzzle)
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	grid := bytes.Split(data, []byte("\n"))
	one, two := map[vec]int{}, map[vec]int{}
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '#' {
				one[vec{x, y, 0, 0}] = 0
				two[vec{x, y, 0, 0}] = 0
			}
		}
	}
	fmt.Println("part 1:", solve(one, 27))
	fmt.Println("part 2:", solve(two, 81))
}
