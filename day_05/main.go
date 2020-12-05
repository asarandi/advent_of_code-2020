//advent of code 2020, day 05, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var (
	fblr   = map[byte]uint{'F': 0, 'B': 1, 'L': 0, 'R': 1}
	seats  map[uint]bool
	hi, lo uint = 0, 1024
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	seats = make(map[uint]bool)
	for _, item := range split {
		n := uint(0)
		for _, bit := range item {
			n = (n << 1) | fblr[bit]
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
		_, ok := seats[lo]
		if !ok {
			break
		}
	}
	fmt.Println("part 1:", hi)
	fmt.Println("part 2:", lo)
}
