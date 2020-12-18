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

func eval(s string, prio bool) int {
	tok := strings.Fields(s)
	if len(tok) == 1 {
		res, _ := strconv.Atoi(tok[0])
		return res
	}
	prio = prio && strings.Index(s, "+") != -1
	i := 1
	for i = range tok {
		if tok[i] == "+" || tok[i] == "*" {
			if prio && tok[i] == "+" {
				break
			} else if !prio {
				break
			}
		}
	}
	repl := calc(tok[i-1 : i+2]...)
	newTok := append(append(tok[:i-1], repl), tok[i+2:]...)
	return eval(strings.Join(newTok, " "), prio)
}

func subexpr(s string, prio bool) int {
	i := strings.LastIndex(s, "(")
	if i == -1 {
		return eval(s, prio)
	}
	j := strings.Index(s[i:], ")") + i
	sub := eval(s[i+1:j], prio)
	repl := fmt.Sprintf("%d", sub)
	return subexpr(s[:i]+repl+s[j+1:], prio)
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
