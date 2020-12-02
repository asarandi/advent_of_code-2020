//advent of code 2020, day 02, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type password struct {
	a, b int
	c    rune
	s    string
}

var puzzle = make(map[password]bool)

func part1() (res int) {
	for p := range puzzle {
		n := strings.Count(p.s, string(p.c))
		if p.a <= n && n <= p.b {
			res++
		}
	}
	return
}

func part2() (res int) {
	for p := range puzzle {
		a := rune(p.s[p.a-1]) == p.c
		b := rune(p.s[p.b-1]) == p.c
		if a != b {
			res++
		}
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)
	lines := strings.TrimSpace(string(data))
	for _, line := range strings.Split(lines, "\n") {
		subs := re.FindStringSubmatch(line)
		if len(subs) != 5 {
			panic(subs)
		}
		a, _ := strconv.Atoi(subs[1])
		b, _ := strconv.Atoi(subs[2])
		c, s := rune(subs[3][0]), subs[4]
		puzzle[password{a, b, c, s}] = true
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
