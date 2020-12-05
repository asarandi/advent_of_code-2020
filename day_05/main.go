//advent of code 2020, day 05, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	bits := map[byte]uint{'F': 0, 'B': 1, 'L': 0, 'R': 1}
	seats := map[uint]bool{}
	n, hi, lo := uint(0), uint(0), uint(1024)

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for _, item := range bytes.Fields(data) {
		n = 0
		for _, c := range item {
			n = (n << 1) | bits[c]
		}
		seats[n] = true
		if n < lo {
			lo = n
		}
		if n > hi {
			hi = n
		}
	}

	for ; lo < hi; lo++ {
		if _, ok := seats[lo]; !ok {
			break
		}
	}

	fmt.Println("part 1:", hi)
	fmt.Println("part 2:", lo)
}
