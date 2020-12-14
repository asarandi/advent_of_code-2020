//advent of code 2020, day 10, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	nums = map[int]int{0: 1}
	max  int
)

func part1() int {
	a, b := 0, 0
	for i := 1; i <= max; i++ {
		if _, ok := nums[i]; ok {
			if _, ok := nums[i-1]; ok {
				a++
			} else {
				b++
			}
		}
	}
	return a * b
}

func part2() int {
	for i := 1; i <= max; i++ {
		if _, ok := nums[i]; ok {
			for j := 1; j <= 3; j++ {
				if v, ok := nums[i-j]; ok {
					nums[i] += v
				}
			}
		}
	}
	return nums[max]
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	for _, s := range strings.Fields(string(data)) {
		n, _ := strconv.Atoi(s)
		if n > max {
			max = n
		}
		nums[n] = 0
	}
	max += 3
	nums[max] = 0
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
