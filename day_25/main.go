//advent of code 2020, day 25, part 1
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	input := strings.Fields(string(data))
	card, _ := strconv.Atoi(input[0])
	door, _ := strconv.Atoi(input[1])
	var i, j, k int
	for i, k = 0, 1; k != card; i++ {
		k = (k * 7) % 20201227
	}
	for j, k = 0, 1; j < i; j++ {
		k = (k * door) % 20201227
	}
	fmt.Println("part 1:", k)
}
