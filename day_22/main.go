//advent of code 2020, day 22, part 1 and 2
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func combat(one, two []int) (bool, []int) {
	for {
		if len(one) == 0 {
			return false, two
		}
		if len(two) == 0 {
			return true, one
		}
		a, b := one[0], two[0]
		one, two = one[1:], two[1:]
		if a > b {
			one = append(one, a, b)
		} else {
			two = append(two, b, a)
		}
	}
	return false, nil
}

func recursiveCombat(one, two []int) (bool, []int) {
	seen1, seen2 := map[string]bool{}, map[string]bool{}
	for {
		if len(one) == 0 {
			return false, two
		}
		if len(two) == 0 {
			return true, one
		}

		h1 := fmt.Sprintf("%#v", one)
		h2 := fmt.Sprintf("%#v", two)
		_, ok1 := seen1[h1]
		_, ok2 := seen2[h2]
		if ok1 || ok2 {
			return true, one
		}
		seen1[h1] = true
		seen2[h2] = true

		a, b := one[0], two[0]
		one, two = one[1:], two[1:]
		oneWins := a > b
		if len(one) >= a && len(two) >= b {
			oneCopy := append([]int(nil), one[:a]...)
			twoCopy := append([]int(nil), two[:b]...)
			oneWins, _ = recursiveCombat(oneCopy, twoCopy)
		}
		if oneWins {
			one = append(one, a, b)
		} else {
			two = append(two, b, a)
		}
	}
	return false, nil
}

func score(_ bool, deck []int) (res int) {
	for i := range deck {
		res += deck[i] * (len(deck) - i)
	}
	return
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(data), "\n\n")
	fields1 := strings.Fields(split[0])[2:]
	fields2 := strings.Fields(split[1])[2:]
	one, two := make([]int, 25), make([]int, 25)
	for i := 0; i < len(one); i++ {
		n, _ := strconv.Atoi(fields1[i])
		m, _ := strconv.Atoi(fields2[i])
		one[i], two[i] = n, m
	}
	fmt.Println("part 1:", score(combat(one, two)))
	fmt.Println("part 1:", score(recursiveCombat(one, two)))
}
