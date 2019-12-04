package day_3

import (
	"advent-of-code-2019/input"
	"fmt"
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/sets/linkedhashset"
	"math"
	"strconv"
	"strings"
)

func Part2() {
	lines := input.Lines("day_3/in.txt")
	firstWire, firstWireDists := lineCoords(toInstArray(lines[0]))
	secondWire, secondWireDists := lineCoords(toInstArray(lines[1]))
	intersections := setIntersection(firstWire, secondWire)

	minDist := math.MaxFloat64
	for _, x := range intersections.Values() {
		coords := x.(pos)
		if coords.X == 0 && coords.Y == 0 {
			continue
		}
		minDist = math.Min(minDist,
			float64(firstWireDists[coords] + secondWireDists[coords]))
	}
	fmt.Println(minDist)
}

func Part1() {
	lines := input.Lines("day_3/in.txt")
	firstWire, _ := lineCoords(toInstArray(lines[0]))
	secondWire, _ := lineCoords(toInstArray(lines[1]))
	intersections := setIntersection(firstWire, secondWire)

	minDist := math.MaxFloat64
	for _, x := range intersections.Values() {
		coords := x.(pos)
		if coords.X == 0 && coords.Y == 0 {
			continue
		}
		minDist = math.Min(minDist,
			math.Abs(float64(coords.X)) + math.Abs(float64(coords.Y)))
	}
	fmt.Println(minDist)
}

func setIntersection(a, b sets.Set) sets.Set {
	out := linkedhashset.New()
	for _, x := range a.Values() {
		if b.Contains(x) {
			out.Add(x)
		}
	}
	return out
}

func lineCoords(in []instr) (sets.Set, map[pos]int) {
	s := linkedhashset.New()
	s.Add(pos{0, 0})
	current := pos{Y:0, X:0}
	totalDist := 0
	distToPoint := map[pos]int{}
	for _, i := range in {
		if i.Dir == "R" {
			for x := 1; x <= i.Magn; x++ {
				totalDist++
				p := pos{current.Y, current.X + x}
				s.Add(p)
				distToPoint[p] = totalDist
			}
			current = s.Values()[s.Size()-1].(pos)
			continue
		}
		if i.Dir == "L" {
			for x := 1; x <= i.Magn; x++ {
				totalDist++
				p := pos{current.Y, current.X - x}
				s.Add(p)
				distToPoint[p] = totalDist
			}
			current = s.Values()[s.Size()-1].(pos)
			continue
		}
		if i.Dir == "U" {
			for x := 1; x <= i.Magn; x++ {
				totalDist++
				p := pos{current.Y + x, current.X}
				s.Add(p)
				distToPoint[p] = totalDist
			}
			current = s.Values()[s.Size()-1].(pos)
			continue
		}
		if i.Dir == "D" {
			for x := 1; x <= i.Magn; x++ {
				totalDist++
				p := pos{current.Y - x, current.X}
				s.Add(p)
				distToPoint[p] = totalDist
			}
			current = s.Values()[s.Size()-1].(pos)
			continue
		}
	}
	return s, distToPoint
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

type pos struct {
	Y int
	X int
}

type instr struct {
	Dir  string
	Magn int
}
