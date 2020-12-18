//advent of code 2020, day 18, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calc(s ...string) string {
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[2])
	if s[1] == "+" {
		return fmt.Sprintf("%d", a+b)
	}
	return fmt.Sprintf("%d", a*b)
}

func eval(s string, f bool) int {
	tok := strings.Fields(s)
	if len(tok) == 1 {
		res, _ := strconv.Atoi(tok[0])
		return res
	}
	f = f && strings.Index(s, "+") != -1
	for i, v := range tok {
		if !(v == "+" || v == "*") ||
			(f && v != "+") {
			continue
		}
		repl := calc(tok[i-1 : i+2]...)
		newTok := append(append(tok[:i-1], repl), tok[i+2:]...)
		return eval(strings.Join(newTok, " "), f)
	}
	return 0
}

func subexpr(s string, f bool) int {
	if i := strings.LastIndex(s, "("); i != -1 {
		j := strings.Index(s[i:], ")") + i
		e := eval(s[i+1:j], f)
		d := fmt.Sprintf("%d", e)
		return subexpr(s[:i]+d+s[j+1:], f)
	}
	return eval(s, f)
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	puzzle := strings.TrimSpace(string(data))
	part1, part2 := 0, 0
	for _, line := range strings.Split(puzzle, "\n") {
		part1 += subexpr(line, false)
		part2 += subexpr(line, true)
	}
	fmt.Println("part 1:", part1)
	fmt.Println("part 2:", part2)
}
