//advent of code 2020, day 06, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	part1, part2 := 0, 0
	for _, group := range bytes.Split(data, []byte("\n\n")) {
		counts := map[byte]int{}
		persons := bytes.Fields(group)
		for _, p := range persons {
			for i := 0; i < len(p); i++ {
				counts[p[i]] += 1
			}
		}
		part1 += len(counts)
		for _, v := range counts {
			if v == len(persons) {
				part2 += 1
			}
		}
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
