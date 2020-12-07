//advent of code 2020, day 07, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type qs struct {
	q int
	s string
}

var bags = map[string][]qs{}

func part1() int {
	var u = map[string]bool{}
	var f func(s string)
	f = func(s string) {
		for k, v := range bags {
			for _, b := range v {
				if b.s == s {
					u[k] = true
					f(k)
				}
			}
		}
	}
	f("shiny gold bag")
	return len(u)
}

func part2() int {
	var f func(string) int
	f = func(s string) (res int) {
		res = 1
		for _, b := range bags[s] {
			res += b.q * f(b.s)
		}
		return
	}
	return f("shiny gold bag") - 1
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	input := strings.ReplaceAll(string(data), "bags", "bag")
	input = strings.ReplaceAll(input, ".", "")
	input = strings.TrimSpace(input)
	for _, line := range strings.Split(input, "\n") {
		pc := strings.Split(line, " contain ")
		ch := []qs{}
		for _, b := range strings.Split(pc[1], ", ") {
			if !strings.Contains(b, "no other bag") {
				s := strings.SplitN(b, " ", 2)
				q, _ := strconv.Atoi(s[0])
				ch = append(ch, qs{q, s[1]})
			}
		}
		bags[pc[0]] = ch
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
