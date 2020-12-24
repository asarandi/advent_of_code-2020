//advent of code 2020, day 21, part 1 and part 2
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	ingredients = map[string]int{} // counts
	allergens   = map[string]map[string]int{}
	unsafe      = map[string]bool{}
)

func part1() int {
	res := 0
	for alle, ingr := range allergens {
		max := 0
		for _, ct := range ingr {
			if ct > max {
				max = ct
			}
		}
		safe := []string{}
		for name, ct := range ingr {
			if ct == max {
				unsafe[name] = true
			} else {
				safe = append(safe, name)
			}
		}
		for _, name := range safe {
			delete(allergens[alle], name)
		}
	}
	for name, ct := range ingredients {
		if !unsafe[name] {
			res += ct
		}
	}
	return res
}

func part2() string {
	identified := map[string]string{}
	for len(identified) != len(allergens) {
		for alle, ingr := range allergens {
			if len(ingr) == 1 {
				var name string
				for name, _ = range ingr {
					identified[name] = alle
				}
				for k := range allergens {
					if k != alle {
						delete(allergens[k], name)
					}
				}
			}
		}
	}
	results := []string{}
	for i := range identified {
		results = append(results, i)
	}
	sort.Slice(results, func(i, j int) bool {
		return identified[results[i]] < identified[results[j]]
	})
	return strings.Join(results, ",")
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " (contains ")
		split[1] = strings.NewReplacer(",", "", ")", "").Replace(split[1])
		for _, i := range strings.Fields(split[0]) {
			ingredients[i]++
		}
		for _, a := range strings.Fields(split[1]) {
			if _, ok := allergens[a]; !ok {
				allergens[a] = map[string]int{}
			}
			for _, i := range strings.Fields(split[0]) {
				allergens[a][i]++
			}
		}
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
