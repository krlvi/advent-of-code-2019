package day_3

import (
	"advent-of-code-2019/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1() {
	lines := input.Lines("day_3/in.txt")

	first := toInstArray(lines[0])
	second := toInstArray(lines[1])

	//fmt.Println(first)
	//fmt.Println(second)

	grid := initialize(20000)

	x := len(grid) / 2
	y := len(grid) / 2

	grid[y][x] = 1

	for _, in := range first {
		sx := x
		sy := y
		if in.Dir == "R" {
			x = x + in.Magn
		} else if in.Dir == "L" {
			x = x - in.Magn
		} else if in.Dir == "U" {
			y = y - in.Magn
		} else if in.Dir == "D" {
			y = y + in.Magn
		}
		flagBetween(grid, sx, sy, x, y, 1)
	}

	x = len(grid) / 2
	y = len(grid) / 2

	for _, in := range second {
		sx := x
		sy := y
		if in.Dir == "R" {
			x = x + in.Magn
		} else if in.Dir == "L" {
			x = x - in.Magn
		} else if in.Dir == "U" {
			y = y - in.Magn
		} else if in.Dir == "D" {
			y = y + in.Magn
		}
		flagBetween(grid, sx, sy, x, y, 5)
	}

	//for i := 0; i< len(grid); i++ {
	//	fmt.Println(grid[i])
	//}
	minDist := 100000.0
	for _, coords := range findInter(grid) {
		d := manhattanDistance(len(grid)/2, len(grid)/2, coords[1], coords[0])
		if d < minDist {
			minDist = d
		}
	}
	fmt.Println(minDist)
}

func manhattanDistance(fromX, fromY, toX, toY int) float64 {
	return math.Abs(float64(toX-fromX)) + math.Abs(float64(toY-fromY))
}

func findInter(grid [][]int) [][]int {
	out := [][]int{}
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid); x++ {
			if grid[y][x] == 6 {
				out = append(out, []int{y, x})
			}
		}
	}
	return out
}

func Part2() {
	//lines := input.Lines("day_3/in.txt")
	//fmt.Println(lines)
}

func flagBetween(grid [][]int, sX, sY, eX, eY, amnt int) {
	if sX < eX {
		for i := sX + 1; i <= eX; i++ {
			grid[sY][i] = grid[sY][i] + amnt
		}
	}
	if sX > eX {
		for i := sX - 1; i >= eX; i-- {
			grid[sY][i] = grid[sY][i] + amnt
		}
	}
	if sY < eY {
		for i := sY + 1; i <= eY; i++ {
			grid[i][sX] = grid[i][sX] + amnt
		}
	}
	if sY > eY {
		for i := sY - 1; i >= eY; i-- {
			grid[i][sX] = grid[i][sX] + amnt
		}
	}
}

func initialize(dim int) [][]int {
	out := [][]int{}
	for i := 0; i < dim; i++ {
		out = append(out, make([]int, dim))
	}
	return out
}

func toInstArray(line string) []instr {
	tokens := strings.Split(line, ",")
	out := []instr{}
	for _, t := range tokens {
		magn, _ := strconv.Atoi(t[1:])
		instr := instr{
			Dir:  string(t[0]),
			Magn: magn,
		}
		out = append(out, instr)
	}
	return out
}

type instr struct {
	Dir  string
	Magn int
}
