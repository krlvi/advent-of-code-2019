package day_1

import (
	"advent-of-code-2019/input"
	"fmt"
)

func Part1() {
	lines := input.Lines("day_1/in.txt")
	digits := input.ToInts(lines)
	sum := 0
	for _, d := range digits {
		sum += (d / 3) - 2
	}
	fmt.Println(sum)
}

func Part2() {
	lines := input.Lines("day_1/in.txt")
	digits := input.ToInts(lines)
	sum := 0
	for _, d := range digits {
		sum += req(d)
	}
	fmt.Println(sum)
}

func req(in int) int {
	f := (in / 3) - 2
	s := f
	for f > 0 {
		f = (f / 3) - 2
		if f > 0 {
			s += f
		}
	}
	return s
}
