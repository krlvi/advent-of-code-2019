package day_2

import (
	"advent-of-code-2019/input"
	"fmt"
	"strings"
)

func Part2() {
	input := in()
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			in := make([]int, len(input))
			copy(in, input)

			in[1] = n
			in[2] = v
			run(in)
			if in[0] == 19690720 {
				fmt.Println(100*n + v)
			}
		}
	}
}

func Part1() {
	in := in()
	in[1] = 12
	in[2] = 2
	run(in)
	fmt.Println(in[0])
}

func run(in []int) []int {
	for i := 0; i < len(in)-3; i = i + 4 {
		op := in[i]
		if op == 1 {
			in[in[i+3]] = in[in[i+1]] + in[in[i+2]]
		} else if op == 2 {
			in[in[i+3]] = in[in[i+1]] * in[in[i+2]]
		} else if op == 99 {
			break
		}
	}
	return in
}

func in() []int {
	lines := input.Lines("day_2/in.txt")
	fields := strings.Split(lines[0], ",")
	return input.ToInts(fields)
}
