package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	nums := make(map[int]bool)
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums[n] = true
	}
	for n := range nums {
		x := 2020 - n
		if _, ok := nums[x]; ok {
			fmt.Println("part 1:", n*x)
			break
		}
	}
loop:
	for n := range nums {
		for m := range nums {
			if n == m {
				continue
			}
			x := 2020 - (n + m)
			if _, ok := nums[x]; ok {
				fmt.Println("part 2:", n*m*x)
				break loop
			}
		}
	}
}
