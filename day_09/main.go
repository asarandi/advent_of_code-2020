//advent of code 2020, day 09, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var numbers = []int{}

func part1() int {
	const w = 25

	var check = func(a []int, n int) bool {
		for j := 0; j < w-1; j++ {
			for k := j; k < w; k++ {
				if a[j]+a[k] == n {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i+w < len(numbers); i++ {
		n := numbers[i+w]
		if !check(numbers[i:i+w], n) {
			return n
		}
	}
	return 0
}

func part2(search int) int {
	for w := 2; w < len(numbers); w++ {
		sum := 0
		for _, n := range numbers[:w] {
			sum += n
		}
		for i := 0; i+w < len(numbers); i++ {
			if i > 0 {
				sum -= numbers[i-1]
				sum += numbers[i+w-1]
			}
			if sum == search {
				sort.IntSlice(numbers[i : i+w]).Sort()
				return numbers[i] + numbers[i+w-1]
			}
		}
	}
	return 0
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	for _, s := range strings.Fields(string(data)) {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	x := part1()
	fmt.Println("part 1:", x)
	fmt.Println("part 2:", part2(x))
}
