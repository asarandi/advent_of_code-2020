//advent of code 2020, day 15, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var puzzle = map[int]int{}

func run(n int) (res int) {
	p := map[int]int{}
	for i, v := range puzzle {
		p[i] = v
	}
	v, seen := 0, false
	for i := len(p) + 1; i <= n; i++ {
		if !seen {
			res = 0
		} else {
			res = i - 1 - v
		}
		v, seen = p[res]
		p[res] = i
	}
	return
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))
	for i, v := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(v)
		puzzle[n] = i + 1
	}
	fmt.Println("part 1:", run(2020))
	fmt.Println("part 2:", run(30000000))
}
