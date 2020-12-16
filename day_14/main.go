//advent of code 2020, day 14, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type item struct {
	mask      string
	addr, val uint64
}

var puzzle = []item{}

func solve() (uint64, uint64) {
	var (
		mem1         = map[uint64]uint64{}
		mem2         = map[uint64]uint64{}
		part1, part2 uint64
	)
	for _, item := range puzzle {
		xct := uint64(strings.Count(item.mask, "X"))
		for j := uint64(0); j < 1<<xct; j++ {
			val1, val2 := item.val, item.val
			addr1, addr2 := item.addr, item.addr
			var i, xpos uint64
			for ; i < 36; i++ {
				switch item.mask[36-1-i] {
				case '0':
					val1 &= ^(1 << i)
				case '1':
					val1 |= (1 << i)
					addr2 |= (1 << i)
				case 'X':
					addr2 &= ^(1 << i)
					addr2 |= ((j >> xpos) & 1) << i
					xpos++
				}
			}
			mem1[addr1] = val1
			mem2[addr2] = val2
		}
	}
	for _, val := range mem1 {
		part1 += val
	}
	for _, val := range mem2 {
		part2 += val
	}
	return part1, part2
}

func main() {
	var (
		reMask   = regexp.MustCompile(`^mask = (\w+)$`)
		reMemset = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
		mask     string
	)
	data, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		if reMask.MatchString(line) {
			mask = reMask.FindStringSubmatch(line)[1]
		} else if reMemset.MatchString(line) {
			memset := reMemset.FindStringSubmatch(line)
			addr, _ := strconv.ParseUint(memset[1], 10, 64)
			val, _ := strconv.ParseUint(memset[2], 10, 64)
			puzzle = append(puzzle, item{mask: mask, addr: addr, val: val})
		}
	}
	part1, part2 := solve()
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
