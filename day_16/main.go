//advent of code 2020, day 16, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var rules = [][]int{}
var myTicket = []int{}
var tickets = [][]int{}

func isValid(rule []int, val int) bool {
	return (rule[0] <= val && val <= rule[1]) ||
		(rule[2] <= val && val <= rule[3])
}

// part1() filters invalid tickets
func part1() int {
	valid, res := [][]int{}, 0
loop:
	for _, t := range tickets {
		for _, n := range t {
			ok := false
			for _, r := range rules {
				ok = ok || isValid(r, n)
			}
			if !ok {
				res += n
				continue loop
			}
		}
		valid = append(valid, t)
	}
	tickets = valid
	return res
}

// does rule `r' match tickets column `c'
func isRC(r, c int) bool {
	nums := []int{}
	for _, t := range tickets {
		nums = append(nums, t[c])
	}
	res := true
	for _, n := range nums {
		res = res && isValid(rules[r], n)
	}
	return res
}

// rule 2 col, col 2 rule
var rtoc = map[int]int{}
var ctor = map[int]int{}

// map rule number to tickets column
func match() map[int]map[int]bool {
	res := map[int]map[int]bool{}
	for i := range rules {
		if _, ok := rtoc[i]; ok {
			continue
		}
		res[i] = map[int]bool{}
		for j := range rules {
			if _, ok := ctor[j]; ok {
				continue
			}
			if isRC(i, j) {
				res[i][j] = true
			}
		}
	}
	return res
}

func mapping() {
loop:
	for len(rtoc) != len(rules) {
		for rule, columns := range match() {
			if len(columns) == 1 {
				for k := range columns {
					rtoc[rule] = k
					ctor[k] = rule
				}
				continue loop
			}
		}
	}
}

func part2() int {
	mapping()
	res := 1
	for i := 0; i < 6; i++ {
		res *= myTicket[rtoc[i]]
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(data), "\n\n")
	rules = parseRules(split[0])
	myTicket = parseTickets(split[1])[0]
	tickets = parseTickets(split[2])
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}

func parseRules(s string) [][]int {
	res := [][]int{}
	re := regexp.MustCompile(`^.*: (\d+)-(\d+) or (\d+)-(\d+)$`)
	s = strings.TrimSpace(s)
	for _, line := range strings.Split(s, "\n") {
		subs := re.FindStringSubmatch(line)
		rule := []int{}
		for i := 1; i < 5; i++ {
			n, _ := strconv.Atoi(subs[i])
			rule = append(rule, n)
		}
		res = append(res, rule)
	}
	return res
}

func parseTickets(s string) [][]int {
	res := [][]int{}
	s = strings.TrimSpace(s)
	split := strings.Split(s, "\n")
	for i := 1; i < len(split); i++ { //skip 1st
		ticket := []int{}
		for _, val := range strings.Split(split[i], ",") {
			n, _ := strconv.Atoi(val)
			ticket = append(ticket, n)
		}
		res = append(res, ticket)
	}
	return res
}
