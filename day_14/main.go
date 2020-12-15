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
	s         *string
	addr, val uint64
}

var (
	puzzle = []item{}
)

func part1() (res uint64) {
	mem := map[uint64]uint64{}
	for _, item := range puzzle {
		s, addr, val := item.s, item.addr, item.val
		for i := 0; i < 36; i++ {
			switch (*s)[36-1-i] {
			case '1':
				val |= 1 << i
			case '0':
				val &= ^(1 << i)
			}
		}
		mem[addr] = val
	}
	for _, val := range mem {
		res += val
	}
	return
}

func part2() (res uint64) {
	mem := map[uint64]uint64{}
	for _, item := range puzzle {
		s, val := item.s, item.val
		xct := uint64(strings.Count(*s, "X"))
		for j := uint64(0); j < 1<<xct; j++ {
			addr, xpos := item.addr, uint64(0)
			for i := 0; i < 36; i++ {
				switch (*s)[36-1-i] {
				case '1':
					addr |= (1 << i)
				case 'X':
					addr &= ^(1 << i)
					addr |= ((j >> xpos) & 1) << i
					xpos++
				}
			}
			mem[addr] = val
		}
	}
	for _, val := range mem {
		res += val
	}
	return
}

func main() {
	var (
		mask     []string
		reMask   = regexp.MustCompile(`^mask = (\w+)$`)
		reMemset = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	)
	data, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(data), "\n") {
		if reMask.MatchString(line) {
			mask = reMask.FindStringSubmatch(line)
		} else if reMemset.MatchString(line) {
			memset := reMemset.FindStringSubmatch(line)
			addr, _ := strconv.ParseUint(memset[1], 10, 64)
			val, _ := strconv.ParseUint(memset[2], 10, 64)
			puzzle = append(puzzle, item{s: &mask[1], addr: addr, val: val})
		}
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
