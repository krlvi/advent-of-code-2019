package day_4

import (
	"fmt"
)

func Part1(){
	min := 136818
	max := 685979
	accepted := 0
	for i := min; i <= max; i++ {
		if increasing(toArr(i)) && anyAdj(toArr(i)) {
			accepted++
		}
	}
	fmt.Println(accepted)
}

func Part2(){
	min := 136818
	max := 685979
	accepted := 0
	for i := min; i <= max; i++ {
		if increasing(toArr(i)) && twoAdj(toArr(i)) {
			accepted++
		}
	}
	fmt.Println(accepted)
}

func toArr(i int) []int {
	a := []int{}
	for i > 0 {
		a = append(a, i % 10)
		i /= 10
	}
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func increasing(n []int) bool {
	last := 0
	for _, i := range n {
		if i < last {
			return false
		}
		last = i
	}
	return true
}

func anyAdj(n []int) bool {
	last := 0
	for _, i := range n {
		if i == last {
			return true
		}
		last = i
	}
	return false
}

func twoAdj(n []int) bool {
	last := 0
	seen := []int{}
	l := 0
	for _, i := range n {
		if i == last {
			l++
		} else {
			seen = append(seen, l)
			l = 0
		}
		last = i
	}
	seen = append(seen, l)
	for _, r := range seen {
		if r == 1 {
			return true
		}
	}
	return false
}
