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

var passports []map[string]string

var fields = map[string]func(string) bool{
	"byr": byr, "iyr": iyr, "eyr": eyr, "hgt": hgt, "hcl": hcl, "ecl": ecl, "pid": pid,
}

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
		pass := make(map[string]string)
		for _, tok := range strings.Fields(string(item)) {
			key := strings.Split(tok, ":")[0]
			value := strings.Split(tok, ":")[1]
			pass[key] = value
		}
		passports = append(passports, pass)
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}

func byr(s string) bool {
	return regexp.MustCompile(`^[\d]{4}$`).MatchString(s) &&
		isrange(s, 1920, 2002)
}

func iyr(s string) bool {
	return regexp.MustCompile(`^[\d]{4}$`).MatchString(s) &&
		isrange(s, 2010, 2020)
}

func eyr(s string) bool {
	return regexp.MustCompile(`^[\d]{4}$`).MatchString(s) &&
		isrange(s, 2020, 2030)
}

func hcl(s string) bool {
	return regexp.MustCompile(`^#[0-9a-f]{6}$`).MatchString(s)
}

func pid(s string) bool {
	return regexp.MustCompile(`^[\d]{9}$`).MatchString(s)
}

func hgt(s string) (res bool) {
	n := len(s) - 2
	value, unit := s[:n], s[n:]
	res = res || (unit == "cm" && isrange(value, 150, 193))
	res = res || (unit == "in" && isrange(value, 59, 76))
	return res
}

func ecl(s string) bool {
	return regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`).MatchString(s)
}

func isrange(s string, min, max int) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return min <= n && n <= max
}
