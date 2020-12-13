//advent of code 2020, day 13, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

var (
	timestamp  int
	busNumbers = map[int]int{}
)

func part1() int {
	leastMinutes, res := 1<<63-1, 0
	for busId := range busNumbers {
		minutes := ((timestamp/busId)+1)*busId - timestamp
		if minutes < leastMinutes {
			leastMinutes = minutes
			res = minutes * busId
		}
	}
	return res
}

func part2() int {
	prod, sum := 1, 0
	modinv := func(a, b int) int {
		return int((new(big.Int)).ModInverse(big.NewInt(int64(a)), big.NewInt(int64(b))).Int64())
	}

	for k := range busNumbers {
		prod *= k
	}
	for k, v := range busNumbers {
		p := prod / k
		n := modinv(p, k) * p * v
		sum += n
	}
	return sum % prod
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(data), "\n")
	timestamp, _ = strconv.Atoi(split[0])
	for i, s := range strings.Split(split[1], ",") {
		if n, err := strconv.Atoi(s); err == nil {
			busNumbers[n] = n - i
		}
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
