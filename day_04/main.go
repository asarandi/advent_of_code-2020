//advent of code 2020, day 04, part 1 and 2
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var fields = map[string]func(string) bool{
	"byr": byr,
	"iyr": iyr,
	"eyr": eyr,
	"hgt": hgt,
	"hcl": hcl,
	"ecl": ecl,
	"pid": pid,
}

var passports []map[string]string

func part1() (res int) {
loop:
	for _, pass := range passports {
		for f := range fields {
			_, ok := pass[f]
			if !ok {
				continue loop
			}
		}
		res++
	}
	return
}

func part2() (res int) {
loop:
	for _, pass := range passports {
		for f, check := range fields {
			value, ok := pass[f]
			if !ok || !check(value) {
				continue loop
			}
		}
		res++
	}
	return
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	split := bytes.Split(bytes.TrimSpace(data), []byte("\n\n"))
	passports = make([]map[string]string, 0)
	for _, item := range split {
		prep := bytes.Split(item, []byte("\n"))
		join := bytes.Join(prep, []byte(" "))
		tokens := strings.Split(string(join), " ")
		pass := make(map[string]string)
		for _, tok := range tokens {
			key := strings.Split(tok, ":")[0]
			value := strings.Split(tok, ":")[1]
			pass[key] = value
		}
		passports = append(passports, pass)
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}

/*
*
 */

func byr(s string) bool {
	return isrange(s, 1920, 2002)
}

func iyr(s string) bool {
	return isrange(s, 2010, 2020)
}

func eyr(s string) bool {
	return isrange(s, 2020, 2030)
}

func hcl(s string) bool {
	return s[0] == '#' && isalnum(s[1:])
}

func pid(s string) bool {
	return len(s) == 9 && isdigit(s)
}

func hgt(s string) (res bool) {
	n := len(s) - 2
	value, unit := s[:n], s[n:]
	res = res || (unit == "cm" && isrange(value, 150, 193))
	res = res || (unit == "in" && isrange(value, 59, 76))
	return res
}

func ecl(s string) bool {
	colors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := colors[s]
	return ok
}

/*
*
 */

var isdigit = regexp.MustCompile(`^[0-9]+$`).MatchString
var islower = regexp.MustCompile(`^[a-z]+$`).MatchString
var isupper = regexp.MustCompile(`^[A-Z]+$`).MatchString
var isalpha = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var isalnum = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func isrange(s string, min, max int) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n >= min && n <= max
}
