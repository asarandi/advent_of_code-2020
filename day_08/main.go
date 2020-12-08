//advent of code 2020, day 08, part 1 and 2
package main

import (
	"fmt"
	"os"
)

type ins struct {
	op  string
	val int
}

var prog = []ins{}

func run() (acc, pc int) {
	seen := make([]bool, len(prog))
	for pc < len(prog) && !seen[pc] {
		seen[pc] = true
		switch prog[pc].op {
		case "jmp":
			pc += prog[pc].val
		case "acc":
			acc += prog[pc].val
			fallthrough
		default:
			pc += 1
		}
	}
	return
}

func part1() int {
	acc, _ := run()
	return acc
}

func part2() int {
	swap := map[string]string{"jmp": "nop", "nop": "jmp"}
	for i := range prog {
		if val, ok := swap[prog[i].op]; ok {
			prog[i].op = val
			acc, pc := run()
			if pc >= len(prog) {
				return acc
			}
			prog[i].op = swap[val]
		}
	}
	return 0
}

func main() {
	fp, _ := os.Open("input.txt")
	for {
		i := ins{}
		_, err := fmt.Fscanf(fp, "%s %d\n", &i.op, &i.val)
		if err != nil {
			break
		}
		prog = append(prog, i)
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
