//advent of code 2020, day 08, part 1 and 2
package main

import (
	"fmt"
	"os"
)

type ins struct {
	op  string
	arg int
}

var bootcode = []ins{}

func run() (acc, pc int) {
	seen := make([]bool, len(bootcode))
	for pc < len(bootcode) && !seen[pc] {
		seen[pc] = true
		switch bootcode[pc].op {
		case "jmp":
			pc += bootcode[pc].arg
		case "acc":
			acc += bootcode[pc].arg
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
	changes := map[string]string{"jmp": "nop", "nop": "jmp"}
	for i := range bootcode {
		newOp, ok := changes[bootcode[i].op]
		if ok {
			bootcode[i].op = newOp
			acc, pc := run()
			bootcode[i].op = changes[newOp] //restore
			if pc >= len(bootcode) {
				return acc
			}
		}
	}
	return 0
}

func main() {
	fp, _ := os.Open("input.txt")
	for {
		i := ins{}
		_, err := fmt.Fscanf(fp, "%s %d\n", &i.op, &i.arg)
		if err != nil {
			break
		}
		bootcode = append(bootcode, i)
	}
	fmt.Println("part 1:", part1())
	fmt.Println("part 2:", part2())
}
